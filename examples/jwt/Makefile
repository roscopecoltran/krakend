.PHONY: all deps build

# This Makefile is a simple example that demonstrates usual steps to build a binary that can be run in the same
# architecture that was compiled in. The "ldflags" in the build assure that any needed dependency is included in the
# binary and no external dependencies are needed to run the service.

KRAKEND_VERSION=$(shell git describe --always --long --dirty --tags)
# BIN_NAME=krakend-jwt-l${KRAKEND_VERSION}

all: deps build

deps:
	go get "github.com/roscopecoltran/krakend/config/viper"
	go get "github.com/roscopecoltran/krakend/proxy"
	go get "github.com/roscopecoltran/krakend/router/gin"
	go get "github.com/roscopecoltran/krakend/logging/gologging"
	go get "github.com/gin-gonic/contrib/secure"
	go get "github.com/aviddiviner/gin-limit"
	go get "gopkg.in/gin-contrib/cors.v1"
	go get "github.com/dgrijalva/jwt-go"
	go get "github.com/gin-gonic/contrib/jwt"

build:
	@gox -verbose -os="darwin linux" -arch="amd64" -output="krakend-jwt-{{.OS}}" .
	# go build -ldflags="-X github.com/roscopecoltran/krakend/core.KrakendVersion=${KRAKEND_VERSION}" -o ${BIN_NAME}
	@echo "You can now use ./${BIN_NAME}"

