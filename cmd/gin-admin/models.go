package main

/*
import (
	"fmt"
	"github.com/roscopecoltran/krakend/config"
)
*/

// Create a GORM-backend model
type User struct {
	gorm.Model
	Name string
}

// Create another GORM-backend model
type Service struct {
	gorm.Model
	Name        string
	Description string
}
