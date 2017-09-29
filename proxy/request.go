package proxy

import (
	"bytes"
	"io"
	"net/url"
	"strings"

	"fmt"
	"github.com/k0kubun/pp"

	"github.com/roscopecoltran/krakend/config"
)

/*
	Refs:
	- https://github.com/eidge/yurl/blob/master/examples/users.yml
	-
*/

// Request contains the data to send to the backend
type Request struct {
	Method  string
	URL     *url.URL
	Query   url.Values
	Path    string
	Body    io.ReadCloser
	Params  map[string]string
	Headers map[string][]string
}

// GeneratePath takes a pattern and updates the path of the request
func (r *Request) GeneratePath(URLPattern string) {
	if len(r.Params) == 0 {
		r.Path = URLPattern
		return
	}
	buff := []byte(URLPattern)
	for k, v := range r.Params {
		key := []byte{}
		key = append(key, "{{."...)
		key = append(key, k...)
		key = append(key, "}}"...)
		buff = bytes.Replace(buff, key, []byte(v), -1)
	}
	r.Path = string(buff)
}

func (r *Request) AddQueryStrings(queryStrings map[string]string) {
	if len(queryStrings) == 0 {
		return
	}
	r.Query = make(url.Values)
	for key, value := range queryStrings {
		paramKey := strings.ToTitle(key)
		if config.Config.Debug.Components.Middlewares {
			fmt.Printf("proxy/request.go > AddQueryStrings(...) > var.queryStrings > key=%s, value=%s , paramKey=%s, r.Params[paramKey]=%s \n", key, value, paramKey, r.Params[paramKey])
		}
		if r.Params[paramKey] != "" {
			value = r.Params[paramKey]
		}
		r.Query.Add(key, value)
	}
	if config.Config.Debug.Components.Middlewares {
		fmt.Println("proxy/request.go > AddQueryStrings(...) > var.r.Query")
		pp.Print(r.Query)
		fmt.Println("proxy/request.go > AddQueryStrings(...) > var.r.Params")
		pp.Print(r.Params)
	}
}

func (r *Request) AddParameters(parameters map[string]string) {
	if r.Method == "GET" {
		r.Params = make(map[string]string)
	}
	if len(parameters) == 0 {
		return
	}
	for key, value := range parameters {
		if r.Params[key] == "" {
			r.Params[key] = strings.Trim(string(value), " ")
		}
	}
}

func (r *Request) AddHeaders(headers []string) {
	if len(headers) == 0 {
		return
	}
	for _, header := range headers {
		parts := strings.Split(header, ":")
		if len(parts) == 2 {
			name := parts[0]
			value := parts[1]
			r.Headers[name] = []string{strings.Trim(value, " ")}
		}
	}
}

// Clone clones itself into a new request
func (r *Request) Clone() Request {
	return Request{
		Method:  r.Method,
		URL:     r.URL,
		Query:   r.Query,
		Path:    r.Path,
		Body:    r.Body,
		Params:  r.Params,
		Headers: r.Headers,
	}
}
