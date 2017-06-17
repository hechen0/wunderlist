package wunderlist

import (
	"fmt"
	"net/http"
	"errors"
	"context"
)

type List struct {
	ID    *int `json:"id,omitempty"`
	Title *string `json:"title,omitempty"`
	Type  *string `json:"type,omitempty"`
}

type ListService service

//Get all Lists a user has permission to
//
//GET a.wunderlist.com/api/v1/lists
//
func (s *ListService) All(ctx context.Context) ([]*List, error) {
	req, err := s.client.NewRequest("GET", "lists", nil)
	if err != nil {
		return nil, err
	}

	var lists []*List

	_, err = s.client.Do(ctx, req, &lists)
	if err != nil {
		return nil, err
	}

	return lists, nil
}

//Get a specific List
//
//GET a.wunderlist.com/api/v1/lists/:id
//
func (s *ListService) Get(ctx context.Context, id int) (*List, error) {
	u := fmt.Sprintf("lists/%v", id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	list := new(List)
	_, err = s.client.Do(ctx, req, list)
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (s *ListService) Create(ctx context.Context, list *List) (*List, error) {
	u := "lists"
	req, err := s.client.NewRequest("POST", u, list)
	if err != nil {
		return nil, err
	}

	newList := new(List)

	resp, err := s.client.Do(ctx, req, newList)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusCreated {
		return nil, errors.New(fmt.Sprintf("expected status: %v, got: %v", http.StatusCreated, resp.Status))
	}

	return newList, nil
}

//
//Update a list by overwriting properties
//
//PATCH a.wunderlist.com/api/v1/lists/:id
func (s *ListService) Update(ctx context.Context, list *List) (*List, error) {
	u := fmt.Sprintf("lists/%d", list.ID)
	req, err := s.client.NewRequest("PATCH", u, list)
	if err != nil {
		return nil, err
	}

	l := new(List)
	_, err = s.client.Do(ctx, req, l)
	if err != nil {
		return nil, err
	}

	return l, nil
}

//Delete a list permanently
//
//DELETE a.wunderlist.com/api/v1/lists/:id
func (s *ListService) Delete(ctx context.Context, id int) (error) {
	u := fmt.Sprintf("lists/%v", id)
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
