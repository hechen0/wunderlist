package wunderlist

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

type AvatarService service

type Image struct {
	Url      *string
	Size     *string
	Fallback *bool
}

//User avatars of different sizes can be fetched or loaded inline in HTML
//https://developer.wunderlist.com/documentation/endpoints/avatar
func (s *AvatarService) Get(ctx context.Context, image *Image) (*Image, error) {
	u := fmt.Sprintf("/avatar")
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusFound {
		return nil, errors.New(fmt.Sprintf("expected AvatarService.Get return status: %v, got: %v", http.StatusFound, resp.Status))
	}

	location := resp.Header.Get("Location")

	if location == "" {
		return nil, errors.New("expected AvatarService.Get return image location got empty")
	}

	image.Url = String(location)

	return image, nil
}
