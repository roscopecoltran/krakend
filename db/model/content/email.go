package cms

import (
	"github.com/jinzhu/gorm"
	// "github.com/sirupsen/logrus"
)

type Email struct {
	gorm.Model
	Email           string `gorm:"column:email" json:"email,omitempty" yaml:"email,omitempty"`
	EmailType       string `default:"pro" gorm:"column:email_type" json:"email_type,omitempty" yaml:"email_type,omitempty"`
	Disabled        bool   `default:"false" gorm:"column:disabled" json:"disabled,omitempty" yaml:"disabled,omitempty"`
	RefServiceID    uint   `gorm:"column:ref_service_id" json:"ref_service_id,omitempty" yaml:"ref_service_id,omitempty"`
	RefServiceURL   string `gorm:"column:ref_service_url" json:"ref_service_url,omitempty" yaml:"ref_service_url,omitempty"`
	RefExternalUrls []Link `gorm:"column:ref_external_urls" json:"ref_external_urls,omitempty" yaml:"ref_external_urls,omitempty"`
}
