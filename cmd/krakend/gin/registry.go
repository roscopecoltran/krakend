package main

/*

import (
	"log"


	"github.com/roscopecoltran/krakend/sd/etcd"
	"github.com/roscopecoltran/krakend/config"
	"github.com/roscopecoltran/krakend/config/viper"
	"github.com/roscopecoltran/krakend/db"
)

func LoadRegistry() {

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

}

*/
