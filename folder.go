package wunderlist

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

//https://developer.wunderlist.com/documentation/endpoints/folder
type FolderService service

type Folder struct {
	Id                 *int    `json:"id,omitempty"`
	CreatedAt          *string `json:"created_at,omitempty"`
	CreatedById        *int    `json:"created_by_id,omitempty"`
	CreatedByRequestId *string `json:"created_by_request_id,omitempty"`
	ListIds            []*int  `json:"list_ids,omitempty"`
	Revision           *int    `json:"revision,omitempty"`
	Title              *string `json:"title,omitempty"`
	Type               *string `json:"type,omitempty"`
	UpdatedAt          *string `json:"updated_at,omitempty"`
	UserId             *int    `json:"user_id,omitempty"`
}

type FolderRevision struct {
	Id       *string `json:"id,omitempty"`
	Revision *string `json:"revision,omitempty"`
	Type     *string `json:"type,omitempty"`
}

//Get all Folders created by the the current User
//
//GET a.wunderlist.com/api/v1/folders
//
func (s *FolderService) All(ctx context.Context) ([]*Folder, error) {
	u := fmt.Sprintf("folders")
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var folders []*Folder
	_, err = s.client.Do(ctx, req, &folders)
	if err != nil {
		return nil, err
	}
	return folders, nil
}

//Get a specific Folder
//
//GET a.wunderlist.com/api/v1/folders/:id
//
func (s *FolderService) Get(ctx context.Context, id int) (*Folder, error) {
	u := fmt.Sprintf("folders/%v", id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	folder := new(Folder)
	_, err = s.client.Do(ctx, req, folder)
	if err != nil {
		return nil, err
	}
	return folder, nil
}

//
//Create a Folder
//
//POST a.wunderlist.com/api/v1/folders
//
func (s *FolderService) Create(ctx context.Context, folder *Folder) (*Folder, error) {
	u := "folders"
	req, err := s.client.NewRequest("POST", u, folder)
	if err != nil {
		return nil, err
	}

	f := new(Folder)

	resp, err := s.client.Do(ctx, req, f)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusCreated {
		return nil, errors.New(fmt.Sprintf("expected status: %v, got: %v", http.StatusCreated, resp.Status))
	}

	return f, nil
}

//Update a Folder by overwriting properties
//
//PATCH a.wunderlist.com/api/v1/folders/:id
//
func (s *FolderService) Update(ctx context.Context, folder *Folder) (*Folder, error) {
	u := fmt.Sprintf("folders/%v", folder.Id)
	req, err := s.client.NewRequest("PATCH", u, folder)
	if err != nil {
		return nil, err
	}

	f := new(Folder)
	_, err = s.client.Do(ctx, req, f)
	if err != nil {
		return nil, err
	}

	return f, nil
}

//Delete a Folder permanently
//
//DELETE a.wunderlist.com/api/v1/folders/:id
//
func (s *FolderService) Delete(ctx context.Context, id int) error {
	u := fmt.Sprintf("folders/%v", id)
	req, err := s.client.NewRequest("DELETE", u, nil)
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

//
//Get Folder Revisions
//
//GET a.wunderlist.com/api/v1/folder_revisions
//
func (s *FolderService) FolderRevisions(ctx context.Context) ([]*FolderRevision, error) {
	u := fmt.Sprintf("folder_revisions")
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var revisions []*FolderRevision
	_, err = s.client.Do(ctx, req, revisions)
	if err != nil {
		return nil, err
	}
	return revisions, nil
}
