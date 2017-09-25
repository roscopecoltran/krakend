package config

import (
	"github.com/jinzhu/gorm"
)

type Provider struct {
	gorm.Model `json:"-" yaml:"-"`
	Name       string `json:"name"  yaml:"name"`
	// Swagger Definitions

}
