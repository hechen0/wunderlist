package wunderlist

import (
	"context"
	"fmt"
	"errors"
)

//TODO: preview broken
type FilePreviewService service

type FilePreview struct {
	Id  *int `json:"id,omitempty"`
	Url *string `json:"url,omitempty"`
}

var SupportFilePreviewType = []string{
	"image/gif",
	"image/jpeg",
	"image/jpg",
	"image/pjpeg",
	"image/png",
	"image/svg+xml",
	"image/tiff",
}

//
// Get a Preview of a specific File
//
func (s *FilePreviewService) Get(ctx context.Context, file *File, size, platform *string) (*FilePreview, error) {

	u := fmt.Sprintf("previews")

	canPreview := false
	for _, support := range SupportFilePreviewType {
		if support == *file.ContentType {
			canPreview = true
		}
	}
	if !canPreview {
		return nil, errors.New(fmt.Sprintf("preview only support type: %+v, provided: %v",
			SupportFilePreviewType, *file.ContentType))
	}

	params := struct {
		FileId   *int `json:"file_id,omitempty"`
		Platform *string `json:"platform,omitempty"`
		Size     *string `json:"size,omitempty"`
	}{file.Id, platform, size}

	req, err := s.client.NewRequest("GET", u, params)
	if err != nil {
		return nil, err
	}

	newFilePreview := new(FilePreview)
	_, err = s.client.Do(ctx, req, newFilePreview)
	if err != nil {
		return nil, err
	}

	return newFilePreview, nil
}
