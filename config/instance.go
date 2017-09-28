package config

import (
	"time"

	"github.com/jinzhu/gorm"
)

type InstanceConfig struct {
	gorm.Model          `json:"-" yaml:"-" toml:"-"`
	Debug               bool   `gorm:"column:debug" default:"false" json:"debug" yaml:"debug" toml:"debug"`
	BaseURL             bool   `gorm:"column:base_url" default:"false" example:"false" json:"base_url" yaml:"base_url" toml:"base_url"`
	BindAddress         string `gorm:"column:bind_address" default:"127.0.0.1" example:"127.0.0.1" json:"bind_address" yaml:"bind_address" toml:"bind_address"`
	HTTPProtocolVersion string `gorm:"column:http_protocol_version" json:"http_protocol_version" example:"1.0" yaml:"http_protocol_version" toml:"http_protocol_version"`
	ImageProxy          bool   `gorm:"column:image_proxy" json:"image_proxy" example:"false" yaml:"image_proxy" toml:"image_proxy"`
	Port                int    `gorm:"column:port" default:"8888" example:"8888" json:"port" yaml:"port" toml:"port"`
	SecretKey           string `gorm:"column:secret_key" default:"ultrasecretkey" example:"ultrasecretkey" json:"secret_key" yaml:"secret_key" toml:"secret_key"`
}

type InstanceHistoryConfig struct {
	gorm.Model   `json:"-" yaml:"-" toml:"-"`
	Debug        bool      `gorm:"column:debug" default:"false" example:"false" json:"debug" yaml:"debug" toml:"debug"`
	InstanceName string    `gorm:"column:instance_name" default:"sniperx" example:"sniperx" json:"instance_name" yaml:"instance_name" toml:"instance_name"`
	StartedAt    time.Time `gorm:"column:started_at" json:"-" yaml:"-" toml:"-"`
}
