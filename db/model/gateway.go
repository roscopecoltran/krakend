package model

import (
	"github.com/jinzhu/gorm"
)

type GatewayConfig struct {
	gorm.Model `json:"-" yaml:"-"`
	Disabled   bool   `default:"false" env:"GATEWAY_DISABLED" json:"disabled"  yaml:"disabled"`
	Debug      bool   `default:"false" env:"GATEWAY_DEBUG" json:"debug"  yaml:"debug"`
	Default    bool   `default:"false" env:"GATEWAY_DEFAULT" json:"default"  yaml:"default"`
	Port       uint   `default:"8080" env:"GATEWAY_PORT"  json:"port"  yaml:"port"`
	Host       string `default:"localhost" env:"GATEWAY_HOST"  json:"host"  yaml:"host"`
	ListenAddr string `default:":8080" env:"GATEWAY_LISTEN_ADDR"  json:"listen_addr" yaml:"listen_addr"`
	Mode       string `env:"GATEWAY_MODE" default:"dev" json:"mode"  yaml:"mode"`
}
