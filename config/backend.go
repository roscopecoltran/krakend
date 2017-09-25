package config

import (
	"time"

	"github.com/jinzhu/gorm"

	"github.com/roscopecoltran/krakend/encoding"
)

// Backend defines how krakend should connect to the backend service (the API resource to consume)
// and how it should process the received response
type Backend struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`
	// the name of the group the response should be moved to. If empty, the response is
	// not changed
	Group string `gorm:"column:group" mapstructure:"group" json:"group" yaml:"group" toml:"group"`
	// HTTP method of the request to send to the backend
	Method string `gorm:"column:method" mapstructure:"method" json:"method" yaml:"method" toml:"method"`
	// Set of hosts of the API
	Host []string `gorm:"-" mapstructure:"host" json:"host" yaml:"host" toml:"host"`
	// HTTP headers of the endpoint to forward
	Header []string `gorm:"-" mapstructure:"header" json:"header" yaml:"header" toml:"header"` // many2many:backends_header;
	// False if the hostname should be sanitized
	HostSanitizationDisabled bool `gorm:"column:disable_host_sanitize" mapstructure:"disable_host_sanitize" json:"disable_host_sanitize" yaml:"disable_host_sanitize" toml:"disable_host_sanitize"`
	// URL pattern to use to locate the resource to be consumed
	URLPattern string `gorm:"column:url_pattern" mapstructure:"url_pattern" json:"url_pattern" yaml:"url_pattern" toml:"url_pattern"`
	// URL pattern to use to locate the resource to be consumed
	QueryStrings map[string]string `gorm:"column:query_strings" mapstructure:"query_strings" json:"query_strings" yaml:"query_strings" toml:"query_strings"`
	// URL pattern to use to locate the resource to be consumed
	Parameters map[string]string `gorm:"column:parameters" mapstructure:"parameters" json:"parameters" yaml:"parameters" toml:"parameters"`
	// set of response fields to remove. If empty, the filter id not used
	Blacklist []string `gorm:"-" mapstructure:"blacklist" json:"blacklist" yaml:"blacklist" toml:"blacklist"` // many2many:backends_blacklist;
	// set of response fields to allow. If empty, the filter id not used
	Whitelist []string `gorm:"-" mapstructure:"whitelist" json:"whitelist" yaml:"whitelist" toml:"whitelist"` // many2many:backends_whitelist;
	// map of response fields to be renamed and their new names
	Mapping map[string]string `gorm:"-" mapstructure:"mapping" json:"mapping" yaml:"mapping" toml:"mapping"` // many2many:backends_mapping;
	// the encoding format
	Encoding string `gorm:"column:encoding" mapstructure:"encoding" json:"encoding" yaml:"encoding" toml:"encoding"`
	// the response to process is a collection
	IsCollection bool `gorm:"column:is_collection" mapstructure:"is_collection" json:"is_collection" yaml:"is_collection" toml:"is_collection"`
	// name of the field to extract to the root. If empty, the formater will do nothing
	Target string `gorm:"column:target" mapstructure:"target" json:"target" yaml:"target" toml:"target"`
	// list of keys to be replaced in the URLPattern
	URLKeys []string `gorm:"-" mapstructure:"url_keys" json:"url_keys" yaml:"url_keys" toml:"url_keys"`
	// number of concurrent calls this endpoint must send to the API
	ConcurrentCalls int `gorm:"column:concurrent_calls" mapstructure:"concurrent_calls" json:"concurrent_calls" yaml:"concurrent_calls" toml:"concurrent_calls"`
	// timeout of this backend
	Timeout time.Duration `gorm:"column:timeout" mapstructure:"timeout" json:"timeout" yaml:"timeout" toml:"timeout"`
	// decoder to use in order to parse the received response from the API
	Decoder encoding.Decoder `gorm:"-" mapstructure:"decoder" json:"decoder" yaml:"decoder" toml:"decoder"`
	// Backend Extra configuration for customized behaviours
	ExtraConfig ExtraConfig `gorm:"-" mapstructure:"extra_config" json:"extra_config" yaml:"-" toml:"extra_config"` // many2many:backends_extra_config;
}
