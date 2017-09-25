package config

import (
	"github.com/jinzhu/gorm"
)

type SearchSettingsConfig struct {
	gorm.Model   `json:"-" yaml:"-" toml:"-"`
	Autocomplete AutocompleteConfig  `gorm:"column:autocomplete" json:"autocomplete" yaml:"autocomplete" toml:"autocomplete"`
	Language     string              `gorm:"column:language" default:"all" json:"language" yaml:"language" example:"all" toml:"language"`
	Queries      SearchQueriesConfig `gorm:"column:queries" json:"queries" yaml:"queries" toml:"queries"`
	Results      SearchResultsConfig `gorm:"column:results" json:"results" yaml:"results" toml:"results"`
}

type SearchQueriesConfig struct {
	gorm.Model      `json:"-" yaml:"-" toml:"-"`
	Cache           bool `gorm:"column:cache" default:"true" json:"cache" yaml:"cache" example:"true" toml:"cache"`
	Expansion       bool `gorm:"column:expansion" default:"true" json:"expansion" yaml:"expansion" example:"true" toml:"expansion"`
	DomainDiscovery bool `gorm:"column:domain_discovery" default:"true" json:"domain_discovery" yaml:"domain_discovery" toml:"domain_discovery" example:"true"`
}

type SearchResultsConfig struct {
	gorm.Model     `json:"-" yaml:"-" toml:"-"`
	Related        bool `gorm:"column:related" default:"true" json:"related" yaml:"related" example:"true" toml:"related"`
	KnowledgeGraph bool `gorm:"column:knowledge_graph" default:"true" json:"knowledge_graph" yaml:"knowledge_graph" example:"true" toml:"knowledge_graph"`
	SafeSearch     bool `gorm:"column:safe_search" default:"false" json:"safe_search" yaml:"safe_search" example:"false" toml:"safe_search"`
	Topics         bool `gorm:"column:topics" default:"true" json:"topics" yaml:"topics" example:"true" toml:"topics"`
	PerPage        int  `gorm:"column:per_page" default:"25" json:"per_page" yaml:"per_page" example:"25" toml:"per_page"`
	MaxPage        int  `gorm:"column:max_page" default:"100" json:"max_page" yaml:"max_page" example:"100" toml:"max_page"`
}
