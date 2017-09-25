package gin

import (
	"io/ioutil"

	// "github.com/qor/admin"
	// "github.com/qor/qor"
	"github.com/gin-gonic/gin"

	"github.com/roscopecoltran/krakend/db"
	"github.com/roscopecoltran/krakend/logging"
)

func AdminHandler(logger logging.Logger) gin.HandlerFunc {

	// db.Admin.MountTo("/admin", mux)
	// r.Any("/admin/*w", gin.WrapH(mux))

	logger.Debug("Admin:", db.ADMIN)

	return func(c *gin.Context) {
		logger.Debug("Method:", c.Request.Method)
		logger.Debug("URL:", c.Request.RequestURI)
		logger.Debug("Query:", c.Request.URL.Query())
		logger.Debug("Params:", c.Params)
		logger.Debug("Headers:", c.Request.Header)
		body, _ := ioutil.ReadAll(c.Request.Body)
		c.Request.Body.Close()
		logger.Debug("Body:", string(body))

		c.JSON(200, gin.H{
			"message": "admin",
		})
	}
}
