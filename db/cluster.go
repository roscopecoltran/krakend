package db

import (
	// "github.com/aws/aws-sdk-go/aws"
	// "github.com/aws/aws-sdk-go/aws/credentials"
	// "github.com/aws/aws-sdk-go/aws/session"
	// "github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/jinzhu/gorm"

	// "gopkg.in/mgo.v2"

	"github.com/allegro/bigcache" // data-cache-bigcache
	"github.com/boltdb/bolt"
	"github.com/bradfitz/gomemcache/memcache" // db-kvs-memcache
	etcd "github.com/coreos/etcd/clientv3"    // db-kvs-etcd
	"github.com/garyburd/redigo/redis"        // db-kvs-redis

	"github.com/cayleygraph/cayley" // db-graph-cayley
	"github.com/jmcvetta/neoism"    // db-graph-neo4j
	// "github.com/ckaznocha/taggraph" // db-graph-taggraph

	"github.com/blevesearch/bleve"
	elastic "gopkg.in/olivere/elastic.v5" // data-index-search

	"github.com/nsqio/go-nsq"
	// "github.com/roscopecoltran/krakend/config"
)

/*
	refs:
	- https://github.com/DeadNumbers/goin/blob/master/bleve.go
	- https://github.com/supple/gorest/blob/master/storage/mongo.go
*/

type DataClusters struct {
	RDB struct {
		Gorm map[string]*gorm.DB
		// DynamoDB map[string]*dynamodb.DynamoDB
	}
	KVS struct {
		Redis    map[string]*redis.Conn
		Bolt     map[string]*bolt.DB
		Memcache map[string]*memcache.Client
		BigCache map[string]*bigcache.BigCache
		Etcd     map[string]*etcd.Client
	}
	//DOC struct {
	//	Mongo map[string]*mgo.Session
	//}
	IDX struct {
		Bleve   map[string]bleve.Index
		Elastic map[string]*elastic.Client
		//Sphinx  map[string]*sphinx.Client
	}
	GRA struct {
		Neo4J  map[string]*neoism.Database
		Cayley map[string]*cayley.Handle
		//TaggGraph   map[string]*taggraph.TagGrapher
	}
	MQ struct {
		NsqProducer map[string]*nsq.Producer
	}
}

func (cluster *DataClusters) initGatewayCluster() {

	// Relational Databases
	cluster.RDB.Gorm = make(map[string]*gorm.DB)

	/*
		dynamoDB := dynamodb.New(session.New(
			&aws.Config{
				Region:      aws.String(c.GetString("db.region")),
				Credentials: credentials.NewEnvCredentials(),
				Endpoint:    aws.String(c.GetString("db.endpoint")),
				DisableSSL:  aws.Bool(c.GetBool("db.disable_ssl")),
			}
		))
		cluster.RDB.Dynamodb = make(map[string]*gorm.DB)
	*/

	// Key-Value Stores
	cluster.KVS.Bolt = make(map[string]*bolt.DB)
	cluster.KVS.Redis = make(map[string]*redis.Conn)
	cluster.KVS.Memcache = make(map[string]*memcache.Client)
	cluster.KVS.BigCache = make(map[string]*bigcache.BigCache)
	cluster.KVS.Etcd = make(map[string]*etcd.Client)

	// cluster.DOC.Mongo = make(map[string]*mgo.Session)

	// Full-Text Indexes
	cluster.IDX.Bleve = make(map[string]bleve.Index)
	cluster.IDX.Elastic = make(map[string]*elastic.Client)
	// cluster.IDX.Sphinx = make(map[string]*sphinx.Client)

	// Graph Databases
	cluster.GRA.Neo4J = make(map[string]*neoism.Database)
	cluster.GRA.Cayley = make(map[string]*cayley.Handle)
	// cluster.GRAPHS.TaggGraph = make(map[string]*taggraph.TagGrapher)

}
