NAME=market_back
PKG=market_back/cmd
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
VERSION=git-$(subst /,-,$(BRANCH))-$(shell date +%Y%m%d%H)-$(shell git describe --always --dirty)
IMAGE_TAG=$(VERSION)
IMAGE_REPO=*
LOCAL_REPO=*
OutDir=proto


linux:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
		go build -v -a -tags netgo -installsuffix netgo -installsuffix cgo -ldflags '-w -s' -ldflags "-X main.Version=$(VERSION)" \
		-o ./build/linux/market_back $(PKG)

darwin:
	GOOS=darwin GOARCH=amd64 \
		go build -a -tags netgo -installsuffix netgo -ldflags "-X main.Version=$(VERSION)" \
		-o ./build/darwin/market_back $(PKG)

m1_osx:
	GOOS=darwin GOARCH=arm64 \
		go build -a -tags netgo -installsuffix netgo -ldflags "-X main.Version=$(VERSION)" \
		-o ./build/darwin/market_back $(PKG)

dist:
	env GOOS=linux GOARCH=amd64 go build -tags=jsoniter -v -o ./build/linux/market_back

build: linux docs

push: linux docs
	docker build -t ${IMAGE_REPO}/market_back:${IMAGE_TAG} .

push_aarch64:
	docker build -f Dockerfile.aarch64 -t ${IMAGE_REPO}/aarch64/market_back:${IMAGE_TAG} .

push_dev: linux docs
	docker build -t ${LOCAL_REPO}/market_back:${IMAGE_TAG} .

docs:
	swag fmt
	swag init --pd -g ./swagger.go -o ./apidocs/

gen:
	mkdir -p ./internal/${OutDir}
	protoc --go_out=./internal/${OutDir} --go-grpc_out=./internal/${OutDir} ./proto/*.proto

.PHONY: linux darwin m1_osx dist push push_aarch64 push_dev docs gen build