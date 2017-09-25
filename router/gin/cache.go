package gin

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/muesli/cache2go"
	// "github.com/gin-gonic/contrib/cache"
)

func NoCache() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Cache-Control", "no-store")
		c.Writer.Header().Set("Pragma", "no-cache")
		c.Next()
	}
}

func UseHTTPCache(cacheKey string, store *cache2go.CacheTable, c *gin.Context) bool {
	t := time.Now()

	cacheRes, err := store.Value(cacheKey)
	if err == nil {
		t = *cacheRes.Data().(*time.Time)
	}

	if CacheHeader(c, t) {
		c.AbortWithStatus(http.StatusNotModified)
		return true
	}

	store.Add(cacheKey, ((30 * 24) * time.Hour), &t)
	return false
}

func CacheHeader(c *gin.Context, lastModified time.Time) bool {
	lastModifiedAt := lastModified.Format(time.RFC1123)

	c.Writer.Header().Add("Cache-Control", "public, max-age=31536000")
	c.Writer.Header().Add("Last-Modified", lastModifiedAt)

	if ifModifiedSince := c.Request.Header.Get("If-Modified-Since"); ifModifiedSince != "" {
		ifModifiedSinceTime, err := time.Parse(time.RFC1123, ifModifiedSince)
		if err != nil {
			return false
		}

		updatedAt, err := time.Parse(time.RFC1123, lastModifiedAt)
		if err != nil {
			return false
		}

		if ifModifiedSinceTime.Sub(updatedAt) < 1 {
			return true
		}
	}

	return false
}
