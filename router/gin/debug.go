package gin

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"

	"github.com/roscopecoltran/krakend/logging"
)

var (
	green   = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	white   = string([]byte{27, 91, 57, 48, 59, 52, 55, 109})
	yellow  = string([]byte{27, 91, 57, 55, 59, 52, 51, 109})
	red     = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	blue    = string([]byte{27, 91, 57, 55, 59, 52, 52, 109})
	magenta = string([]byte{27, 91, 57, 55, 59, 52, 53, 109})
	cyan    = string([]byte{27, 91, 57, 55, 59, 52, 54, 109})
	reset   = string([]byte{27, 91, 48, 109})
)

// DebugHandler creates a dummy handler function, useful for quick integration tests
func DebugHandler(logger logging.Logger) gin.HandlerFunc {
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
			"message": "pong",
		})
	}
}
