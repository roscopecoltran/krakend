package config

import (
	"github.com/jinzhu/gorm"
)

type OutgoingConfig struct {
	gorm.Model      `json:"-" yaml:"-" toml:"-"`
	PoolConnections int    `gorm:"column:pool_connections" default:"100" json:"pool_connections" yaml:"pool_connections" example:"100" toml:"pool_connections" example:"100"`
	PoolMaxsize     int    `gorm:"column:pool_maxsize" default:"10" json:"pool_maxsize" yaml:"pool_maxsize" example:"10" toml:"pool_maxsize" example:"10"`
	RequestTimeout  int    `gorm:"column:request_timeout" default:"2" json:"request_timeout" yaml:"request_timeout" example:"2" toml:"request_timeout" example:"2"`
	UseragentSuffix string `gorm:"column:useragent_suffix" default:"SNIPER-X" json:"useragent_suffix" yaml:"useragent_suffix" toml:"useragent_suffix"`
}
