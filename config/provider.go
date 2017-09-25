package config

import (
	"time"

	"github.com/jinzhu/gorm"
)

type ProviderConfig struct {
	gorm.Model      `json:"-" yaml:"-" toml:"-"`
	Disabled        bool          `gorm:"column:disabled" default:"false" json:"disabled" yaml:"disabled" toml:"disabled"`
	Name            string        `gorm:"column:name" json:"name,omitempty" example:"arch linux wiki" yaml:"name,omitempty" toml:"name,omitempty"`
	Slug            string        `gorm:"column:slug" json:"slug,omitempty" yaml:"slug,omitempty" toml:"slug,omitempty"`
	BaseURL         string        `gorm:"column:base_url" json:"base_url,omitempty" yaml:"base_url,omitempty" toml:"base_url,omitempty"`
	Categories      string        `gorm:"column:categories" json:"categories" yaml:"categories" toml:"categories"`
	ContentQuery    string        `gorm:"column:content_query" json:"content_query" yaml:"content_query" toml:"content_query"`
	ContentXpath    string        `gorm:"column:content_xpath" json:"content_xpath" yaml:"content_xpath" toml:"content_xpath"`
	Engine          string        `gorm:"column:engine" json:"engine,omitempty" example:"archlinux" yaml:"engine,omitempty" toml:"engine,omitempty"`
	FirstPageNum    int           `gorm:"column:first_page_num" json:"first_page_num" yaml:"first_page_num" toml:"first_page_num"`
	NumberOfResults int           `gorm:"column:number_of_results" json:"number_of_results" yaml:"number_of_results" toml:"number_of_results"`
	PageSize        int           `gorm:"column:page_size" json:"page_size" yaml:"page_size" toml:"page_size"`
	Paging          bool          `gorm:"column:paging" json:"paging" yaml:"paging" toml:"paging"`
	ResultsQuery    string        `gorm:"column:results_query" json:"results_query" yaml:"results_query" toml:"results_query"`
	ResultsXpath    string        `gorm:"column:results_xpath" json:"results_xpath" yaml:"results_xpath" toml:"results_xpath"`
	SearchType      string        `gorm:"column:search_type" json:"search_type,omitempty" yaml:"search_type,omitempty" toml:"search_type,omitempty"`
	SearchURL       string        `gorm:"column:search_url" json:"search_url,omitempty" yaml:"search_url,omitempty" toml:"search_url,omitempty"`
	Shortcut        string        `gorm:"column:shortcut" json:"shortcut" example:"al" yaml:"shortcut" example:"al" toml:"shortcut" example:"al"`
	SuggestionXpath string        `gorm:"column:suggestion_xpath" json:"suggestion_xpath" yaml:"suggestion_xpath" toml:"suggestion_xpath"`
	TitleQuery      string        `gorm:"column:title_query" json:"title_query" yaml:"title_query" toml:"title_query"`
	TitleXpath      string        `gorm:"column:title_xpath" json:"title_xpath" yaml:"title_xpath" toml:"title_xpath"`
	URL             string        `gorm:"column:url" json:"url,omitempty" yaml:"url,omitempty" toml:"url,omitempty"`
	URLQuery        string        `gorm:"column:url_query" json:"url_query" yaml:"url_query" toml:"url_query"`
	URLXpath        string        `gorm:"column:url_xpath" json:"url_xpath" yaml:"url_xpath" toml:"url_xpath"`
	Weight          int           `gorm:"column:weight" json:"weight" yaml:"weight" toml:"weight"`
	Timeout         time.Duration `gorm:"column:timeout" json:"timeout" yaml:"timeout" default:"2" toml:"timeout" default:"2"`
	ConcurrentCalls int           `gorm:"column:concurrent_calls" json:"concurrent_calls" yaml:"concurrent_calls" default:"10" toml:"concurrent_calls" default:"10" `
	Specs           SpecsConfig   `gorm:"column:specs" json:"engine,specs" yaml:"engine,specs" toml:"engine,specs"`
}
