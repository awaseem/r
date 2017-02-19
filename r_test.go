package r

import "testing"
import "net/http"
import "net/http/httptest"

func TestSetHeader(t *testing.T) {
	r := New().SetHeader("hello", "world")
	if r.Headers["hello"] != "world" {
		t.Errorf("Failed to set headers for request object")
	}
}

func TestSetMethod(t *testing.T) {
	r := New().SetMethod("PATCH")
	if r.Method != "PATCH" {
		t.Errorf("Failed to set method for request object")
	}
}

func TestSetURL(t *testing.T) {
	r := New().SetURL("https://google.com")
	if r.URL != "https://google.com" {
		t.Errorf("Failed to set URL for request object")
	}
}

func TestSetBody(t *testing.T) {
	r := New().SetBody([]byte("body"))
	if r.Body.String() != "body" {
		t.Errorf("Failed to set body for request object")
	}
}

func TestSetBodyString(t *testing.T) {
	r := New().SetBodyString("body")
	if r.Body.String() != "body" {
		t.Errorf("Failed to set body for request object")
	}
}

func TestSetParam(t *testing.T) {
	r := New().SetParam("hello", "world")
	if r.QueryParams["hello"] != "world" {
		t.Errorf("Failed to set query parameter")
	}
}

func TestDefault(t *testing.T) {
	r := New().Default()
	if r.Headers["Content-Type"] != "application/json" {
		t.Errorf("Failed to set default json header")
	}
}

func TestGet(t *testing.T) {
	r := New().Get("https://google.com")
	if r.URL != "https://google.com" {
		t.Errorf("Failed to set URL for request object")
	}
	if r.Method != "GET" {
		t.Errorf("Failed to set GET method for request object")
	}
	if r.Headers["Content-Type"] != "application/json" {
		t.Errorf("Failed to set default json header for request object")
	}
}

func TestPost(t *testing.T) {
	r := New().Post("https://google.com")
	if r.URL != "https://google.com" {
		t.Errorf("Failed to set URL for request object")
	}
	if r.Method != "POST" {
		t.Errorf("Failed to set GET method for request object")
	}
	if r.Headers["Content-Type"] != "application/json" {
		t.Errorf("Failed to set default json header for request object")
	}
}

func TestPut(t *testing.T) {
	r := New().Put("https://google.com")
	if r.URL != "https://google.com" {
		t.Errorf("Failed to set URL for request object")
	}
	if r.Method != "PUT" {
		t.Errorf("Failed to set GET method for request object")
	}
	if r.Headers["Content-Type"] != "application/json" {
		t.Errorf("Failed to set default json header for request object")
	}
}

func TestDelete(t *testing.T) {
	r := New().Delete("https://google.com")
	if r.URL != "https://google.com" {
		t.Errorf("Failed to set URL for request object")
	}
	if r.Method != "DELETE" {
		t.Errorf("Failed to set GET method for request object")
	}
	if r.Headers["Content-Type"] != "application/json" {
		t.Errorf("Failed to set default json header for request object")
	}
}

func TestSetQueryParams(t *testing.T) {
	url := "www.test.com"
	params := make(map[string]string)
	params["hello"] = "world"
	r, err := setQueryParams(url, params)
	if err != nil {
		t.Error("errored query params test case when its not suppose to...")
	}
	if r != "www.test.com?hello=world" {
		t.Errorf("Failed to set query parameters")
	}
}

func TestSetHeaders(t *testing.T) {
	httpR, _ := http.NewRequest("POST", "test", nil)
	headers := make(map[string]string)
	headers["hello"] = "world"
	setHeaders(httpR, headers)
	if httpR.Header.Get("hello") != "world" {
		t.Errorf("Failed set headers for request object")
	}
}

func TestGetRequest(t *testing.T) {
	// setup
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("ok"))
		if err != nil {
			t.Errorf("Failed to write response")
		}
	}))
	defer server.Close()
	r, _ := New().Get(server.URL).Send()
	if r.StatusCode != 200 {
		t.Error("Failed to request data from httpbin")
	}
}

func TestParseJSON(t *testing.T) {
	// setup
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(`{"origin": "ok"}`))
		if err != nil {
			t.Errorf("Failed to write response")
		}
	}))
	defer server.Close()
	type res struct {
		Origin string `json:"origin"`
	}
	resB := res{}
	r, _ := New().Get(server.URL).Send()
	ParseJSON(r, &resB)
	if resB.Origin == "" {
		t.Errorf("Failed to parse httpbin resposne")
	}
}
