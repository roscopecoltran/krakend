package main

import (
	"flag"
	"log"
	"os"

	gorilla "github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/urfave/negroni"
	"github.com/zbindenren/negroni-prometheus"
	"gopkg.in/unrolled/secure.v1"

	"github.com/roscopecoltran/krakend/config/viper"
	"github.com/roscopecoltran/krakend/logging/gologging"
	"github.com/roscopecoltran/krakend/proxy"
	"github.com/roscopecoltran/krakend/router/mux"
	krakendnegroni "github.com/roscopecoltran/krakend/router/negroni"
)

func main() {
	port := flag.Int("p", 0, "Port of the service")
	logLevel := flag.String("l", "ERROR", "Logging level")
	debug := flag.Bool("d", false, "Enable the debug")
	configFile := flag.String("c", "./conf.d/krakend/configuration.json", "Path to the configuration filename")
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

	secureMiddleware := secure.New(secure.Options{
		AllowedHosts:          []string{"127.0.0.1:8080", "example.com", "ssl.example.com"},
		SSLRedirect:           false,
		SSLHost:               "ssl.example.com",
		SSLProxyHeaders:       map[string]string{"X-Forwarded-Proto": "https"},
		STSSeconds:            315360000,
		STSIncludeSubdomains:  true,
		STSPreload:            true,
		FrameDeny:             true,
		ContentTypeNosniff:    true,
		BrowserXssFilter:      true,
		ContentSecurityPolicy: "default-src 'self'",
	})

	r := gorilla.NewRouter()
	m := negroniprometheus.NewMiddleware("serviceName")
	r.Handle("/__metrics", prometheus.Handler())

	cfg := krakendnegroni.DefaultConfigWithRouter(proxy.DefaultFactory(logger), logger, r, []negroni.Handler{m})
	cfg.Middlewares = []mux.HandlerMiddleware{secureMiddleware}

	mux.NewFactory(cfg).New().Run(serviceConfig)
}
