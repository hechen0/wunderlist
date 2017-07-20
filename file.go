package wunderlist

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

// https://developer.wunderlist.com/documentation/endpoints/file
type FileService service

type File struct {
	ContentType        *string      `json:"content_type,omitempty"`
	CreatedAt          *string      `json:"created_at,omitempty"`
	CreatedByRequestId *string      `json:"created_by_request_id,omitempty"`
	FileName           *string      `json:"file_name,omitempty"`
	FileSize           *json.Number `json:"file_size,omitempty"`
	Id                 *int         `json:"id,omitempty"`
	LocalCreatedAt     *string      `json:"local_created_at,omitempty"`
	Revision           *int         `json:"revision,omitempty"`
	TaskId             *int         `json:"task_id,omitempty"`
	Type               *string      `json:"type,omitempty"`
	UpdatedAt          *string      `json:"updated_at,omitempty"`
	Url                *string      `json:"url,omitempty"`
	UserId             *int         `json:"user_id,omitempty"`
}

//
// Get Files for a Task or List
//
func (s *FileService) All(ctx context.Context, task_or_list interface{}) ([]*File, error) {
	t, id, err := ParseTaskOrList(task_or_list)

	if err != nil {
		return nil, err
	}

	u := fmt.Sprintf("files?%v=%v", t, id)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var files []*File

	_, err = s.client.Do(ctx, req, &files)
	if err != nil {
		return nil, err
	}

	return files, nil
}

//
// Get a specific File
//
func (s *FileService) Get(ctx context.Context, id int) (*File, error) {
	u := fmt.Sprintf("files/%v", id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	file := new(File)
	_, err = s.client.Do(ctx, req, file)
	if err != nil {
		return nil, err
	}

	return file, nil
}

//
// Create a file
//
func (s *FileService) Create(ctx context.Context, task *Task, upload *Upload) (*File, error) {
	u := "files"

	timezone, _ := time.LoadLocation("Asia/Shanghai")
	params := struct {
		UploadId       *int   `json:"upload_id"`
		TaskId         *int   `json:"task_id"`
		LocalCreatedAt string `json:"local_created_at"`
	}{upload.Id, task.Id, time.Now().In(timezone).String()}

	req, err := s.client.NewRequest("POST", u, params)
	if err != nil {
		return nil, err
	}

	newFile := new(File)

	resp, err := s.client.Do(ctx, req, newFile)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusCreated {
		return nil, errors.New(fmt.Sprintf("expected status: %v, got: %v", http.StatusCreated, resp.Status))
	}

	return newFile, nil
}

//
//Delete a file permanently
//
func (s *FileService) Delete(ctx context.Context, file *File) error {
	u := fmt.Sprintf("files/%v", file.Id)
	req, err := s.client.NewRequest("DELETE", u, file)
	if err != nil {
		return err
	}
	resp, err := s.client.Do(ctx, req, nil)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return errors.New(fmt.Sprintf("expected status: %v, got: %v", http.StatusNoContent, resp.Status))
	}

	return nil
}
