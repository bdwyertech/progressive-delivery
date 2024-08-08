package pdtesting

import (
	"context"
	"net"
	"testing"

	"github.com/go-logr/logr"
	pb "github.com/weaveworks/progressive-delivery/pkg/api/prog"
	"github.com/weaveworks/progressive-delivery/pkg/kube"
	"github.com/weaveworks/progressive-delivery/pkg/server"
	"github.com/weaveworks/progressive-delivery/pkg/services/crd"
	"github.com/weaveworks/weave-gitops/core/clustersmngr"
	"github.com/weaveworks/weave-gitops/core/clustersmngr/cluster"
	"github.com/weaveworks/weave-gitops/core/clustersmngr/fetcher"
	"github.com/weaveworks/weave-gitops/core/nsaccess/nsaccessfakes"
	wkube "github.com/weaveworks/weave-gitops/pkg/kube"
	"github.com/weaveworks/weave-gitops/pkg/server/auth"
	"github.com/weaveworks/weave-gitops/pkg/testutils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	v1 "k8s.io/api/core/v1"
	v1a "k8s.io/client-go/kubernetes/typed/authorization/v1"
	"k8s.io/client-go/rest"
)

func MakeGRPCServer(
	t *testing.T,
	cfg *rest.Config,
	k8sEnv *testutils.K8sTestEnv,
) pb.ProgressiveDeliveryServiceClient {
	log := logr.Discard()
	ctx := context.Background()

	cl, err := RestConfigToCluster(k8sEnv.Rest)
	if err != nil {
		t.Fatal(err)
	}

	fetcher := fetcher.NewSingleClusterFetcher(cl)

	nsChecker := nsaccessfakes.FakeChecker{}
	nsChecker.FilterAccessibleNamespacesStub = func(ctx context.Context, _ v1a.AuthorizationV1Interface, n []v1.Namespace) ([]v1.Namespace, error) {
		// Pretend the user has access to everything
		return n, nil
	}

	clustersManager := clustersmngr.NewClustersManager([]clustersmngr.ClusterFetcher{fetcher}, &nsChecker, log)

	_ = clustersManager.UpdateClusters(ctx)
	_ = clustersManager.UpdateNamespaces(ctx)

	opts := server.ServerOpts{
		ClustersManager: clustersManager,
		CRDService:      crd.NewNoCacheFetcher(clustersManager),
		Logger:          logr.Discard(),
	}

	pdServer, _ := server.NewProgressiveDeliveryServer(opts)
	lis := bufconn.Listen(1024 * 1024)
	principal := auth.NewUserPrincipal(auth.Token("1234"))
	s := grpc.NewServer(
		withClientsPoolInterceptor(clustersManager, cfg, principal),
	)

	pb.RegisterProgressiveDeliveryServiceServer(s, pdServer)

	dialer := func(context.Context, string) (net.Conn, error) {
		return lis.Dial()
	}

	go func(tt *testing.T) {
		if err := s.Serve(lis); err != nil {
			tt.Error(err)
		}
	}(t)

	conn, err := grpc.DialContext(
		context.Background(),
		"bufnet",
		grpc.WithContextDialer(dialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		s.GracefulStop()
		conn.Close()
	})

	return pb.NewProgressiveDeliveryServiceClient(conn)
}

func RestConfigToCluster(cfg *rest.Config) (cluster.Cluster, error) {
	return cluster.NewSingleCluster(
		cluster.DefaultCluster,
		cfg,
		kube.CreateScheme(),
		wkube.UserPrefixes{},
		cluster.DefaultKubeConfigOptions...,
	)
}

func withClientsPoolInterceptor(clustersManager clustersmngr.ClustersManager, config *rest.Config, user *auth.UserPrincipal) grpc.ServerOption {
	return grpc.UnaryInterceptor(func(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if err := clustersManager.UpdateClusters(ctx); err != nil {
			return nil, err
		}
		if err := clustersManager.UpdateNamespaces(ctx); err != nil {
			return nil, err
		}

		clustersManager.UpdateUserNamespaces(ctx, user)

		ctx = auth.WithPrincipal(ctx, user)

		clusterClient, err := clustersManager.GetImpersonatedClient(ctx, user)
		if err != nil {
			return nil, err
		}

		ctx = context.WithValue(ctx, clustersmngr.ClustersClientCtxKey, clusterClient)

		return handler(ctx, req)
	})
}
