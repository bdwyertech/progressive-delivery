.PHONY: proto test lint tools dependencies js-lib publish clean dev-server

CURRENT_DIR := $(shell pwd)

proto: ## Generate protobuf files
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
	  github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
	  google.golang.org/protobuf/cmd/protoc-gen-go
	go install github.com/grpc-ecosystem/protoc-gen-grpc-gateway-ts
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0
	go install github.com/bufbuild/buf/cmd/buf@v1.1.0
	buf generate

test: ## Run tests
	go test -v ./...

lint:
	$(CURRENT_DIR)/tools/go-lint

tools: ## Install Go tools
	cat $(CURRENT_DIR)/tools/tools.go | awk -F'"' '/_/{print $$2}' | xargs -tI % go install %

dependencies: tools ## Install build dependencies
	$(CURRENT_DIR)/tools/download-deps.sh $(CURRENT_DIR)/tools/dependencies.toml

create-dev-cluster:
	$(CURRENT_DIR)/tools/bin/kind create cluster --name pd-dev

delete-dev-cluster:
	$(CURRENT_DIR)/tools/bin/kind delete cluster --name pd-dev

dev-cluster: ## Start a dev server with Tilt
	$(CURRENT_DIR)/tools/bin/tilt up

ui/lib/node_modules:
	cd ui/lib && npm install

ui/lib/dist/index.js: ui/lib/node_modules
	cd ui/lib && yarn compile

ui/lib/dist/package.json: ui/lib/package.json
	cp ui/lib/package.json ui/lib/dist
	cp ui/lib/.npmrc ui/lib/dist

js-lib: ui/lib/dist/index.js ui/lib/dist/package.json

publish-js-lib: js-lib
	cd ui/lib/dist && npm publish --access=public

clean:
	rm -rf ui/lib/dist
