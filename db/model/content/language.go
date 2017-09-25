package cms

import (
	"errors"

	"github.com/jinzhu/gorm"
	// "github.com/sirupsen/logrus"
)

// Service represents a hosting service like Github
type Language struct {
	gorm.Model `json:"-" yaml:"-"`
	//sorting.SortingDESC
	Name string `gorm:"type:varchar(128);not null" yaml:"name,omitempty" json:"name,omitempty"`
}
