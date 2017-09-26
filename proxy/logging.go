package proxy

import (
	"context"
	"time"

	"fmt"
	"github.com/k0kubun/pp"

	"github.com/roscopecoltran/krakend/config"
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

// NewLoggingMiddleware creates proxy middleware for logging requests and responses
func NewLoggingMiddleware(logger logging.Logger, name string) Middleware {
	return func(next ...Proxy) Proxy {
		if len(next) > 1 {
			panic(ErrTooManyProxies)
		}
		return func(ctx context.Context, request *Request) (*Response, error) {
			begin := time.Now()
			logger.Info(name, "proxy/logging.go > Calling backend")
			logger.Debug("proxy/logging.go > Request", request)
			result, err := next[0](ctx, request)
			// path := ctx.Request.URL.Path

			if config.Config.Debug.Components.Middlewares {
				fmt.Println("proxy/logging.go > var.request")
				pp.Print(request)
				fmt.Println("proxy/logging.go > var.result:")
				pp.Print(result)
				fmt.Println("proxy/logging.go > var.err:")
				pp.Print(err)

				/*
					end := time.Now()
					latency := end.Sub(begin)

					clientIP := ctx.Request.ClientIP()
					method := ctx.Request.Method
					statusCode := ctx.statusCode
					statusColor := colorForStatus(statusCode)
					methodColor := colorForMethod(method)

					fmt.Fprintf(out, "%v |%s %3d %s| %13v | %s |%s  %s %-7s %s\n",
						end.Format("2006/01/02 - 15:04:05"),
						statusColor, statusCode, reset,
						latency,
						clientIP,
						methodColor, reset, method,
						path,
					)
				*/

			}

			// logger.Debug("proxy/logging.go > Result", result)
			logger.Info(name, "proxy/logging.go > Call to backend took", time.Since(begin).String())
			if err != nil {
				logger.Warning(name, "proxy/logging.go > Call to backend failed:", err.Error())
				return result, err
			}
			if result == nil {
				logger.Warning(name, "proxy/logging.go > Call to backend returned a null response")
			}
			return result, err
		}
	}
}

func colorForStatus(code int) string {
	switch {
	case code >= 200 && code < 300:
		return green
	case code >= 300 && code < 400:
		return white
	case code >= 400 && code < 500:
		return yellow
	default:
		return red
	}
}

func colorForMethod(method string) string {
	switch method {
	case "GET":
		return blue
	case "POST":
		return cyan
	case "PUT":
		return yellow
	case "DELETE":
		return red
	case "PATCH":
		return green
	case "HEAD":
		return magenta
	case "OPTIONS":
		return white
	default:
		return reset
	}
}
