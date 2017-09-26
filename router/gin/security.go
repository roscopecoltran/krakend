package gin

import (
	b64 "encoding/base64"
	"fmt"
	"os"
	"strings"

	// "github.com/gin-contrib/jwt"
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

/*
//JWTAuth JSON WEB TOKEN Auth middleware
//See: https://en.wikipedia.org/wiki/Cross-origin_resource_sharing
func JWTAuth(secretPassword string) gin.HandlerFunc {
	return jwt.Auth(secretPassword)
}
*/

func ApiToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: validate tokens
		c.Next()
	}
}

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
Depecrated: AllowedMethods, AbortOnError,AllowedHeaders
*/
/*
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
*/

/*
func (r ginRouter) getRawXSRFToken() string {
	token, err := r.cfg.Engine.Cookie("_xsrf")
	if err != nil {
		return ""
	}
	return token
}

func (r ginRouter) checkXSRFToken() bool {
	token := r.cfg.Engine.FormValue("xsrf_token")
	if token == "" {
		return false
	}
	_, tokenA, _ := decodeXSRFToken(token)
	_, tokenB, _ := decodeXSRFToken(r.getRawXSRFToken())
	return compareToken(tokenA, tokenB)
}

func (r ginRouter) xsrfToken() string {
	maskedToken := r.cfg.Engine.getRawXSRFToken()
	if maskedToken == "" {
		maskedToken = newXSRFToken()
		r.ctx.SetCookie("_xsrf", maskedToken, 3600)
	}
	_, tokenBytes, err := decodeXSRFToken(maskedToken)
	if err != nil {
		maskedToken = newXSRFToken()
		r.ctx.SetCookie("_xsrf", maskedToken, 3600)
		_, tokenBytes, _ = decodeXSRFToken(maskedToken)
	}
	maskBytes := randomBytes(4)
	maskedTokenBytes := append(maskBytes, websocketMask(maskBytes, tokenBytes)...)
	return hex.EncodeToString(maskedTokenBytes)
}
*/

// Tor
