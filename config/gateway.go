package config

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

type GatewayConfig struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`
	UUID       uuid.UUID `gorm:"-" storm:"uuid" storm:"uuid" json:"uuid" yaml:"uuid" toml:"uuid"`
	Disabled   bool      `default:"false" env:"GATEWAY_DISABLED" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug      bool      `default:"false" env:"GATEWAY_DEBUG" json:"debug" yaml:"debug" toml:"debug"`
	Default    bool      `default:"false" env:"GATEWAY_DEFAULT" json:"default" yaml:"default" toml:"default"`
	Port       uint      `default:"8080" env:"GATEWAY_PORT" json:"port" yaml:"port" toml:"port"`
	Host       string    `default:"localhost" env:"GATEWAY_HOST" json:"host" yaml:"host" toml:"host"`
	ListenAddr string    `default:":8080" env:"GATEWAY_LISTEN_ADDR" json:"listen_addr" yaml:"listen_addr" toml:"listen_addr"`
	Mode       string    `env:"GATEWAY_MODE" default:"dev" json:"mode" yaml:"mode" toml:"mode"`
}

/*
type Customer struct {
	Id        uuid.UUID `storm:"id" storm:"index"`
	CreatedAt int64     `storm:"index"`
	Name      string    `storm:"index" form:"name" binding:"required,min=3,max=40"`
	Email     string    `storm:"index" form:"email" binding:"required,min=3,max=40"`
}
*/
