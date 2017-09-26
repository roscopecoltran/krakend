package proxy

import (
	"context"
	"errors"
	// "fmt"
	"io"
	"net/http"

	"fmt"
	"github.com/k0kubun/pp"

	"github.com/roscopecoltran/krakend/config"
	"github.com/roscopecoltran/krakend/encoding"
	// "github.com/roscopecoltran/krakend/logging"
	// "github.com/Jeffail/gabs"
)

// ErrInvalidStatusCode is the error returned by the http proxy when the received status code
// is not a 200 nor a 201
var ErrInvalidStatusCode = errors.New("Invalid status code")

// HTTPClientFactory creates http clients based with the received context
type HTTPClientFactory func(ctx context.Context) *http.Client

// NewHTTPClient just creates a http default client
func NewHTTPClient(_ context.Context) *http.Client { return http.DefaultClient }

var httpProxy = CustomHTTPProxyFactory(NewHTTPClient)

// HTTPProxyFactory returns a BackendFactory. The Proxies it creates will use the received net/http.Client
func HTTPProxyFactory(client *http.Client) BackendFactory {
	return CustomHTTPProxyFactory(func(_ context.Context) *http.Client { return client })
}

// CustomHTTPProxyFactory returns a BackendFactory. The Proxies it creates will use the received HTTPClientFactory
func CustomHTTPProxyFactory(cf HTTPClientFactory) BackendFactory {
	return func(backend *config.Backend) Proxy {
		return NewHTTPProxy(backend, cf, backend.Decoder)
	}
}

// HTTPRequestExecutor defines the interface of the request executor for the HTTP transport protocol
type HTTPRequestExecutor func(ctx context.Context, req *http.Request) (*http.Response, error)

// DefaultHTTPRequestExecutor creates a HTTPRequestExecutor with the received HTTPClientFactory
func DefaultHTTPRequestExecutor(clientFactory HTTPClientFactory) HTTPRequestExecutor {
	return func(ctx context.Context, req *http.Request) (*http.Response, error) {
		return clientFactory(ctx).Do(req.WithContext(ctx))
	}
}

// NewHTTPProxy creates a http proxy with the injected configuration, HTTPClientFactory and Decoder
func NewHTTPProxy(remote *config.Backend, clientFactory HTTPClientFactory, decode encoding.Decoder) Proxy {
	return NewHTTPProxyWithHTTPExecutor(remote, DefaultHTTPRequestExecutor(clientFactory), decode)
}

// NewHTTPProxyWithHTTPExecutor creates a http proxy with the injected configuration, HTTPRequestExecutor and Decoder
func NewHTTPProxyWithHTTPExecutor(remote *config.Backend, requestExecutor HTTPRequestExecutor, dec encoding.Decoder) Proxy {
	ef := NewEntityFormatter(remote.Target, remote.Whitelist, remote.Blacklist, remote.Group, remote.Mapping)
	return NewHTTPProxyDetailed(remote, requestExecutor, DefaultHTTPStatusHandler, DefaultHTTPResponseParserFactory(HTTPResponseParserConfig{dec, ef}))
}

// HTTPResponseParser defines how of the response is parsed from http.Response to Response object
type HTTPResponseParser func(context.Context, *http.Response) (*Response, error)

// DefaultHTTPResponseParserConfig defines a default HTTPResponseParserConfig
var DefaultHTTPResponseParserConfig = HTTPResponseParserConfig{
	func(r io.Reader, v *map[string]interface{}) error { return nil },
	EntityFormatterFunc{func(entity Response) Response { return entity }},
}

// HTTPResponseParserConfig contains the config for a given HttpResponseParser
type HTTPResponseParserConfig struct {
	dec encoding.Decoder
	ef  EntityFormatter
}

// DefaultHTTPResponseParserFactory creates HTTPResponseParser from a given HTTPResponseParserConfig
type HTTPResponseParserFactory func(HTTPResponseParserConfig) HTTPResponseParser

