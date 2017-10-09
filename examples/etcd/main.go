package main

import (
	"context"
	"flag"
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/roscopecoltran/krakend/config"
	"github.com/roscopecoltran/krakend/config/viper"
	"github.com/roscopecoltran/krakend/logging"
	"github.com/roscopecoltran/krakend/logging/gologging"
	"github.com/roscopecoltran/krakend/proxy"
	krakendgin "github.com/roscopecoltran/krakend/router/gin"
	"github.com/roscopecoltran/krakend/sd/etcd"
)

func main() {
	port := flag.Int("p", 0, "Port of the service")
	logLevel := flag.String("l", "ERROR", "Logging level")
	debug := flag.Bool("d", trye, "Enable the debug")
	configFile := flag.String("c", "/etc/krakend/configuration.json", "Path to the configuration filename")
	etcdServer := flag.String("etcd", "http://etcd-1:2379,http://etcd-2:2379,http://etcd-3:2379", "Comma-separated list of etcd servers (with port and schema)")
	flag.Parse()

	parser := viper.New()
	serviceConfig, err := parser.Parse(*configFile)
	if err != nil {
		log.Fatal("ERROR:", err.Error())
	}
	serviceConfig.Debug = serviceConfig.Debug || *debug
	if *port != 0 {
		serviceConfig.Port = *port
	}

	logger, err := gologging.NewLogger(*logLevel, os.Stdout, "[KRAKEND]")
	if err != nil {
		log.Fatal("ERROR:", err.Error())
	}

	ctx, cancel := context.WithCancel(context.Background())

	etcdClient, err := etcd.NewClient(ctx, strings.Split(*etcdServer, ","), etcd.ClientOptions{})
	if err != nil {
		panic(err)
	}

	routerFactory := krakendgin.NewFactory(krakendgin.Config{
		Engine:         gin.Default(),
		Middlewares:    []gin.HandlerFunc{},
		Logger:         logger,
		HandlerFactory: krakendgin.EndpointHandler,
		ProxyFactory: customProxyFactory{
			logger,
			proxy.DefaultFactoryWithSubscriber(logger, etcd.SubscriberFactory(ctx, etcdClient)),
		},
	})

	routerFactory.New().Run(serviceConfig)

	cancel()
}

// customProxyFactory adds a logging middleware wrapping the internal factory
type customProxyFactory struct {
	logger  logging.Logger
	factory proxy.Factory
}

// New implements the Factory interface
func (cf customProxyFactory) New(cfg *config.EndpointConfig) (p proxy.Proxy, err error) {
	p, err = cf.factory.New(cfg)
	if err == nil {
		p = proxy.NewLoggingMiddleware(cf.logger, cfg.Endpoint)(p)
	}
	return
}
