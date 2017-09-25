package gin

import (
	"time"

	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/roscopecoltran/krakend/logging"
)

// ref. https://github.com/stevokk/netflixplusplus/blob/master/api/main.go
func CORSMiddleware(logger logging.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Debug("Method:", c.Request.Method)
		logger.Debug("URL:", c.Request.RequestURI)
		logger.Debug("Query:", c.Request.URL.Query())
		logger.Debug("Params:", c.Params)
		logger.Debug("Current-Headers:", c.Request.Header)
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		logger.Debug("New-Headers:", c.Request.Header)
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

/*
CORS allow CORS.
See: https://en.wikipedia.org/wiki/Cross-origin_resource_sharing
*/
func CORS() gin.HandlerFunc {
	return cors.New(cors.Config{
		AbortOnError:    false,
		AllowAllOrigins: true,
		// AllowedOrigins:   []string{"*"}, // TODO: set GUI url
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "HEAD", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}
