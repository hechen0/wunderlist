package wunderlist

import (
	"fmt"
	"net/http"
	"errors"
	"context"
)

type List struct {
	ID    *int `json:"id, omitempty"`
	Title *string `json:"title, omitempty"`
	Type  *string `json:"type, omitempty"`
}

type listService service

//Get all Lists a user has permission to
//
//GET a.wunderlist.com/api/v1/lists
//
//Response
//Status: 200
//
//json
//[
//  {
//      "id": 83526310,
//	"created_at": "2013-08-30T08:29:46.203Z",
//	"title": "Read Later",
//	"list_type": "list",
//	"type": "list",
//	"revision": 10
//  }
//]
func (s *listService) All(ctx context.Context) ([]*List, error) {
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
//Response
//
//Status: 200
//
//json
//{
//"id": 83526310,
//"created_at": "2013-08-30T08:29:46.203Z",
//"title": "Read Later",
//"list_type": "list",
//"type": "list",
//"revision": 10
//}
func (s *listService) Get(ctx context.Context, id int) (*List, error) {
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

//Create a list
//
//POST a.wunderlist.com/api/v1/lists
//Data
//
//NAME	TYPE	NOTES
//title	string	required. maximum length is 255 characters
//Request body example
//
//json
//{
//"title": "Hallo"
//}
//Response
//
//Status: 201
//
//json
//{
//"id": 83526310,
//"created_at": "2013-08-30T08:29:46.203Z",
//"title": "Read Later",
//"revision": 1000,
//"type": "list"
//}
func (s *listService) Create(ctx context.Context, list *List) (*List, error) {
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
//Data
//
//NAME	TYPE	NOTES
//revision	integer	required
//title	string	maximum length is 255 characters
//Request body example
//
//json
//{
//"revision": 1000,
//"title": "Hallo"
//}
//Response
//
//Status 200
//
//json
//{
//"id": 409233670,
//"revision": 1001,
//"title": "Hello",
//"type": "list"
//}
func (s *listService) Update() (error) {
	return nil
}

//Delete a list permanently
//
//DELETE a.wunderlist.com/api/v1/lists/:id
//Params
//
//NAME	TYPE	NOTES
//revision	integer	required
//Response
//
//Status 204
func (s *listService) Delete(id int) (err error) {
	return
}
