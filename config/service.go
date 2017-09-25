package config

import (
	"time"

	"github.com/jinzhu/gorm"
)

// example: https://github.com/devopsfaith/krakend/blob/master/docs/CONFIG.md#json-example

// ServiceConfig defines the krakend service
type ServiceConfig struct {
	gorm.Model `json:"-" yaml:"-"`
	// set of endpoint definitions
	Endpoints []*EndpointConfig `gorm:"-" mapstructure:"endpoints" json:"endpoints" yaml:"endpoints"`
	// defafult timeout
	Timeout time.Duration `gorm:"column:timeout" mapstructure:"timeout" json:"timeout" yaml:"timeout"`
	// default TTL for GET
	CacheTTL time.Duration `gorm:"column:cache_ttl" mapstructure:"cache_ttl" json:"cache_ttl" yaml:"cache_ttl"`
	// default set of hosts
	Host []string `gorm:"-" mapstructure:"host" json:"host" yaml:"host"` // many2many:service_hosts;
	// port to bind the krakend service
	Port int `gorm:"column:port" mapstructure:"port" json:"port" yaml:"port"`
	// version code of the configuration
	Version int `gorm:"column:version" mapstructure:"version" json:"version" yaml:"version"`
	// run krakend in debug mode
	Debug     bool      `gorm:"column:debug" mapstructure:"debug" json:"debug" yaml:"debug"`
	uriParser URIParser `gorm:"-" json:"-" yaml:"-"`
}
