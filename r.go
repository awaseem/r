package r

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
	Body []byte
	// Query parameters for the url
	QueryParams map[string]string
}

// New create a new request object with no defaults
func New() *Request {
	r := Request{}
	r.Headers = make(map[string]string)
	r.QueryParams = make(map[string]string)
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

// SetBody set the request body
func (r *Request) SetBody(body []byte) *Request {
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
