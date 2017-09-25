package config

import (
	"github.com/jinzhu/gorm"
)

type StoreConfig struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`
	RDB        StoreRDBConfig `json:"database" yaml:"database" toml:"database"`
	KVS        StoreKVSConfig `json:"kvstore" yaml:"kvstore" toml:"kvstore"`
	IDX        StoreIDXConfig `json:"fulltext" yaml:"fulltext" toml:"fulltext"`
	GRA        StoreGRAConfig `json:"graphdb" yaml:"graphdb" toml:"graphdb"`
}

type DatabaseConfig struct {
	gorm.Model     `json:"-" yaml:"-" toml:"-"`
	Driver         string `json:"driver" yaml:"driver" toml:"driver"`
	Adapter        string `default:"sqlite" json:"adapter" yaml:"adapter" toml:"adapter"`
	Disabled       bool   `default:"false" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug          bool   `default:"true" json:"debug" yaml:"debug" toml:"debug"`
	Name           string `json:"name" yaml:"name" toml:"name"`
	Host           string `default:"localhost" json:"host" yaml:"host" toml:"host"`
	Port           string `json:"port" yaml:"port" toml:"port"`
	User           string `json:"user" yaml:"user" toml:"user"`
	Password       string `json:"password" yaml:"password" toml:"password"`
	SSL            bool   `default:"false" json:"ssl_mode" yaml:"ssl_mode" toml:"ssl_mode"`
	CACertificates string `json:"ca_certificates" yaml:"ca_certificates" toml:"ca_certificates"`
}

type StoreRDBConfig struct {
	gorm.Model   `json:"-" yaml:"-" toml:"-"`
	Adapter      string `env:"RDB_ADAPTER" default:"sqlite" json:"adapter" yaml:"adapter" toml:"adapter"` // mysql, postgres, sqlite3
	Disabled     bool   `default:"false" env:"RDB_DISABLED" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug        bool   `default:"true" env:"RDB_DEBUG" json:"debug" yaml:"debug" toml:"debug"`
	PrefixPath   string `env:"RDB_PREFIX_PATH" default:"./shared/data/rdb" json:"prefix_path" yaml:"prefix_path" toml:"prefix_path"`
	Name         string `env:"RDB_NAME" default:"sniper" json:"name" yaml:"name" toml:"name"`
	Host         string `env:"RDB_HOST" default:"localhost" json:"host" yaml:"host" toml:"host"`
	Port         string `env:"RDB_PORT" default:"3306" json:"port" yaml:"port" toml:"port"`
	User         string `env:"RDB_USER" json:"user" yaml:"user" toml:"user"`
	Password     string `env:"RDB_PASSWORD" json:"password" yaml:"password" toml:"password"`
	MaxIdleConns int    `default:"2" env:"RDB_MAX_IDLE_CONNECTIONS" json:"max_idle_connections" yaml:"max_idle_connections" toml:"max_idle_connections"`
	MaxOpenConns int    `default:"4" env:"RDB_MAX_OPEN_CONNECTIONS" json:"max_open_connections" yaml:"max_open_connections" toml:"max_open_connections"`
}

type StoreIDXConfig struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`
	Adapter    string `env:"FULLTEXT_IDX_ADAPTER" default:"elastic" json:"adapter" yaml:"adapter" toml:"adapter"` // bleve, elastic, sphinx
	Disabled   bool   `default:"false" env:"FULLTEXT_IDX_DISABLED" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug      bool   `default:"true" env:"FULLTEXT_IDX_DEBUG" json:"debug" yaml:"debug" toml:"debug"`
	Name       string `env:"FULLTEXT_IDX_NAME" default:"sniper" json:"name" yaml:"name" toml:"name"`
	Host       string `env:"FULLTEXT_IDX_HOST" default:"localhost" json:"host" yaml:"host" toml:"host"`
	Port       string `env:"FULLTEXT_IDX_PORT" default:"9200" json:"port" yaml:"port" toml:"port"`
	User       string `env:"FULLTEXT_IDX_USER" json:"user" yaml:"user" toml:"user"`
	Password   string `env:"FULLTEXT_IDX_PASSWORD" json:"password" yaml:"password" toml:"password"`
}

type StoreKVSConfig struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`
	Adapter    string `env:"KVS_ADAPTER" default:"boltdb" json:"adapter" yaml:"adapter" toml:"adapter"` // boltdb, etcd, redis, memcache, bigcache
	Disabled   bool   `default:"false" env:"KVS_DISABLED" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug      bool   `default:"true" env:"KVS_DEBUG" json:"debug" yaml:"debug" toml:"debug"`
	Name       string `env:"KVS_INDEX_NAME" default:"krakend" json:"name" yaml:"name" toml:"name"`
	PrefixPath string `env:"KVS_INDEX_PREFIX_PATH" default:"./shared/data/bucket"  json:"prefix_path" yaml:"prefix_path" toml:"prefix_path"`
	Hosts      string `env:"KVS_HOSTS" default:"localhost" json:"host" yaml:"host" toml:"host"`
}

type StoreGRAConfig struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`
	Adapter    string `env:"GRAPHDB_ADAPTER" default:"neo4j" json:"adapter" yaml:"adapter" toml:"adapter"` // meo4j, dgraph, cayley
	Disabled   bool   `default:"false" env:"GRAPHDB_DISABLED" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug      bool   `default:"true" env:"GRAPHDB_DEBUG" json:"debug" yaml:"debug" toml:"debug"`
	Name       string `env:"GRAPHDB_NAME" default:"sniper" json:"name" yaml:"name" toml:"name"`
	Host       string `env:"GRAPHDB_HOST" default:"localhost" json:"host" yaml:"host" toml:"host"`
	Port       string `env:"GRAPHDB_PORT" default:"7474" json:"port" yaml:"port" toml:"port"`
	User       string `env:"GRAPHDB_USER" json:"user" yaml:"user" toml:"user"`
	Password   string `env:"GRAPHDB_PASSWORD" json:"password" yaml:"password" toml:"password"`
}
