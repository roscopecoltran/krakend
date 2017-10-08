package main

import (
	"context"
	"flag"
	"log"
	"os"
	"strings"
	"time"

	"github.com/aviddiviner/gin-limit"
	"github.com/gin-contrib/secure"
	"github.com/gin-gonic/contrib/cache"
	"github.com/gin-gonic/gin"
	"github.com/roscopecoltran/configor"

	"github.com/roscopecoltran/krakend/config"
	"github.com/roscopecoltran/krakend/config/viper"
	"github.com/roscopecoltran/krakend/db"
	"github.com/roscopecoltran/krakend/logging"
	"github.com/roscopecoltran/krakend/logging/gologging"
	"github.com/roscopecoltran/krakend/proxy"
	krakendgin "github.com/roscopecoltran/krakend/router/gin"
	"github.com/roscopecoltran/krakend/sd/etcd"
	/*
		// "github.com/roscopecoltran/krakend/db/model"
		// "github.com/gin-contrib/static"
		// "github.com/gregjones/httpcache"
		// "github.com/gin-contrib/cache"
		// "github.com/gin-contrib/cache/persistence"
		// "github.com/k0kubun/pp"
	*/)

func main() {

	flag.Parse()

	defaultConfigFiles = append(defaultConfigFiles, *configFile)

	if err := configor.Load(&config.Config, defaultConfigFiles...); err != nil {
		log.Fatal("ERROR while loading the config struct:", err.Error())
	}

	if err := db.NewStack(); err != nil {
		log.Fatal("ERROR:", err.Error())
	}

	parser := viper.New()
	serviceConfig, err := parser.Parse(*configFile)
	if err != nil {
		log.Fatal("ERROR:", err.Error())
	}
	serviceConfig.Debug = serviceConfig.Debug || *debug
	if *port != 0 {
		serviceConfig.Port = *port
	}

	logger, err := gologging.NewLogger(*logLevel, os.Stdout, "[SNIPER]")
	if err != nil {
		log.Fatal("ERROR:", err.Error())
	}

	// pp.Println(config.Config)
	if err := db.MigrateDB(false, false, false, config.DefaultGatewayModels...); err != nil {
		//logging.Fatal("Can not load default models into the RDB store.")
		log.Fatal("ERROR:", err.Error())
	}

	ctx, cancel := context.WithCancel(context.Background())

	etcdClient, err := etcd.NewClient(ctx, strings.Split(*etcdServer, ","), etcd.ClientOptions{})
	if err != nil {
		panic(err)
	}

	store := cache.NewInMemoryStore(time.Minute)
	// tp := httpcache.NewMemoryCacheTransport()
	// client := http.Client{Transport: tp}

	// pp.Println(store)

	// https://github.com/gin-contrib/secure/blob/master/example/code1/example.go
	mws := []gin.HandlerFunc{
		secure.New(secure.Config{
			AllowedHosts: []string{"127.0.0.1:8080", "localhost:8080", "example.com", "ssl.example.com"},
			SSLRedirect:  false,
			SSLHost:      "ssl.example.com",
			// SSLProxyHeaders:       map[string]string{"X-Forwarded-Proto": "https"},
			STSSeconds:            315360000,
			STSIncludeSubdomains:  true,
			FrameDeny:             true,
			ContentTypeNosniff:    true,
			BrowserXssFilter:      true,
			ContentSecurityPolicy: "default-src 'self'",
		}),
		limit.MaxAllowed(20),
	}

	// routerFactory := krakendgin.DefaultFactory(proxy.DefaultFactory(logger), logger)

	routerFactory := krakendgin.NewFactory(krakendgin.Config{
		Engine: gin.Default(),
		/*
			ProxyFactory: proxy.NewDefaultFactory( // default
				proxy.HTTPProxyFactory(&client),
				logger,
			),
		*/
		/*
			ProxyFactory: customProxyFactory{ // gin (ref. https://github.com/devopsfaith/krakend/blob/master/examples/gin/main.go#L66-L71)
				logger,
				proxy.DefaultFactory(logger),
			},
		*/
		ProxyFactory: customProxyFactory{
			logger,
			proxy.DefaultFactoryWithSubscriber( // etcd (ref. https://github.com/devopsfaith/krakend/blob/master/examples/etcd/main.go)
				logger,
				etcd.SubscriberFactory(ctx, etcdClient),
			),
		},
		Middlewares: mws,
		Logger:      logger,
		HandlerFactory: func(configuration *config.EndpointConfig, proxy proxy.Proxy) gin.HandlerFunc {
			return cache.CachePage(store, configuration.CacheTTL, krakendgin.EndpointHandler(configuration, proxy))
		},
	})

	if err := configor.Dump(config.Config, "all", "yaml,toml,json", "./dump"); err != nil {
		log.Fatal("ERROR while dumping the config struct:", err.Error())
	}

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
