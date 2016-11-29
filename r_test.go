package main

import "testing"

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
