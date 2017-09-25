package config

import (
	"github.com/jinzhu/gorm"
)

type SpecsConfig struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`
	Disabled   bool      `gorm:"column:disabled" default:"false" json:"disabled" yaml:"disabled" toml:"disabled"`
	API        APIConfig `gorm:"column:api" json:"api,omitempty" yaml:"api,omitempty" toml:"api,omitempty"`
}
