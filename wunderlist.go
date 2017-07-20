package wunderlist

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	baseURL = "https://a.wunderlist.com/api/v1/"
)

type Client struct {
	client *http.Client

	// Base URL for API requests.
	BaseURL *url.URL

	// Client ID and access token
	Auth *Auth

	// Reuse a single struct instead of allocating one for each service on the heap.
	share service

	// Services used for talking to different parts of the API.
	Lists   *ListService
	Folders *FolderService
	Avatar  *AvatarService
}

type Auth struct {
	Token    string
	ClientId string
}

type service struct {
	client *Client
}

func NewClient() *Client {

	base, _ := url.Parse(baseURL)

	c := &Client{BaseURL: base, client: http.DefaultClient}
	c.share.client = c

	c.Lists = (*ListService)(&c.share)
	c.Folders = (*FolderService)(&c.share)
	c.Avatar = (*AvatarService)(&c.share)

	return c
}

func (c *Client) SetAuth(auth *Auth) {
	c.Auth = auth
}

func (c *Client) NewRequest(method, urlStr string, body interface{}) (req *http.Request, err error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return
	}

	u := c.BaseURL.ResolveReference(rel)

	var buf io.ReadWriter

	if body != nil {
		buf = new(bytes.Buffer)

		if err = json.NewEncoder(buf).Encode(body); err != nil {
			return
		}
	}

	req, err = http.NewRequest(method, u.String(), buf)
	if err != nil {
		return
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	return
}

func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {

	req = req.WithContext(ctx)

	req.Header.Set("X-Access-Token", c.Auth.Token)
	req.Header.Set("X-Client-ID", c.Auth.ClientId)

	resp, err := c.client.Do(req)

	if err != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
	}

	if v != nil {
		if err = json.NewDecoder(resp.Body).Decode(v); err != nil {
			return nil, err
		}
	}

	return resp, nil
}

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
//func Bool(v bool) *bool { return &v }

// Int is a helper routine that allocates a new int value
// to store v and returns a pointer to it.
func Int(v int) *int { return &v }

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string { return &v }

// Get task or list
func ParseTaskOrList(task_or_list interface{}) (id int, t string, err error) {
	switch task_or_list.(type) {
	case Task:
		return *Task(task_or_list).Id, "task_id", nil
	case List:
		return *List(task_or_list).Id, "list_id", nil
	default:
		return nil, nil, errors.New(fmt.Sprintf("expect Task or List, got: %v", task_or_list))
	}
}
