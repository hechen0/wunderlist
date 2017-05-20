package wunderlist

import (
	"net/url"
	"net/http"
	"bufio"
	"io"
	"encoding/json"
)

const (
	version = "1"
	baseURL = "https://a.wunderlist.com/"
	userAgent = "hechen0/wunderlist " + version
)

type Client struct {

	// Base URL for API requests.
	BaseURL *url.URL

	// User agent used when communicating with the API.
	UserAgent string

	// Reuse a single struct instead of allocating one for each service on the heap.
	share service

	// Services used for talking to different parts of the API.
	Lists *listService
}


type service struct {
	client *Client
}

func NewClient() *Client{

	base, _ := url.Parse(baseURL)

	c := &Client{UserAgent: userAgent, BaseURL: base}
	c.share.client = c

	c.Lists = (*listService)(&c.share)
	return c
}

func (c *Client) NewRequest(method, urlStr string, body interface{}) (req *http.Request,err error){
	rel, err := url.Parse(urlStr)
	if err != nil {
		return
	}

	u := c.BaseURL.ResolveReference(rel)

	var buf io.ReadWriter

	if body != nil {
		buf = new(bufio.ReadWriter)

		// check request body valid
		if err = json.NewEncoder(buf).Encode(body); err != nil{
			return
		}
	}

	req, err = http.NewRequest(method, u.String(), buf)
	if err != nil {
		return
	}

	return
}