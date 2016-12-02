<p align="center">
  <img src="http://i.imgur.com/jv4RVHD.png">
</p>

# r
[![Build Status](https://travis-ci.org/awaseem/r.svg?branch=master)](https://travis-ci.org/awaseem/r)
[![Go Report Card](https://goreportcard.com/badge/github.com/awaseem/r)](https://goreportcard.com/report/github.com/awaseem/r)

Easiest way to make requests.

# Getting Started

```
go get github.com/awaseem/r
```

# Documentation

See [Go Doc](https://godoc.org/github.com/awaseem/r) for more informations on methods

# Examples 

The general idea is to build the request and use the `send()` method to get the response.

If you want to make a simple GET request, here's what you can do:

```
res, err := r.New().Get("http://example.com/").Send()
```

The response you get back is just `http.Response`. If you want to do a simple POST request with a json payload, here's what you can do: 

```
// custom struct type for body of request
body := Body {
  Value: "test",
}
// pass the struct as a parameter to sendJSON
res, err := r.New().Post("http://example.com/").SendJSON(body)
```

There is also a helper method that allows you to take the body of a response and put it into a struct

```
body := Body {}
res, _ := r.New().Post("http://example.com/").Send()

// body will hold the values from the JSON response
err := r.ParseJSON(res, &body)
```  

# Defaults

The following is the only header is set as default `"Content-Type" = "application/json"`