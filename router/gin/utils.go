package gin

import (
	// "errors"
	"io/ioutil"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	// "github.com/gin-contrib/size"
	// "github.com/roscopecoltran/krakend/config"
	"github.com/roscopecoltran/krakend/logging"
)

/*
GZIP compress responses
*/
func GZIP() gin.HandlerFunc {
	return gzip.Gzip(gzip.DefaultCompression)
}

/*
// SendStatus takes an integer and sets the response status to the integer given.
func (r ginRouter) SendStatus(statusCode int) {
	r.cfg.Engine.statusCode = statusCode
	r.ctx.Response.WriteHeader(statusCode)
}

// StatusCode returns the status code that golf has sent.
func (r ginRouter) StatusCode() int {
	return r.ctx.statusCode
}

// SetHeader sets the header entries associated with key to the single element value. It replaces any existing values associated with key.
func (r ginRouter) SetHeader(key, value string) {
	r.ctx.Response.Header().Set(key, value)
}

// AddHeader adds the key, value pair to the header. It appends to any existing values associated with key.
func (r ginRouter) AddHeader(key, value string) {
	r.ctx.Response.Header().Add(key, value)
}

// Header gets the first value associated with the given key. If there are no values associated with the key, Get returns "".
func (r ginRouter) Header(key string) string {
	return r.ctx.Request.Header.Get(key)
}

// Query method retrieves the form data, return empty string if not found.
func (r ginRouter) Query(key string, index ...int) (string, error) {
	if val, ok := r.ctx.Request.Form[key]; ok {
		if len(index) == 1 {
			return val[index[0]], nil
		}
		return val[0], nil
	}
	return "", errors.New("Query key not found")
}

// Param method retrieves the parameters from url
// If the url is /:id/, then id can be retrieved by calling `ctx.Param(id)`
func (r ginRouter) Param(key string) string {
	val, _ := r.ctx.Params.ByName(key)
	return val
}

// Redirect method sets the response as a 302 redirection.
func (r ginRouter) Redirect(url string) {
	r.ctx.SetHeader("Location", url)
	r.ctx.SendStatus(302)
}

// Redirect301 method sets the response as a 301 redirection.
func (r ginRouter) Redirect301(url string) {
	r.ctx.SetHeader("Location", url)
	r.ctx.SendStatus(301)
}

// Abort method returns an HTTP Error by indicating the status code, the corresponding
// handler inside `App.errorHandler` will be called, if user has not set
// the corresponding error handler, the defaultErrorHandler will be called.
func (r ginRouter) Abort(statusCode int, data ...map[string]interface{}) {
	r.ctx.App.handleError(r.ctx, statusCode, data...)
}

// Cookie returns the value of the cookie by indicating the key.
func (r ginRouter) Cookie(key string) (string, error) {
	c, err := r.ctx.Request.Cookie(key)
	if err != nil {
		return "", err
	}
	return c.Value, nil
}

// SetCookie set cookies for the request. If expire is 0, create a session cookie.
func (r ginRouter) SetCookie(key string, value string, expire int) {
	now := time.Now()
	cookie := &http.Cookie{
		Name:   key,
		Value:  value,
		Path:   "/",
		MaxAge: expire,
	}
	if expire != 0 {
		expireTime := now.Add(time.Duration(expire) * time.Second)
		cookie.Expires = expireTime
	}
	http.SetCookie(r.ctx.Response, cookie)
}

// ClientIP implements a best effort algorithm to return the real client IP, it parses
// X-Real-IP and X-Forwarded-For in order to work properly with reverse-proxies such us: nginx or haproxy.
// This method is taken from https://github.com/gin-gonic/gin
func (r ginRouter) ClientIP() string {
	clientIP := strings.TrimSpace(ctx.requestHeader("X-Real-Ip"))
	if len(clientIP) > 0 {
		return clientIP
	}
	clientIP = r.ctx.requestHeader("X-Forwarded-For")
	if index := strings.IndexByte(clientIP, ','); index >= 0 {
		clientIP = clientIP[0:index]
	}
	clientIP = strings.TrimSpace(clientIP)
	if len(clientIP) > 0 {
		return clientIP
	}
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.ctx.Request.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}
*/

func SizeHandler(logger logging.Logger) gin.HandlerFunc {
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
