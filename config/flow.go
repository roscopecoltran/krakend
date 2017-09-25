package config

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

type FlowConfig struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`
	UUID       uuid.UUID `gorm:"-" storm:"uuid" storm:"uuid" json:"uuid" yaml:"uuid" toml:"uuid"`
	Disabled   bool      `gorm:"column:disabled" default:"false" json:"disabled" yaml:"disabled" toml:"disabled"`
	Name       string    `gorm:"column:name" json:"name" yaml:"name" example:"Github Stars" toml:"name" example:"Github Stars"`
	Slug       string    `gorm:"column:slug" json:"slug" yaml:"slug" example:"github-stars" toml:"slug" example:"github-stars"`
	// Call
}