// DefaultHTTPResponseParserFactory is the default implementation of HTTPResponseParserFactory
func DefaultHTTPResponseParserFactory(cfg HTTPResponseParserConfig) HTTPResponseParser {
	return func(ctx context.Context, resp *http.Response) (*Response, error) {
		var data map[string]interface{}
		err := cfg.dec(resp.Body, &data)
		resp.Body.Close()

		if config.Config.Debug.Components.Servers {

			fmt.Println("proxy/http.go > DefaultHTTPResponseParserFactory(..) > var.resp.Header")
			pp.Print(resp.Header)

			fmt.Println("proxy/http.go > DefaultHTTPResponseParserFactory(..) > var.resp.StatusCode")
			pp.Println(resp.StatusCode)

			fmt.Println("proxy/http.go > DefaultHTTPResponseParserFactory(..) > var.data")
			pp.Print(data)

			if err != nil {
				fmt.Println("proxy/http.go > DefaultHTTPResponseParserFactory(..) > var.err")
				pp.Print(err)
			}

		}

		if err != nil {
			return nil, err
		}

		newResponse := Response{Data: data, IsComplete: true}
		newResponse = cfg.ef.Format(newResponse)
		return &newResponse, nil
	}
}

// HTTPStatusHandler defines how we tread the http response code
type HTTPStatusHandler func(context.Context, *http.Response) (*http.Response, error)

// DefaultHTTPCodeHandler is the default implementation of HTTPStatusHandler
func DefaultHTTPStatusHandler(ctx context.Context, resp *http.Response) (*http.Response, error) {
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, ErrInvalidStatusCode
	}

	return resp, nil
}

// NewHTTPProxyDetailed creates a http proxy with the injected configuration, HTTPRequestExecutor, Decoder and HTTPResponseParser
func NewHTTPProxyDetailed(remote *config.Backend, requestExecutor HTTPRequestExecutor, ch HTTPStatusHandler, rp HTTPResponseParser) Proxy {
	return func(ctx context.Context, request *Request) (*Response, error) {
		requestToBakend, err := http.NewRequest(request.Method, request.URL.String(), request.Body)

		if config.Config.Debug.Components.Proxies {
			if err != nil {
				fmt.Println("proxy/http.go > NewHTTPProxyDetailed(..) > NewRequest(...) > var.err")
				pp.Print(err)
			}
		}

		if err != nil {
			return nil, err
		}
		requestToBakend.Header = request.Headers

		resp, err := requestExecutor(ctx, requestToBakend)
		requestToBakend.Body.Close()

		if config.Config.Debug.Components.Proxies {

			fmt.Println("proxy/http.go > NewHTTPProxyDetailed(..) > var.request.Method")
			pp.Println(request.Method)

			fmt.Println("proxy/http.go > NewHTTPProxyDetailed(..) > var.request.Body")
			pp.Println(request.Body)

			fmt.Println("proxy/http.go > NewHTTPProxyDetailed(..) > var.request.URL.String()")
			pp.Println(request.URL.String())

			// fmt.Println("proxy/http.go > NewHTTPProxyDetailed(..) > requestToBakend.StatusCode")
			// pp.Println(requestToBakend.StatusCode)

			fmt.Println("proxy/http.go > NewHTTPProxyDetailed(..) > requestToBakend.Header")
			pp.Println(requestToBakend.Header)

			// fmt.Println("proxy/http.go > NewHTTPProxyDetailed(..) > var.resp")
			// pp.Print(resp)

			if err != nil {
				fmt.Println("proxy/http.go > NewHTTPProxyDetailed(..) > var.err")
				pp.Print(err)
			}

		}

		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
		if err != nil {
			return nil, err
		}

		resp, err = ch(ctx, resp)
		if err != nil {
			return nil, err
		}

		return rp(ctx, resp)
	}
}

// NewRequestBuilderMiddleware creates a proxy middleware that parses the request params received
// from the outter layer and generates the path to the backend endpoints
func NewRequestBuilderMiddleware(remote *config.Backend) Middleware {
	return func(next ...Proxy) Proxy {
		if len(next) > 1 {
			panic(ErrTooManyProxies)
		}
		return func(ctx context.Context, request *Request) (*Response, error) {
			r := request.Clone()
			r.AddQueryStrings(remote.QueryStrings)
			r.AddHeaders(remote.Header)
			r.GeneratePath(remote.URLPattern)
			if remote.Method != "GET" {
				r.AddParameters(remote.Parameters)
			} else {
				//r.Params.Clear()
				r.Params = make(map[string]string)
			}
			r.Method = remote.Method
			if config.Config.Debug.Components.Middlewares {
				fmt.Println("proxy/http.go > NewRequestBuilderMiddleware(..) > var.r")
				pp.Print(request.Method)
			}
			return next[0](ctx, &r)
		}
	}
}
