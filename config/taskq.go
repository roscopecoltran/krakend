package config

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

/*
	Refs:
	- https://github.com/xh3b4sd/taskq
	- https://github.com/richardwilkes/taskqueue
	- https://github.com/dogtools/dog
*/

type TaskQConfig struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`
	UUID       uuid.UUID `gorm:"-" storm:"uuid" storm:"uuid" json:"uuid" yaml:"uuid" toml:"uuid"`
	Disabled   bool      `default:"false" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug      bool      `default:"false" json:"debug" yaml:"debug" toml:"debug"`
	Type       string    `json:"type" yaml:"type" toml:"type"` // series, parallel
	Action     string    `json:"action" yaml:"action" toml:"action"`
	Order      int       `json:"order" yaml:"order" toml:"order"`
	Priority   int       `json:"priority" yaml:"priority" toml:"priority"`
}

/*
import (
	"github.com/dogtools/dog"
)
var Tasks     []dog.Dogfile   `json:"tasks" yaml:"tasks" toml:"tasks"`
*/
