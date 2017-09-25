package config

import (
	"github.com/jinzhu/gorm"
)

type LinkConfig struct {
	gorm.Model  `json:"-" yaml:"-" toml:"-"`
	URL         string               `gorm:"column:url" json:"url" yaml:"url" toml:"url"`
	ContentType string               `gorm:"column:content_type" json:"content_type" yaml:"content_type" toml:"content_type"`
	Encoding    string               `gorm:"column:encoding" json:"encoding" yaml:"encoding" toml:"encoding"`
	Attr        LinkAttributesConfig `gorm:"column:attributes" json:"attributes" yaml:"attributes" toml:"attributes"`
	Comments    string               `gorm:"column:comments" json:"comments" yaml:"comments" toml:"comments"`
	Topics      []TopicConfig        `gorm:"column:link_topics" json:"topics" yaml:"topics" toml:"topics"`
}

type LinkAttributesConfig struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`
	Href       string            `gorm:"column:href" json:"href" yaml:"href" toml:"href"`
	Alt        string            `gorm:"column:alt" json:"language" yaml:"alt" toml:"alt"`
	Target     string            `gorm:"column:target" json:"target" yaml:"target" toml:"target"`
	Data       map[string]string `gorm:"column:data" json:"data" yaml:"data" toml:"data"`
}
