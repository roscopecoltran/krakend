.PHONY: all deps build

# This Makefile is a simple example that demonstrates usual steps to build a binary that can be run in the same
# architecture that was compiled in. The "ldflags" in the build assure that any needed dependency is included in the
# binary and no external dependencies are needed to run the service.

KRAKEND_VERSION=$(shell git describe --always --long --dirty --tags)
BIN_NAME=krakend_gin_${KRAKEND_VERSION}

all: deps build

deps:
	go get -v "github.com/mattn/goreman"
	go get -v "github.com/coreos/etcd"
	go get -v "github.com/roscopecoltran/krakend/sd/etcd"
	go get -v "github.com/roscopecoltran/krakend/config/viper"
	go get -v "github.com/roscopecoltran/krakend/proxy"
	go get -v "github.com/roscopecoltran/krakend/router/gin"
	go get -v "github.com/roscopecoltran/krakend/logging/gologging"
	go get -v "github.com/gin-contrib/cache"
	go get -v "github.com/gin-contrib/secure"
	go get -v "github.com/aviddiviner/gin-limit"

build:
	go build -a -ldflags="-X github.com/roscopecoltran/krakend/core.KrakendVersion=${KRAKEND_VERSION}" -o ${BIN_NAME}
	@echo "You can now use ./${BIN_NAME}"
