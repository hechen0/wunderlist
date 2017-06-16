package wunderlist

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"net/url"
)

var (
	mux *http.ServeMux

	client *Client

	server *httptest.Server
)

func setup(){
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	client = NewClient()

	u, _ := url.Parse(server.URL)
	client.BaseURL = u
}

func teardown(){
	server.Close()
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

func TestNewClient(t *testing.T) {
	c := NewClient()

	if got, want := c.BaseURL.String(), baseURL; got != want {
		t.Errorf("NewClient BaseURL is %v, want %v", got, want)
	}

	if got, want := c.UserAgent, userAgent; got != want {
		t.Errorf("NewClient UserAgent is %v, want %v", got, want)
	}
}