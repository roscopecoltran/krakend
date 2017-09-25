package config

import (
	"time"

	"github.com/jinzhu/gorm"
)

// EndpointConfig defines the configuration of a single endpoint to be exposed
// by the krakend service
type EndpointConfig struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`
	// url pattern to be registered and exposed to the world
	Endpoint string `gorm:"column:endpoint" mapstructure:"endpoint" json:"endpoint" yaml:"endpoint" toml:"endpoint"`
	// HTTP method of the endpoint (GET, POST, PUT, etc)
	Method string `gorm:"column:method" mapstructure:"method" json:"method" yaml:"method" toml:"method"`
	// HTTP headers of the endpoint to forward
	Header []string `gorm:"-" mapstructure:"header" json:"header" yaml:"header" toml:"header"` // many2many:endpoints_config;
	// set of definitions of the backends to be linked to this endpoint
	Backend []*Backend `gorm:"-" mapstructure:"backend" json:"backend"  yaml:"backend" toml:"backend"`
	// number of concurrent calls this endpoint must send to the backends
	ConcurrentCalls int `gorm:"column:concurrent_calls" mapstructure:"concurrent_calls" json:"concurrent_calls"  yaml:"concurrent_calls" toml:"concurrent_calls"`
	// timeout of this endpoint
	Timeout time.Duration `gorm:"column:timeout" mapstructure:"timeout" json:"timeout"  yaml:"timeout" toml:"timeout"`
	// duration of the cache header
	CacheTTL time.Duration `gorm:"column:cache_ttl" mapstructure:"cache_ttl" json:"cache_ttl"  yaml:"cache_ttl" toml:"cache_ttl"`
	// list of query string params to be extracted from the URI
	QueryStrings []string `gorm:"-" mapstructure:"query_strings" json:"query_strings"  yaml:"query_strings" toml:"query_strings"` // many2many:endpoints_query_string;
	// Endpoint Extra configuration for customized behaviour
	ExtraConfig ExtraConfig `gorm:"-" mapstructure:"extra_config" json:"-"  yaml:"-" toml:"-"` // many2many:endpoints_extra_config;
}
