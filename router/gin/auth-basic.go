package gin

import (
	b64 "encoding/base64"
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/roscopecoltran/krakend/logging"
)

// AuthRequired require Authorization with token Basic64 encoded
func AuthRequiredHandler(logger logging.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")
		if header == "" || !strings.Contains(header, "Basic") {
			c.AbortWithStatus(403)
		}

		value := strings.Split(header, " ")

		if len(value) != 2 {
			c.AbortWithStatus(403)
		}

		token := value[1]
		sEnc := b64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", os.Getenv("ADMIN_USERNAME"), os.Getenv("ADMIN_PASSWORD"))))
		if token != sEnc {
			c.AbortWithStatus(403)
		}
	}
}
