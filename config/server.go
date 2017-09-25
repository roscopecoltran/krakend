package config

import (
	"github.com/jinzhu/gorm"
)

type ServerConfig struct {
	gorm.Model `json:"-" yaml:"-"`
	Port       int    `default:"8080" env:"SERVER_PORT"  json:"port"  yaml:"port"`
	Host       string `default:"localhost" env:"SERVER_HOST"  json:"host"  yaml:"host"`
	ListenAddr string `default:":8080" env:"SERVER_LISTEN_ADDR"  json:"listen_addr" yaml:"listen_addr"`
	Mode       string `env:"SERVER_MODE" default:"dev" json:"mode"  yaml:"mode"`
}
