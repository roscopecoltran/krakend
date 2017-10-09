package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/qor/admin"
	"github.com/qor/qor"
)

func main() {

	DB, _ := gorm.Open("sqlite3", "./shared/data/krakend/gateway.db") // update with krakend-admin packages later
	DB.AutoMigrate(&User{}, &Product{})

	// Initalize
	Admin := admin.New(&qor.Config{DB: DB})

	// Create resources from GORM-backend model
	Admin.AddResource(&User{})
	Admin.AddResource(&Service{})

	// Register route
	mux := http.NewServeMux()
	// amount to /admin, so visit `/admin` to view the admin interface
	Admin.MountTo("/admin", mux)

	fmt.Println("Listening on: 9000")

	mux := http.NewServeMux()
	Admin.MountTo("/admin", mux)

	r := gin.Default()
	r.Any("/admin/*w", gin.WrapH(mux))
	r.Run()

}
