package wunderlist

import "net/url"

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

	baseurl, _ := url.Parse(baseURL)

	c := &Client{UserAgent: userAgent, BaseURL: baseurl}
	c.share = service{}

	c.Lists = (*listService)(&c.share)
	return c
}