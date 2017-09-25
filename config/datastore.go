package config

import (
	"github.com/jinzhu/gorm"
)

type DatabaseConfig struct {
	gorm.Model `json:"-" yaml:"-"`
	RDB        RdbConfig `json:"database" yaml:"database"`
	KVS        KvsConfig `json:"kvstore" yaml:"kvstore"`
	IDX        IdxConfig `json:"fulltext" yaml:"fulltext"`
	GRA        GraConfig `json:"graphdb" yaml:"graphdb"`
}

type DatastoreConfig struct {
	gorm.Model     `json:"-" yaml:"-"`
	Driver         string `json:"driver" yaml:"driver"`
	Adapter        string `default:"sqlite" json:"adapter"  yaml:"adapter"`
	Disabled       bool   `default:"false" json:"disabled" yaml:"disabled"`
	Debug          bool   `default:"true" json:"debug" yaml:"debug"`
	Name           string `json:"name"  yaml:"name"`
	Host           string `default:"localhost" json:"host"  yaml:"host"`
	Port           string `json:"port"  yaml:"port"`
	User           string `json:"user"  yaml:"user"`
	Password       string `json:"password"  yaml:"password"`
	SSL            bool   `default:"false" json:"ssl_mode" yaml:"ssl_mode"`
	CACertificates string `json:"ca_certificates"  yaml:"ca_certificates"`
}

type KvsConfig struct {
	gorm.Model `json:"-" yaml:"-"`
	Adapter    string `env:"KVS_ADAPTER" default:"boltdb" json:"adapter"  yaml:"adapter"` // boltdb, etcd, redis, memcache, bigcache
	Disabled   bool   `default:"false" env:"KVS_DISABLED" json:"disabled"  yaml:"disabled"`
	Debug      bool   `default:"true" env:"KVS_DEBUG" json:"debug"  yaml:"debug"`
	Name       string `env:"KVS_INDEX_NAME" default:"krakend" json:"name"  yaml:"name"`
	PrefixPath string `env:"KVS_INDEX_PREFIX_PATH" default:"./data/bucket"  json:"prefix_path"  yaml:"prefix_path"`
	Hosts      string `env:"KVS_HOSTS" default:"localhost" json:"host"  yaml:"host"`
}

type RdbConfig struct {
	gorm.Model   `json:"-" yaml:"-"`
	Adapter      string `env:"RDB_ADAPTER" default:"sqlite" json:"adapter"  yaml:"adapter"` // mysql, postgres, sqlite3
	Disabled     bool   `default:"false" env:"RDB_DISABLED" json:"disabled"  yaml:"disabled"`
	Debug        bool   `default:"true" env:"RDB_DEBUG" json:"debug"  yaml:"debug"`
	PrefixPath   string `env:"RDB_PREFIX_PATH" default:"./data/db" json:"prefix_path"  yaml:"prefix_path"`
	Name         string `env:"RDB_NAME" default:"sniper" json:"name"  yaml:"name"`
	Host         string `env:"RDB_HOST" default:"localhost" json:"host"  yaml:"host"`
	Port         string `env:"RDB_PORT" default:"3306" json:"port"  yaml:"port"`
	User         string `env:"RDB_USER" json:"user"  yaml:"user"`
	Password     string `env:"RDB_PASSWORD" json:"password"  yaml:"password"`
	MaxIdleConns int    `default:"2" env:"RDB_MAX_IDLE_CONNECTIONS" json:"max_idle_connections"  yaml:"max_idle_connections"`
	MaxOpenConns int    `default:"4" env:"RDB_MAX_OPEN_CONNECTIONS" json:"max_open_connections"  yaml:"max_open_connections"`
}

type IdxConfig struct {
	gorm.Model `json:"-" yaml:"-"`
	Adapter    string `env:"FULLTEXT_IDX_ADAPTER" default:"elastic" json:"adapter"  yaml:"adapter"` // bleve, elastic, sphinx
	Disabled   bool   `default:"false" env:"FULLTEXT_IDX_DISABLED" json:"disabled"  yaml:"disabled"`
	Debug      bool   `default:"true" env:"FULLTEXT_IDX_DEBUG" json:"debug"  yaml:"debug"`
	Name       string `env:"FULLTEXT_IDX_NAME" default:"sniper" json:"name"  yaml:"name"`
	Host       string `env:"FULLTEXT_IDX_HOST" default:"localhost" json:"host"  yaml:"host"`
	Port       string `env:"FULLTEXT_IDX_PORT" default:"9200" json:"port"  yaml:"port"`
	User       string `env:"FULLTEXT_IDX_USER" json:"user"  yaml:"user"`
	Password   string `env:"FULLTEXT_IDX_PASSWORD" json:"password"  yaml:"password"`
}

type GraConfig struct {
	gorm.Model `json:"-" yaml:"-"`
	Adapter    string `env:"GRAPHDB_ADAPTER" default:"neo4j" json:"adapter"  yaml:"adapter"` // meo4j, dgraph, cayley
	Disabled   bool   `default:"false" env:"GRAPHDB_DISABLED" json:"disabled"  yaml:"disabled"`
	Debug      bool   `default:"true" env:"GRAPHDB_DEBUG" json:"debug"  yaml:"debug"`
	Name       string `env:"GRAPHDB_NAME" default:"sniper" json:"name"  yaml:"name"`
	Host       string `env:"GRAPHDB_HOST" default:"localhost" json:"host"  yaml:"host"`
	Port       string `env:"GRAPHDB_PORT" default:"7474" json:"port"  yaml:"port"`
	User       string `env:"GRAPHDB_USER" json:"user"  yaml:"user"`
	Password   string `env:"GRAPHDB_PASSWORD" json:"password"  yaml:"password"`
}