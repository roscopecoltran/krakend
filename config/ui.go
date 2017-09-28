package config

import (
	"github.com/jinzhu/gorm"
)

type WebInterfaceConfig struct {
	gorm.Model    `json:"-" yaml:"-" toml:"-"`
	Debug         bool   `gorm:"column:debug" default:"false" example:"false" json:"debug" yaml:"debug" toml:"debug"`
	DefaultLocale string `gorm:"column:default_locale" json:"default_locale" yaml:"default_locale" toml:"default_locale"`
	DefaultTheme  string `gorm:"column:default_theme" json:"default_theme" example:"oscar" yaml:"default_theme" example:"oscar" toml:"default_theme" example:"oscar"`
	StaticPath    string `gorm:"column:static_path" json:"static_path" yaml:"static_path" toml:"static_path"`
	TemplatesPath string `gorm:"column:templates_path" json:"templates_path" yaml:"templates_path" toml:"templates_path"`
}
