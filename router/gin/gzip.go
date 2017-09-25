package gin

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

/*
GZIP compress responses
*/
func GZIP() gin.HandlerFunc {
	return gzip.Gzip(gzip.DefaultCompression)
}
