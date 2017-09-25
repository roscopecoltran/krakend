package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/aviddiviner/gin-limit"
	"github.com/gin-gonic/contrib/cache"
	"github.com/gin-gonic/contrib/secure"
	"github.com/gin-gonic/gin"
	// "github.com/gin-gonic/contrib/static"

	// "github.com/k0kubun/pp"
	"github.com/roscopecoltran/configor"

	"github.com/roscopecoltran/krakend/config"
	"github.com/roscopecoltran/krakend/config/viper"
	"github.com/roscopecoltran/krakend/db"
	//"github.com/roscopecoltran/krakend/db/model"
	"github.com/roscopecoltran/krakend/logging"
	"github.com/roscopecoltran/krakend/logging/gologging"
	"github.com/roscopecoltran/krakend/proxy"
	krakendgin "github.com/roscopecoltran/krakend/router/gin"
	// plug_openapi "github.com/roscopecoltran/krakend/plugins/openapi"
)

var defaultConfigFiles = []string{
	"shared/conf.d/krakend/stores.yaml",
	"shared/conf.d/krakend/apis.yaml",
	"shared/conf.d/krakend/proxy.yaml",
	"shared/conf.d/krakend/credentials.yaml",
	"shared/conf.d/krakend/schemas.yaml",
	"shared/conf.d/krakend/cloud.yaml",
	"shared/conf.d/krakend/common.yaml",
	"shared/conf.d/krakend/flows.yaml",
	"shared/conf.d/krakend/tasks.yaml",
	"shared/conf.d/krakend/frontend.yaml",
	"shared/conf.d/krakend/backend.yaml",
	"shared/conf.d/krakend/environment.yaml",
	"shared/conf.d/krakend/locales.yaml",
	"shared/conf.d/krakend/endpoints.yaml",
	"shared/conf.d/krakend/providers.yaml"}

func main() {
	port := flag.Int("p", 0, "Port of the service")
	logLevel := flag.String("l", "DEBUG", "Logging level")
	debug := flag.Bool("d", true, "Enable the debug")
	configFile := flag.String("c", "./conf.d/krakend/config.json", "Path to the configuration filename")
	flag.Parse()

	defaultConfigFiles = append(defaultConfigFiles, *configFile)

	// $ export CONFIGOR_ENV=prod # Will load `config.json`, `config.prod.json` if it exists `config.prod.json` will overwrite `config.json`'s configuration
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

	store := cache.NewInMemoryStore(time.Minute)

	// pp.Println(store)

	mws := []gin.HandlerFunc{
		secure.Secure(secure.Options{
			AllowedHosts:          []string{"127.0.0.1:8080", "example.com", "ssl.example.com"},
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
		limit.MaxAllowed(20),
	}

	// routerFactory := krakendgin.DefaultFactory(proxy.DefaultFactory(logger), logger)

	routerFactory := krakendgin.NewFactory(krakendgin.Config{
		Engine:       gin.Default(),
		ProxyFactory: customProxyFactory{logger, proxy.DefaultFactory(logger)},
		Middlewares:  mws,
		Logger:       logger,
		HandlerFactory: func(configuration *config.EndpointConfig, proxy proxy.Proxy) gin.HandlerFunc {
			return cache.CachePage(store, configuration.CacheTTL, krakendgin.EndpointHandler(configuration, proxy))
		},
	})

	if err := configor.Dump(config.Config, "all", "yaml,toml,json", "./dump"); err != nil {
		log.Fatal("ERROR while dumping the config struct:", err.Error())
	}

	routerFactory.New().Run(serviceConfig)
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
