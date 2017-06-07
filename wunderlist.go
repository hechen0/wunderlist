package wunderlist

import (
	"net/url"
	"net/http"
	"bufio"
	"io"
	"encoding/json"
)

const (
	version   = "1"
	baseURL   = "https://a.wunderlist.com/"
	userAgent = "hechen0/wunderlist " + version
)

type Client struct {
	client *http.Client

	// Base URL for API requests.
	BaseURL *url.URL

	// User agent used when communicating with the API.
	UserAgent string

	// Reuse a single struct instead of allocating one for each service on the heap.
	share service

	// Services used for talking to different parts of the API.
	Lists *listService

	// access token
	token string
}

type service struct {
	client *Client
}

func NewClient() *Client {

	base, _ := url.Parse(baseURL)

	c := &Client{UserAgent: userAgent, BaseURL: base, client: http.DefaultClient}
	c.share.client = c

	c.Lists = (*listService)(&c.share)
	return c
}

func (c *Client) SetToken(token string) {
	c.token = token
}

func (c *Client) NewRequest(method, urlStr string, body interface{}) (req *http.Request, err error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return
	}

	u := c.BaseURL.ResolveReference(rel)

	var buf io.ReadWriter

	if body != nil {
		buf = new(bufio.ReadWriter)

		// check request body valid
		if err = json.NewEncoder(buf).Encode(body); err != nil {
			return
		}
	}

	req, err = http.NewRequest(method, u.String(), buf)
	if err != nil {
		return
	}

	return
}

func (c *Client) Do(req *http.Request, v interface{}) (resp *http.Response, err error) {
	resp, err = c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if v != nil {
		if err = json.NewDecoder(resp.Body).Decode(v); err != nil {
			return nil, err
		}
	}

	return
}

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool { return &v }

// Int is a helper routine that allocates a new int value
// to store v and returns a pointer to it.
func Int(v int) *int { return &v }

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string { return &v }
