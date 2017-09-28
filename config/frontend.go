package config

import (
	"github.com/jinzhu/gorm"
)

type AutocompleteConfig struct {
	gorm.Model      `json:"-" yaml:"-" toml:"-"`
	Disabled        bool                               `gorm:"column:disabled" default:"false" json:"disabled,omitempty" yaml:"disabled,omitempty" example:"false" toml:"disabled,omitempty"`
	Debug           bool                               `gorm:"column:debug" default:"false" json:"debug" yaml:"debug" toml:"debug"`
	Cache           bool                               `gorm:"column:cache" default:"true" json:"cache,omitempty" yaml:"cache,omitempty" example:"true" toml:"cache,omitempty"`
	ConcurrentCalls int                                `gorm:"column:concurrent_calls" default:"2" json:"concurrent_calls,omitempty" yaml:"concurrent_calls,omitempty" toml:"concurrent_calls,omitempty" example:"2"`
	Default         string                             `gorm:"column:default" default:"google" json:"default,omitempty" yaml:"default,omitempty" toml:"default,omitempty" example:"google"`
	Providers       []AutocompleteEngineSettingsConfig `gorm:"column:providers" json:"providers,omitempty" yaml:"providers,omitempty" toml:"providers,omitempty"`
}

type AutocompleteEngineSettingsConfig struct {
	gorm.Model     `json:"-" yaml:"-" toml:"-"`
	Name           string `gorm:"column:name" json:"name,omitempty" yaml:"name,omitempty" toml:"name,omitempty" example:"bing"`
	Disabled       bool   `gorm:"column:disabled" default:"false" json:"disabled,omitempty" yaml:"disabled,omitempty" toml:"disabled,omitempty" example:"false"`
	Debug          bool   `gorm:"column:debug" default:"false" json:"debug" yaml:"debug" toml:"debug"`
	Slug           string `gorm:"column:slug" json:"slug,omitempty" yaml:"slug,omitempty" toml:"slug,omitempty" example:"bing-autocomplete" `
	Priority       int    `gorm:"column:priority" default:"1" json:"priority,omitempty" yaml:"priority,omitempty" toml:"priority,omitempty" example:"1"`
	InputMinLength int    `gorm:"column:input_min_length" default:"3" json:"input_min_length,omitempty" yaml:"input_min_length,omitempty" toml:"input_min_length,omitempty" example:"3"`
	InputMaxLength int    `gorm:"column:input_max_length" default:"32" json:"input_max_length,omitempty" yaml:"input_max_length,omitempty" toml:"input_max_length,omitempty" example:"32"`
	ResultsMin     int    `gorm:"column:results_min" default:"1" json:"results_min,omitempty" yaml:"results_min,omitempty" toml:"results_min,omitempty" example:"1"`
	ResultsMax     int    `gorm:"column:results_max" default:"10" json:"results_max,omitempty" yaml:"results_max,omitempty" toml:"results_max,omitempty" example:"10"`
}
