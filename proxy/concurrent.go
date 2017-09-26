package proxy

import (
	"context"
	"errors"
	"time"

	"fmt"
	"github.com/k0kubun/pp"

	"github.com/roscopecoltran/krakend/config"
)

// NewConcurrentMiddleware creates a proxy middleware that enables sending several requests concurrently
func NewConcurrentMiddleware(remote *config.Backend) Middleware {
	if remote.ConcurrentCalls == 1 {
		panic(ErrTooManyProxies)
	}
	serviceTimeout := time.Duration(75*remote.Timeout.Nanoseconds()/100) * time.Nanosecond

	return func(next ...Proxy) Proxy {
		if len(next) > 1 {
			panic(ErrTooManyProxies)
		}

		return func(ctx context.Context, request *Request) (*Response, error) {
			localCtx, cancel := context.WithTimeout(ctx, serviceTimeout)

			results := make(chan *Response, remote.ConcurrentCalls)
			failed := make(chan error, remote.ConcurrentCalls)

			for i := 0; i < remote.ConcurrentCalls; i++ {
				go processConcurrentCall(localCtx, next[0], request, results, failed)
			}

			var response *Response
			var err error

			for i := 0; i < remote.ConcurrentCalls; i++ {
				select {
				case response = <-results:
					if response != nil && response.IsComplete {
						cancel()
						return response, nil
					}
				case err = <-failed:
				case <-ctx.Done():
				}
			}
			cancel()
			return response, err
		}
	}
}

var errNullResult = errors.New("invalid response")

func processConcurrentCall(ctx context.Context, next Proxy, request *Request, out chan<- *Response, failed chan<- error) {
	localCtx, cancel := context.WithCancel(ctx)
	result, err := next(localCtx, request)

	if config.Config.Debug.Components.Proxies {
		// fmt.Println("proxy/concurrent.go > var.request.Code")
		// pp.Print(request.Code)
		// fmt.Println("proxy/concurrent.go > var.request.Header")
		// pp.Print(request.Header)
		// fmt.Println("proxy/concurrent.go > var.result")
		// pp.Print(result)
		if err != nil {
			fmt.Println("proxy/concurrent.go > var.err")
			pp.Print(err)
		}
	}
	if err != nil {
		failed <- err
		cancel()
		return
	}
	if result == nil {
		failed <- errNullResult
		cancel()
		return
	}
	select {
	case out <- result:
	case <-ctx.Done():
		failed <- ctx.Err()
	}
	cancel()
}
