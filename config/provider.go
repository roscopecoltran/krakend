package config

import (
	// "errors"
	"time"

	"github.com/jinzhu/gorm"
)

type ProviderConfig struct {
	gorm.Model      `json:"-" yaml:"-" toml:"-"`
	Disabled        bool          `gorm:"column:disabled" default:"false" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug           bool          `gorm:"column:debug" default:"false" example:"false" json:"debug" yaml:"debug" toml:"debug"`
	Name            string        `gorm:"column:name" json:"name,omitempty" example:"arch linux wiki" yaml:"name,omitempty" toml:"name,omitempty"`
	Slug            string        `gorm:"column:slug" json:"slug,omitempty" yaml:"slug,omitempty" toml:"slug,omitempty"`
	BaseURL         string        `gorm:"column:base_url" json:"base_url,omitempty" yaml:"base_url,omitempty" toml:"base_url,omitempty"`
	Topics          []TopicConfig `gorm:"many2many:provider_topics;" json:"topics,omitempty" yaml:"topics,omitempty" toml:"topics,omitempty"`
	Categories      string        `gorm:"column:categories" json:"categories,omitempty" yaml:"categories,omitempty" toml:"categories,omitempty"`
	ContentQuery    string        `gorm:"column:content_query" json:"content_query,omitempty" yaml:"content_query,omitempty" toml:"content_query,omitempty"`
	ContentXpath    string        `gorm:"column:content_xpath" json:"content_xpath,omitempty" yaml:"content_xpath,omitempty" toml:"content_xpath,omitempty"`
	Engine          string        `gorm:"column:engine" json:"engine,omitempty" example:"archlinux" yaml:"engine,omitempty" toml:"engine,omitempty"`
	FirstPageNum    int           `gorm:"column:first_page_num" default:"1" json:"first_page_num,omitempty" yaml:"first_page_num,omitempty" toml:"first_page_num,omitempty"`
	NumberOfResults int           `gorm:"column:number_of_results" default:"100" json:"number_of_results,omitempty" yaml:"number_of_results,omitempty" toml:"number_of_results,omitempty"`
	PageSize        int           `gorm:"column:page_size" default:"100" json:"page_size,omitempty" yaml:"page_size,omitempty" toml:"page_size,omitempty"`
	Paging          bool          `gorm:"column:paging" default:"true" json:"paging,omitempty" yaml:"paging,omitempty" toml:"paging,omitempty"`
	ResultsQuery    string        `gorm:"column:results_query" json:"results_query,omitempty" yaml:"results_query,omitempty" toml:"results_query,omitempty"`
	ResultsXpath    string        `gorm:"column:results_xpath" json:"results_xpath,omitempty" yaml:"results_xpath,omitempty" toml:"results_xpath,omitempty"`
	SearchType      string        `gorm:"column:search_type" json:"search_type,omitempty" yaml:"search_type,omitempty" toml:"search_type,omitempty"`
	SearchURL       string        `gorm:"column:search_url" json:"search_url,omitempty" yaml:"search_url,omitempty" toml:"search_url,omitempty"`
	Shortcut        string        `gorm:"column:shortcut" json:"shortcut,omitempty" yaml:"shortcut,omitempty" toml:"shortcut,omitempty" example:"al"`
	SuggestionXpath string        `gorm:"column:suggestion_xpath" json:"suggestion_xpath,omitempty" yaml:"suggestion_xpath,omitempty" toml:"suggestion_xpath,omitempty"`
	TitleQuery      string        `gorm:"column:title_query" json:"title_query,omitempty" yaml:"title_query,omitempty" toml:"title_query,omitempty"`
	TitleXpath      string        `gorm:"column:title_xpath" json:"title_xpath,omitempty" yaml:"title_xpath,omitempty" toml:"title_xpath,omitempty"`
	URL             string        `gorm:"column:url" json:"url,omitempty" yaml:"url,omitempty" toml:"url,omitempty"`
	URLQuery        string        `gorm:"column:url_query" json:"url_query,omitempty" yaml:"url_query,omitempty" toml:"url_query,omitempty"`
	URLXpath        string        `gorm:"column:url_xpath" json:"url_xpath,omitempty" yaml:"url_xpath,omitempty" toml:"url_xpath,omitempty"`
	Weight          int           `gorm:"column:weight" json:"weight,omitempty" yaml:"weight,omitempty" toml:"weight,omitempty"`
	Timeout         time.Duration `gorm:"column:timeout" json:"timeout,omitempty" yaml:"timeout,omitempty" default:"2" toml:"timeout,omitempty" default:"2"`
	ConcurrentCalls int           `gorm:"column:concurrent_calls" default:"10" json:"concurrent_calls,omitempty" yaml:"concurrent_calls,omitempty" toml:"concurrent_calls,omitempty"`
	Specs           []SpecsConfig `gorm:"many2many:provider_specs;" json:"specs,omitempty" yaml:"specs,omitempty" toml:"specs,omitempty"`
}

/*
// ProviderResult wraps a provider and an error
type ProviderResult struct {
	Provider *ProviderConfig
	Error    error
}
*/
