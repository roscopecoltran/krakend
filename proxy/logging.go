package proxy

import (
	"context"
	"time"

	"github.com/roscopecoltran/krakend/logging"
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
			logger.Debug("proxy/logging.go > Result", result)
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
