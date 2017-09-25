package proxy

import (
	"bytes"
	// "fmt"
	// "github.com/k0kubun/pp"
	"io"
	"net/url"
	"strings"
	// "github.com/roscopecoltran/krakend/logging"
)

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
	for key, value := range queryStrings {
		if r.Params[key] != "" {
			value = r.Params[key]
		}
		r.Query.Add(key, value)
	}
}

func (r *Request) AddParameters(parameters map[string]string) {
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
