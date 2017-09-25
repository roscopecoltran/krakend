package config

import (
	"github.com/jinzhu/gorm"
)

type TopicConfig struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`
	Label      string `gorm:"column:label" json:"label" yaml:"label" toml:"label"`
	Class      string `gorm:"column:class" json:"class" yaml:"class" toml:"class"`
	Group      string `gorm:"column:group" json:"group" yaml:"group" toml:"group"`
}
