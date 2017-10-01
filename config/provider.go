package config

import (
	// "errors"
	"time"

	"github.com/jinzhu/gorm"
)

/*
	Refs:
	- https://github.com/mattruggio/search_engine_totals/blob/master/config/engines.yml
	- https://github.com/htano/summary_dev/blob/master/lib/contents-extractor/config/domain-extractor.yml
	- https://github.com/rgilbert1/article-subscriber/blob/master/sources.yml
	- https://github.com/asciimoo/searx/blob/master/searx/engines/google.py
	- https://github.com/asciimoo/searx/blob/master/searx/settings.yml

*/

type ProviderConfig struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`
	Disabled   bool `gorm:"column:disabled" default:"false" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug      bool `gorm:"column:debug" default:"false" example:"false" json:"debug" yaml:"debug" toml:"debug"`

	// Profile
	Name                  string                       `gorm:"column:name" json:"name,omitempty" example:"arch linux wiki" yaml:"name,omitempty" toml:"name,omitempty"`
	Slug                  string                       `gorm:"column:slug" json:"slug,omitempty" yaml:"slug,omitempty" toml:"slug,omitempty"`
	Shortcut              string                       `gorm:"column:shortcut" json:"shortcut,omitempty" yaml:"shortcut,omitempty" toml:"shortcut,omitempty" example:"al"`
	URL                   string                       `gorm:"column:url" json:"url,omitempty" yaml:"url,omitempty" toml:"url,omitempty"`
	BaseURL               string                       `gorm:"column:base_url" json:"base_url,omitempty" yaml:"base_url,omitempty" toml:"base_url,omitempty"`
	Engine                string                       `gorm:"column:engine" json:"engine,omitempty" example:"archlinux" yaml:"engine,omitempty" toml:"engine,omitempty"`
	SearchType            string                       `gorm:"column:search_type" json:"search_type,omitempty" yaml:"search_type,omitempty" toml:"search_type,omitempty"`
	SearchURL             string                       `gorm:"column:search_url" json:"search_url,omitempty" yaml:"search_url,omitempty" toml:"search_url,omitempty"`
	SearchString          string                       `gorm:"column:search_string" json:"search_string,omitempty" yaml:"search_string,omitempty" toml:"search_string,omitempty"`
	AcceptHeader          []string                     `gorm:"column:accept_header" json:"accept_header,omitempty" yaml:"accept_header,omitempty" toml:"accept_header,omitempty"`
	SafeSearch            bool                         `gorm:"column:safesearch" default:"false" json:"safesearch,omitempty" yaml:"safesearch,omitempty" toml:"safesearch,omitempty" example:"true"`
	SafeSearchTypes       map[string]string            `gorm:"column:safesearch_types" json:"safesearch_types,omitempty" yaml:"safesearch_types,omitempty" toml:"safesearch_types,omitempty"`
	CountryToHostname     map[string]string            `gorm:"column:country_to_hostname" json:"country_to_hostname,omitempty" yaml:"country_to_hostname,omitempty" toml:"country_to_hostname,omitempty"`
	Weight                int                          `gorm:"column:weight" json:"weight,omitempty" yaml:"weight,omitempty" toml:"weight,omitempty"`
	ConcurrentCalls       int                          `gorm:"column:concurrent_calls" default:"10" json:"concurrent_calls,omitempty" yaml:"concurrent_calls,omitempty" toml:"concurrent_calls,omitempty"`
	Timeout               time.Duration                `gorm:"column:timeout" json:"timeout,omitempty" yaml:"timeout,omitempty" default:"2" toml:"timeout,omitempty" default:"2"`
	TimeRangeSupport      bool                         `gorm:"column:time_range_support" default:"false" json:"time_range_support,omitempty" yaml:"time_range_support,omitempty" toml:"time_range_support,omitempty" example:"true"`
	TimeRangeString       string                       `gorm:"column:time_range_string" json:"time_range_string,omitempty" yaml:"time_range_string,omitempty" toml:"time_range_string,omitempty" example:"true"`
	TimeRangeDict         map[string]string            `gorm:"column:time_range_dict" json:"time_range_dict,omitempty" yaml:"time_range_dict,omitempty" toml:"time_range_dict,omitempty" example:"true"`
	FirstPageNum          int                          `gorm:"column:first_page_num" default:"1" json:"first_page_num,omitempty" yaml:"first_page_num,omitempty" toml:"first_page_num,omitempty"`
	NumberOfResults       int                          `gorm:"column:number_of_results" default:"100" json:"number_of_results,omitempty" yaml:"number_of_results,omitempty" toml:"number_of_results,omitempty"`
	PageSize              int                          `gorm:"column:page_size" default:"100" json:"page_size,omitempty" yaml:"page_size,omitempty" toml:"page_size,omitempty"`
	Paging                bool                         `gorm:"column:paging" default:"true" json:"paging,omitempty" yaml:"paging,omitempty" toml:"paging,omitempty"`
	LanguageSupport       bool                         `gorm:"column:language_support" default:"true" json:"language_support,omitempty" yaml:"language_support,omitempty" toml:"language_support,omitempty"`
	SupportedLanguagesURL string                       `gorm:"column:supported_languages_url"  json:"supported_languages_url,omitempty" yaml:"supported_languages_url,omitempty" toml:"supported_languages_url,omitempty"`
	SupportedLanguages    map[string]map[string]string `gorm:"column:supported_languages"  json:"supported_languages,omitempty" yaml:"supported_languages,omitempty" toml:"supported_languages,omitempty"`
	LangURLS              map[string]map[string]string `gorm:"column:lang_urls"  json:"lang_urls,omitempty" yaml:"lang_urls,omitempty" toml:"lang_urls,omitempty"`
	MainLangs             map[string]string            `gorm:"column:main_langs" json:"main_langs,omitempty" yaml:"main_langs,omitempty" toml:"main_langs,omitempty"`
	Specs                 []SpecsConfig                `gorm:"many2many:provider_specs;" json:"specs,omitempty" yaml:"specs,omitempty" toml:"specs,omitempty"`
	Topics                []TopicConfig                `gorm:"many2many:provider_topics;" json:"topics,omitempty" yaml:"topics,omitempty" toml:"topics,omitempty"`
	Categories            string                       `gorm:"column:categories" json:"categories,omitempty" yaml:"categories,omitempty" toml:"categories,omitempty"`
	ZeroIndicationStrings string                       `gorm:"column:zero_indication_strings" json:"zero_indication_strings,omitempty" yaml:"zero_indication_strings,omitempty" toml:"zero_indication_strings,omitempty"`

	// Paths (to refactor and optimize)
	ArticleQuery            string `gorm:"column:article_query" json:"article_query,omitempty" yaml:"article_query,omitempty" toml:"article_query,omitempty"`
	ArticleXpath            string `gorm:"column:article_xpath" json:"article_xpath,omitempty" yaml:"article_xpath,omitempty" toml:"article_xpath,omitempty"`
	AuthorQuery             string `gorm:"column:author_query" json:"author_query,omitempty" yaml:"author_query,omitempty" toml:"author_query,omitempty"`
	AuthorXpath             string `gorm:"column:author_xpath" json:"author_xpath,omitempty" yaml:"author_xpath,omitempty" toml:"author_xpath,omitempty"`
	ContentQuery            string `gorm:"column:content_query" json:"content_query,omitempty" yaml:"content_query,omitempty" toml:"content_query,omitempty"`
	ContentXpath            string `gorm:"column:content_xpath" json:"content_xpath,omitempty" yaml:"content_xpath,omitempty" toml:"content_xpath,omitempty"`
	ContentMiscQuery        string `gorm:"column:content_misc_query" json:"content_misc_query,omitempty" yaml:"content_misc_query,omitempty" toml:"content_misc_query,omitempty"`
	ContentMiscXpath        string `gorm:"column:content_misc_xpath" json:"content_misc_xpath,omitempty" yaml:"content_misc_xpath,omitempty" toml:"content_misc_xpath,omitempty"`
	ResultsQuery            string `gorm:"column:results_query" json:"results_query,omitempty" yaml:"results_query,omitempty" toml:"results_query,omitempty"`
	ResultsXpath            string `gorm:"column:results_xpath" json:"results_xpath,omitempty" yaml:"results_xpath,omitempty" toml:"results_xpath,omitempty"`
	SpellingSuggestionQuery string `gorm:"column:spelling_suggestion_query" json:"spelling_suggestion_query,omitempty" yaml:"spelling_suggestion_query,omitempty" toml:"spelling_suggestion_query,omitempty"`
	SpellingSuggestionXpath string `gorm:"column:spelling_suggestion_xpath" json:"spelling_suggestion_xpath,omitempty" yaml:"spelling_suggestion_xpath,omitempty" toml:"spelling_suggestion_xpath,omitempty"`
	SuggestionQuery         string `gorm:"column:suggestion_query" json:"suggestion_query,omitempty" yaml:"suggestion_query,omitempty" toml:"suggestion_query,omitempty"`
	SuggestionXpath         string `gorm:"column:suggestion_xpath" json:"suggestion_xpath,omitempty" yaml:"suggestion_xpath,omitempty" toml:"suggestion_xpath,omitempty"`
	TagsQuery               string `gorm:"column:tags_query" json:"tags_query,omitempty" yaml:"tags_query,omitempty" toml:"tags_query,omitempty"`
	TagsXpath               string `gorm:"column:tags_xpath" json:"tags_xpath,omitempty" yaml:"tags_xpath,omitempty" toml:"tags_xpath,omitempty"`
	TitleQuery              string `gorm:"column:title_query" json:"title_query,omitempty" yaml:"title_query,omitempty" toml:"title_query,omitempty"`
	TitleXpath              string `gorm:"column:title_xpath" json:"title_xpath,omitempty" yaml:"title_xpath,omitempty" toml:"title_xpath,omitempty"`
	ThumbnailQuery          string `gorm:"column:thumbnail_query" json:"thumbnail_query,omitempty" yaml:"thumbnail_query,omitempty" toml:"thumbnail_query,omitempty"`
	ThumbnailXpath          string `gorm:"column:thumbnail_xpath" json:"thumbnail_xpath,omitempty" yaml:"thumbnail_xpath,omitempty" toml:"thumbnail_xpath,omitempty"`
	URLQuery                string `gorm:"column:url_query" json:"url_query,omitempty" yaml:"url_query,omitempty" toml:"url_query,omitempty"`
	URLXpath                string `gorm:"column:url_xpath" json:"url_xpath,omitempty" yaml:"url_xpath,omitempty" toml:"url_xpath,omitempty"`

	Paths PathConfig `gorm:"-" json:"paths,omitempty" yaml:"paths,omitempty" toml:"paths,omitempty"`
}

/*
// ProviderResult wraps a provider and an error
type ProviderResult struct {
	Provider *ProviderConfig
	Error    error
}
*/
