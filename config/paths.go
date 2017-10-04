package config

//import (
// "errors"
//"time"
// "github.com/jinzhu/gorm"
//)

type PathConfig struct {
	JSON  PathBlocksConfig `gorm:"-" json:"json,omitempty" yaml:"json,omitempty" toml:"json,omitempty"`
	XPATH PathBlocksConfig `gorm:"-" json:"xpath,omitempty" yaml:"xpath,omitempty" toml:"xpath,omitempty"`
	XML   PathBlocksConfig `gorm:"-" json:"xml,omitempty" yaml:"xml,omitempty" toml:"xml,omitempty"`
}

type PathBlocksConfig struct {
	Article             string `gorm:"column:article" json:"article,omitempty" yaml:"article,omitempty" toml:"article,omitempty"`
	Author              string `gorm:"column:author" json:"author,omitempty" yaml:"author,omitempty" toml:"author,omitempty"`
	Content             string `gorm:"column:content" json:"content,omitempty" yaml:"content,omitempty" toml:"content,omitempty"`
	ContentMisc         string `gorm:"column:content_misc" json:"content_misc,omitempty" yaml:"content_misc,omitempty" toml:"content_misc,omitempty"`
	RelatedLinks        string `gorm:"column:related_links" json:"related_links,omitempty" yaml:"related_links,omitempty" toml:"related_links,omitempty"`
	Results             string `gorm:"column:results" json:"results,omitempty" yaml:"results,omitempty" toml:"results,omitempty"`
	SpellingSuggestions string `gorm:"column:spelling_suggestions" json:"spelling_suggestions,omitempty" yaml:"spelling_suggestions,omitempty" toml:"spelling_suggestions,omitempty"`
	Suggestions         string `gorm:"column:suggestions" json:"suggestions,omitempty" yaml:"suggestions,omitempty" toml:"suggestions,omitempty"`
	Tags                string `gorm:"column:tags" json:"tags,omitempty" yaml:"tags,omitempty" toml:"tags,omitempty"`
	Title               string `gorm:"column:title" json:"title,omitempty" yaml:"title,omitempty" toml:"title,omitempty"`
	Thumbnail           string `gorm:"column:thumbnail" json:"thumbnail,omitempty" yaml:"thumbnail,omitempty" toml:"thumbnail,omitempty"`
	URL                 string `gorm:"column:url" json:"url,omitempty" yaml:"url,omitempty" toml:"url,omitempty"`
}
