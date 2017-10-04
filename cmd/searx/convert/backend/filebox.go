package backend

import (
	"github.com/qor/filebox"
	"github.com/qor/roles"
	// "github.com/roscopecoltran/krakend/cmd/searx/convert/config"
	// "github.com/roscopecoltran/krakend/cmd/searx/convert/config/auth"
)

var Filebox *filebox.Filebox

func init() {
	Filebox = filebox.New("./shared/public/downloads")
	// Filebox.SetAuth(auth.AdminAuth{})
	dir := Filebox.AccessDir("/")
	dir.SetPermission(roles.Allow(roles.Read, "admin"))
}
