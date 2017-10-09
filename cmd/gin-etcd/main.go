package main

import (
	"context"
	"flag"
	"log"
	"os"
	"strings"
	"time"

	"github.com/aviddiviner/gin-limit"
	"github.com/gin-gonic/contrib/cache"
	"github.com/gin-gonic/contrib/secure"
	"github.com/gin-gonic/gin"
	"gopkg.in/gin-contrib/cors.v1"

	// "github.com/gin-gonic/contrib/jwt"
	// "github.com/gregjones/httpcache"

	"github.com/roscopecoltran/krakend/config"
	"github.com/roscopecoltran/krakend/config/viper"
	"github.com/roscopecoltran/krakend/logging"
	"github.com/roscopecoltran/krakend/logging/gologging"
	"github.com/roscopecoltran/krakend/proxy"

	"github.com/roscopecoltran/krakend/sd/etcd"

	krakendgin "github.com/roscopecoltran/krakend/router/gin"
)

var (
	port               = flag.Int("p", 0, "Port of the service")
	logLevel           = flag.String("l", "ERROR", "Logging level")
	debug              = flag.Bool("d", true, "Enable the debug")
	configFile         = flag.String("c", "/etc/krakend/configuration.json", "Path to the configuration filename")
	etcdServer         = flag.String("etcd", "http://etcd-1:2379,http://etcd-2:2379,http://etcd-3:2379", "Comma-separated list of etcd servers (with port and schema)")
	limitMaxReqAllowed = flag.Int("max-req", 20, "Maximum allowed concurrent requests")
)

var (
	isCache     = flag.Bool("cache", true, "Enable the cache for the proxy gateway")
	isCacheHttp = flag.Bool("cache-http", false, "Enable the http-cache for responses")
	cacheDriver = flag.String("cache-driver", "inmemory", "Cache driver (available: inmemory, memecachec, redis)")
)

var (
	jwtSecret = flag.String("jwt-secret", "KrakenDrulez123.4567890!", "Secret for signing jwt")
	jwtIssuer = flag.String("jwt-issuer", "http://example.com/", "Issuer for the jwt")
	jwtPort   = flag.Int("jwt-port", 8090, "Port for the jwt generator api")
	jwsTTL    = flag.Duration("jwt-ttl", 1*time.Hour, "Expiration for the JWT")
)

var (
	allowedHosts   = flag.String("hosts", "127.0.0.1:8080,127.0.0.1:8085,127.0.0.1:8090,example.com,ssl.example.com", "Comma-separated list of allowed hosts")
	allowedOrigins = flag.String("cors-origins", "http://example.com", "Comma-separated list of CORS allowed origins")
	allowedMethods = flag.String("cors-methods", "HEAD,GET,POST,PUT,PATCH,DELETE", "Comma-separated list of CORS allowed methods")
	allowedHeaders = flag.String("cors-headers", "Origin,Authorization,Content-Type", "Comma-separated list of CORS allowed headers")
	exposedHeaders = flag.String("cors-headers-exposed", "Content-Length", "Comma-separated list of CORS exposed headers")
	corsTTL        = flag.Duration("cors-ttl", 12*time.Hour, "Max age for the CORS layer")
)

func main() {

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

	// etcd
	ctx, cancel := context.WithCancel(context.Background())

	etcdClient, err := etcd.NewClient(ctx, strings.Split(*etcdServer, ","), etcd.ClientOptions{})
	if err != nil {
		panic(err)
	}

	/*
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
	*/

	// cache
	store := cache.NewInMemoryStore(time.Minute)
	/*
		if *isCacheHttp {
			tp := httpcache.NewMemoryCacheTransport()
			client := http.Client{Transport: tp}
		}
	*/

	mws := []gin.HandlerFunc{
		secure.Secure(secure.Options{
			AllowedHosts:          strings.Split(*allowedHosts, ","),
			SSLRedirect:           false,
			SSLHost:               "ssl.example.com",
			SSLProxyHeaders:       map[string]string{"X-Forwarded-Proto": "https"},
			STSSeconds:            315360000,
			STSIncludeSubdomains:  true,
			FrameDeny:             true,
			ContentTypeNosniff:    true,
			BrowserXssFilter:      true,
			ContentSecurityPolicy: "default-src 'self'",
		}),
		limit.MaxAllowed(*limitMaxReqAllowed),
		cors.New(cors.Config{
			AllowOrigins:     strings.Split(*allowedOrigins, ","),
			AllowMethods:     strings.Split(*allowedMethods, ","),
			AllowHeaders:     strings.Split(*allowedHeaders, ","),
			ExposeHeaders:    strings.Split(*exposedHeaders, ","),
			AllowCredentials: true,
			MaxAge:           *corsTTL,
		}),
		//jwt.Auth(*jwtSecret),
	}

	// routerFactory := krakendgin.DefaultFactory(proxy.DefaultFactory(logger), logger)
	routerFactory := krakendgin.NewFactory(krakendgin.Config{
		Engine: gin.Default(),
		// ProxyFactory: customProxyFactory{logger, proxy.DefaultFactory(logger)},
		// ProxyFactory: proxy.NewDefaultFactory(proxy.HTTPProxyFactory(&client), logger),
		ProxyFactory: customProxyFactory{
			logger,
			proxy.DefaultFactoryWithSubscriber(logger, etcd.SubscriberFactory(ctx, etcdClient)),
		},
		Middlewares: mws,
		Logger:      logger,
		HandlerFactory: func(configuration *config.EndpointConfig, proxy proxy.Proxy) gin.HandlerFunc {
			return cache.CachePage(store, configuration.CacheTTL, krakendgin.EndpointHandler(configuration, proxy))
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
