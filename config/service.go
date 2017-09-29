package config

import (
	"time"
	//"github.com/jinzhu/gorm"
)

// example: https://github.com/devopsfaith/krakend/blob/master/docs/CONFIG.md#json-example

// ServiceConfig defines the krakend service
/*
type ServiceConfig struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`
	// set of endpoint definitions
	Endpoints []*EndpointConfig `gorm:"-" mapstructure:"endpoints" json:"endpoints" yaml:"endpoints" toml:"endpoints"`
	// defafult timeout
	Timeout time.Duration `gorm:"column:timeout" default:"3000" mapstructure:"timeout" json:"timeout" yaml:"timeout" toml:"timeout"`
	// default TTL for GET
	CacheTTL time.Duration `gorm:"column:cache_ttl" mapstructure:"cache_ttl" json:"cache_ttl" yaml:"cache_ttl" toml:"cache_ttl"`
	// default set of hosts
	Host []string `gorm:"-" mapstructure:"host" json:"host" yaml:"host" toml:"host"` // many2many:service_hosts;
	// port to bind the krakend service
	Port int `gorm:"column:port" mapstructure:"port" json:"port" yaml:"port" toml:"port"`
	// version code of the configuration
	Version int `gorm:"column:version" mapstructure:"version" json:"version" yaml:"version" toml:"version"`
	// run krakend in debug mode
	Debug     bool      `gorm:"column:debug" mapstructure:"debug" json:"debug" yaml:"debug" toml:"debug"`
	uriParser URIParser `gorm:"-" json:"-" yaml:"-" toml:"-"`
}
*/

// ServiceConfig defines the krakend service
type ServiceConfig struct {
	// set of endpoint definitions
	Endpoints []*EndpointConfig `mapstructure:"endpoints"`
	// Processors configuration for handling data received or cached
	Processors ProcessorConfig `mapstructure:"processors"`
	// Tasks/Queue order to fetch and aggregate data
	TaskQ []TaskQConfig `mapstructure:"taskq"`
	// defafult timeout
	Timeout time.Duration `mapstructure:"timeout"`
	// default TTL for GET
	CacheTTL time.Duration `mapstructure:"cache_ttl"`
	// default set of hosts
	Host []string `mapstructure:"host"`
	// port to bind the krakend service
	Port int `mapstructure:"port"`
	// version code of the configuration
	Version int `mapstructure:"version"`
	// run krakend in debug mode
	Debug bool `mapstructure:"debug"`
	// run krakend in debug mode
	Admin bool `mapstructure:"admin"`
	// URI Parsing object
	uriParser URIParser
}
