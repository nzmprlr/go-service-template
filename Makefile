MODULE := $(shell awk 'NR==1 {print $$2}' go.mod)

.PHONY: build

all: build run

build: GO_BUILD_MODULE = -X "$(MODULE)/config.module"=$(MODULE)
build: GO_BUILD_ID = -X "$(MODULE)/config.id"=$(shell git rev-list --count HEAD)
build: GO_BUILD_REV = -X "$(MODULE)/config.rev"=$(shell git rev-parse --short HEAD)
build: GO_BUILT_AT = -X "$(MODULE)/config.builtAt"=$(shell date -u +%Y-%m-%d_%T_%Z)
build:
	go build -o $(MODULE) -ldflags "$(GO_BUILD_MODULE) $(GO_BUILD_ID) $(GO_BUILD_REV) $(GO_BUILT_AT)"

run:
	./$(MODULE)

test: --mockgen
	go test -cover ./...

test-cover: --mockgen
	go test -coverprofile=$(MODULE)-cover.txt ./...
	go tool cover -html=$(MODULE)-cover.txt -o $(MODULE)-cover.html
	open $(MODULE)-cover.html

clean:
	find . -name "$(MODULE)*" -maxdepth 1 -type f -print -delete

docker: docker-build docker-run

docker-build: TAG ?= $(shell git rev-parse --short HEAD)
docker-build:
	docker build -t $(MODULE):$(TAG) .

docker-run: APP_ENV ?= local
docker-run: PORT ?= 8080
docker-run: TAG ?= $(shell git rev-parse --short HEAD)
docker-run:
	docker run -it --rm -p $(PORT):8080 -e APP_ENV=$(APP_ENV) $(MODULE):$(TAG)

--mockgen: MOCKGEN_VER = 1.5.0
--mockgen:
	@if ! test -f `go env GOPATH`/bin/mockgen; then go install github.com/golang/mock/mockgen@v$(MOCKGEN_VER); fi
	`go env GOPATH`/bin/mockgen -source api/service.go -destination mock/mock-service.go $(MODULE)
	`go env GOPATH`/bin/mockgen -source api/event.go -destination mock/mock-event.go $(MODULE)
	`go env GOPATH`/bin/mockgen -source api/data.go -destination mock/mock-data.go $(MODULE)
	`go env GOPATH`/bin/mockgen -source api/call.go -destination mock/mock-call.go $(MODULE)

godoc: 
	@if ! test -f `go env GOPATH`/bin/godoc; then go install golang.org/x/tools/cmd/godoc; fi
	`go env GOPATH`/bin/godoc -play -http=:35035&
	@sleep 5; python -m webbrowser http://localhost:35035/pkg/$(MODULE)
kill-godoc:
	pkill godoc

redoc: generate-swagger --redoc --kill-swagger
swagger: generate-swagger --swagger --kill-swagger

markdown: generate-swagger
	swagger generate markdown -f=$(MODULE)-swagger.yaml --output=$(MODULE)-swagger.md

--kill-swagger:
	@sleep 7; pkill swagger

--swagger:
	swagger serve -q -F=swagger -p=9035 $(MODULE)-swagger.yaml&

--redoc:
	swagger serve -q -p=9035 $(MODULE)-swagger.yaml&

generate-swagger: --check-swagger
	$(eval SED=$(if $(filter Darwin, $(shell uname -s)), sed -i '', sed -i))
	$(eval TEMP=$(shell mktemp -d))
	@cp doc/swagger.go $(TEMP)/swagger.go
	@$(SED) "s/{REVISION}/$(shell git rev-parse --short HEAD)/g" doc/swagger.go
	@$(SED) "s/{SWAGGER_HOST}/$(SWAGGER_HOST)/g" doc/swagger.go
	swagger generate spec -o $(MODULE)-swagger.yaml --scan-models
	@cp $(TEMP)/swagger.go doc/swagger.go
	@rm -rf $(TEMP)

--check-swagger:
ifndef SWAGGER_HOST
	$(error SWAGGER_HOST variable is required to make.)
endif
	@if ! hash swagger &> /dev/null; then echo "swagger(go-swagger) is required to make: https://goswagger.io/install.html > Docker image installation not working!"; exit 1; fi

init-template: 
ifneq '$(MODULE)' '{MODULE}'
	$(error template already initialized, module: $(MODULE))
endif
ifndef INIT
	$(error INIT variable is required to make.)
endif
	$(eval SED=$(if $(filter Darwin, $(shell uname -s)), sed -i '', sed -i))
	@find . -name "*.go" -maxdepth 3 -type f -print -exec $(SED) "s/{MODULE}/$(INIT)/g" {} +
	@$(SED) "s/{MODULE}/$(INIT)/g" .gitignore
	@$(SED) "s/{MODULE}/$(INIT)/g" .dockerignore
	@$(SED) "s/{MODULE}/$(INIT)/g" Dockerfile
	@$(SED) "s/{MODULE}/$(INIT)/g" go.mod
	@rm -rf .git
	git init
	@if [ "`git config --global user.name`" == "" ]; then git config --global user.name "`whoami`"; fi;
	@if [ "`git config --global user.email`" == "" ]; then git config --global user.email "`whoami`"; fi;
	git commit --allow-empty -m 'initial commit'
