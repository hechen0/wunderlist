package wunderlist

import (
	"fmt"
	"net/http"
	"errors"
	"context"
)

// https://developer.wunderlist.com/documentation/endpoints/list
type TaskService service

type Task struct {
	Id                 *int `json:"id,omitempty"`
	Title              *string `json:"title,omitempty"`
	Type               *string `json:"type,omitempty"`
	Completed          *bool `json:"completed,omitempty"`
	CreatedAt          *string `json:"created_at,omitempty"`
	CreatedById        *int `json:"created_by_id,omitempty"`
	CreatedByRequestId *string `json:"created_by_request_id,omitempty"`
	ListId             *int `json:"list_id,omitempty"`
	Revision           *int `json:"revision,omitempty"`
	Starred            *bool `json:"starred,omitempty"`
}

// Get all Lists a user has permission to
//
// GET a.wunderlist.com/api/v1/lists
//
func (s *TaskService) All(ctx context.Context) ([]*Task, error) {
	req, err := s.client.NewRequest("GET", "tasks", nil)
	if err != nil {
		return nil, err
	}

	var tasks []*Task

	_, err = s.client.Do(ctx, req, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

// Get a specific List
//
// GET a.wunderlist.com/api/v1/lists/:id
//
func (s *TaskService) Get(ctx context.Context, id int) (*Task, error) {
	u := fmt.Sprintf("tasks/%v", id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	task := new(Task)
	_, err = s.client.Do(ctx, req, task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

//
// Create a list
//
func (s *TaskService) Create(ctx context.Context, task *Task) (*Task, error) {
	u := "tasks"
	req, err := s.client.NewRequest("POST", u, task)
	if err != nil {
		return nil, err
	}

	newTask := new(Task)

	resp, err := s.client.Do(ctx, req, newTask)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusCreated {
		return nil, errors.New(fmt.Sprintf("expected status: %v, got: %v", http.StatusCreated, resp.Status))
	}

	return newTask, nil
}

//
//Update a list by overwriting properties
//
func (s *TaskService) Update(ctx context.Context, task *Task) (*Task, error) {
	u := fmt.Sprintf("tasks/%d", task.Id)
	req, err := s.client.NewRequest("PATCH", u, task)
	if err != nil {
		return nil, err
	}

	l := new(Task)
	_, err = s.client.Do(ctx, req, l)
	if err != nil {
		return nil, err
	}

	return l, nil
}

//
//Delete a list permanently
//
func (s *TaskService) Delete(ctx context.Context, id int) (error) {
	u := fmt.Sprintf("tasks/%v", id)
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
