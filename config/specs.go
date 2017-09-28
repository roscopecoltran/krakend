package config

import (
	"github.com/jinzhu/gorm"
)

var DefaultSpecsType []string = []string{"raml", "wsdl", "wadl", "oai"} // (swagger/oai, wadl, wsdl, raml) "swagger"

/*
	Refs:
	- https://github.com/ghchinoy/atmotool/blob/master/apis/create.go#L80
*/

type SpecsConfig struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`
	Disabled   bool      `gorm:"column:disabled" default:"false" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug      bool      `gorm:"column:debug" default:"false" example:"false" json:"debug" yaml:"debug" toml:"debug"`
	API        APIConfig `gorm:"many2many:specs_configs;" json:"api,omitempty" yaml:"api,omitempty" toml:"api,omitempty"`
}
