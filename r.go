package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const (
	getMethod    = "GET"
	postMethod   = "POST"
	putMethod    = "PUT"
	deleteMethod = "DELETE"
)

// Request is the baseline object for the request
type Request struct {
	// Headers for the request
	Headers map[string]string
	// Type for the request (i.e GET, POST, PUT, DELETE)
	Method string
	// URL for the request
	URL string
	// Body for the request
	Body *bytes.Buffer
	// Query parameters for the url
	QueryParams map[string]string
}

// New create a new request object with no defaults
func New() *Request {
	r := Request{}
	r.Headers = make(map[string]string)
	r.QueryParams = make(map[string]string)
	r.Body = new(bytes.Buffer)
	return &r
}

// SetHeader set a single header for the request
func (r *Request) SetHeader(key string, value string) *Request {
	r.Headers[key] = value
	return r
}

// SetMethod set the request type
func (r *Request) SetMethod(method string) *Request {
	r.Method = method
	return r
}

// SetURL set the request URL
func (r *Request) SetURL(URL string) *Request {
	r.URL = URL
	return r
}

// SetBody set the request body (defaults to SetBodyBytes)
func (r *Request) SetBody(body []byte) *Request {
	return r.SetBodyBytes(body)
}

// SetBodyBytes set the request body as an array of bytes
func (r *Request) SetBodyBytes(body []byte) *Request {
	r.Body = bytes.NewBuffer(body)
	return r
}

// SetBodyString set the body as a string
func (r *Request) SetBodyString(body string) *Request {
	r.Body = bytes.NewBuffer([]byte(body))
	return r
}

// SetBodyBuffer set the request body as a buffer object
func (r *Request) SetBodyBuffer(body *bytes.Buffer) *Request {
	r.Body = body
	return r
}

// SetParam set query parameters
func (r *Request) SetParam(key string, value string) *Request {
	r.QueryParams[key] = value
	return r
}

// Default set any defaults
func (r *Request) Default() *Request {
	r.SetHeader("Content-Type", "application/json")
	return r
}

// Get helper method for GET requests
func (r *Request) Get(url string) *Request {
	return r.SetMethod(getMethod).SetURL(url).Default()
}

// Post helper method for POST requests
func (r *Request) Post(url string) *Request {
	return r.SetMethod(postMethod).SetURL(url).Default()
}

// Put helper method for PUT requests
func (r *Request) Put(url string) *Request {
	return r.SetMethod(putMethod).SetURL(url).Default()
}

// Delete helper method for DELETE requests
func (r *Request) Delete(url string) *Request {
	return r.SetMethod(deleteMethod).SetURL(url).Default()
}

// Send method to send the request
func (r *Request) Send() (*http.Response, error) {
	req, err := http.NewRequest(r.Method, r.URL, r.Body)
	if err != nil {
		return nil, err
	}
	setHeaders(req, r.Headers)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return resp, nil
}

// SendJSON method to send json payloads, it takes a generic interface and transforms it into JSON
func (r *Request) SendJSON(body interface{}) (*http.Response, error) {
	b := new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(body); err != nil {
		return nil, err
	}
	return r.SetBodyBuffer(b).Send()
}

func setHeaders(r *http.Request, headers map[string]string) {
	for k, v := range headers {
		r.Header.Set(k, v)
	}
}
