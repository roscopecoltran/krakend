package main

import (
	"flag"
	"os"
	// 	"github.com/roscopecoltran/configor"
)

const (
	_sep = string(os.PathSeparator)
)

var (
	flagConfigPath     = flag.String("config", "", "Path to look for a config file. (directory)")
	port               = flag.Int("p", 0, "Port of the service")
	logLevel           = flag.String("l", "DEBUG", "Logging level")
	debug              = flag.Bool("d", true, "Enable the debug")
	configFile         = flag.String("c", "./conf.d/krakend/config.json", "Path to the configuration filename")
	etcdServer         = flag.String("etcd", "http://127.0.0.1:4001", "Comma-separated list of etcd servers (with port and schema)")
	defaultConfigFiles = []string{
		"shared/conf.d/krakend/debug.yaml",       // debugging
		"shared/conf.d/krakend/stores.yaml",      // datastores (bolt, gorm, dgraph, ...)
		"shared/conf.d/krakend/apis.yaml",        // apis (swagger, raml, ...)
		"shared/conf.d/krakend/proxy.yaml",       // proxies (morty, tor, ...)
		"shared/conf.d/krakend/credentials.yaml", // oauth2
		"shared/conf.d/krakend/schemas.yaml",     //
		"shared/conf.d/krakend/cloud.yaml",       // cloud deploymenet (docker, docker-compose, crane, k8s)
		"shared/conf.d/krakend/common.yaml",      //
		"shared/conf.d/krakend/flows.yaml",       // sub-requests (inserie, inparallel, flows)
		"shared/conf.d/krakend/tasks.yaml",       // local tasks (grep, fuzzy search,...)
		"shared/conf.d/krakend/frontend.yaml",    // admin web ui (react, material, vuejs)
		"shared/conf.d/krakend/backend.yaml",     // admin web ui (qor-admin, aor, ng-admin, vuejs-bulma)
		"shared/conf.d/krakend/environment.yaml", // env variables (oauth2, personal tokens, basic auth)
		"shared/conf.d/krakend/locales.yaml",     // translation, i18n, i10n
		"shared/conf.d/krakend/endpoints.yaml",   // krakend provider's endpoints (format: json, xpath, rss, plaintext)
		"shared/conf.d/krakend/providers.yaml"}   // providers profiles (swagger and data leafs mappings)
)

/*
func loadConfig() {
	defaultConfigFiles = append(defaultConfigFiles, *configFile)

	if err := configor.Load(&config.Config, defaultConfigFiles...); err != nil {
		log.Fatal("ERROR while loading the config struct:", err.Error())
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

}
*/

/*

- Notes:
	1. Switching between production and dev settings files to load:
	   - $ export CONFIGOR_ENV=prod # Will load `config.json`, `config.prod.json` if it exists `config.prod.json` will overwrite `config.json`'s configuration
*/
