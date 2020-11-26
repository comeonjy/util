// Licensed to Elasticsearch B.V under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.
//
// Code generated from specification version 8.0.0: DO NOT EDIT

package esapi

import (
	"context"
	"net/http"
	"strings"
)

func newIndicesResolveIndexFunc(t Transport) IndicesResolveIndex {
	return func(name []string, o ...func(*IndicesResolveIndexRequest)) (*Response, error) {
		var r = IndicesResolveIndexRequest{Name: name}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// IndicesResolveIndex returns information about any matching indices, aliases, and data streams
//
// This API is experimental.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/master/indices-resolve-index-api.html.
//
type IndicesResolveIndex func(name []string, o ...func(*IndicesResolveIndexRequest)) (*Response, error)

// IndicesResolveIndexRequest configures the Indices Resolve Index API request.
//
type IndicesResolveIndexRequest struct {
	Name []string

	ExpandWildcards string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r IndicesResolveIndexRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "GET"

	path.Grow(1 + len("_resolve") + 1 + len("index") + 1 + len(strings.Join(r.Name, ",")))
	path.WriteString("/")
	path.WriteString("_resolve")
	path.WriteString("/")
	path.WriteString("index")
	path.WriteString("/")
	path.WriteString(strings.Join(r.Name, ","))

	params = make(map[string]string)

	if r.ExpandWildcards != "" {
		params["expand_wildcards"] = r.ExpandWildcards
	}

	if r.Pretty {
		params["pretty"] = "true"
	}

	if r.Human {
		params["human"] = "true"
	}

	if r.ErrorTrace {
		params["error_trace"] = "true"
	}

	if len(r.FilterPath) > 0 {
		params["filter_path"] = strings.Join(r.FilterPath, ",")
	}

	req, err := newRequest(method, path.String(), nil)
	if err != nil {
		return nil, err
	}

	if len(params) > 0 {
		q := req.URL.Query()
		for k, v := range params {
			q.Set(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	if len(r.Header) > 0 {
		if len(req.Header) == 0 {
			req.Header = r.Header
		} else {
			for k, vv := range r.Header {
				for _, v := range vv {
					req.Header.Add(k, v)
				}
			}
		}
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	res, err := transport.Perform(req)
	if err != nil {
		return nil, err
	}

	response := Response{
		StatusCode: res.StatusCode,
		Body:       res.Body,
		Header:     res.Header,
	}

	return &response, nil
}

// WithContext sets the request context.
//
func (f IndicesResolveIndex) WithContext(v context.Context) func(*IndicesResolveIndexRequest) {
	return func(r *IndicesResolveIndexRequest) {
		r.ctx = v
	}
}

// WithExpandWildcards - whether wildcard expressions should get expanded to open or closed indices (default: open).
//
func (f IndicesResolveIndex) WithExpandWildcards(v string) func(*IndicesResolveIndexRequest) {
	return func(r *IndicesResolveIndexRequest) {
		r.ExpandWildcards = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f IndicesResolveIndex) WithPretty() func(*IndicesResolveIndexRequest) {
	return func(r *IndicesResolveIndexRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f IndicesResolveIndex) WithHuman() func(*IndicesResolveIndexRequest) {
	return func(r *IndicesResolveIndexRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f IndicesResolveIndex) WithErrorTrace() func(*IndicesResolveIndexRequest) {
	return func(r *IndicesResolveIndexRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f IndicesResolveIndex) WithFilterPath(v ...string) func(*IndicesResolveIndexRequest) {
	return func(r *IndicesResolveIndexRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
//
func (f IndicesResolveIndex) WithHeader(h map[string]string) func(*IndicesResolveIndexRequest) {
	return func(r *IndicesResolveIndexRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
//
func (f IndicesResolveIndex) WithOpaqueID(s string) func(*IndicesResolveIndexRequest) {
	return func(r *IndicesResolveIndexRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
