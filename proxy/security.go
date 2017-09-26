package proxy

/*
import (
	"log"
	"net/http/httputil"
)

// XSRFProtectionMiddleware is the built-in middleware for XSRF protection.
func XSRFProtectionMiddleware(next HandlerFunc) HandlerFunc {
	fn := func(ctx *Context) {
		if ctx.Request.Method == "POST" || ctx.Request.Method == "PUT" || ctx.Request.Method == "DELETE" {
			if !ctx.checkXSRFToken() {
				ctx.Abort(403)
				return
			}
		}
		next(ctx)
	}
	return fn
}

*/
