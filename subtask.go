package wunderlist

import (
	"fmt"
	"net/http"
	"errors"
	"context"
)

// https://developer.wunderlist.com/documentation/endpoints/subtask
type SubtaskService service

type Subtask struct {
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
	DueDate            *string `json:"due_date,omitempty"`
	RecurrenceCount    *string `json:"recurrence_count,omitempty"`
	RecurrenceType     *string `json:"recurrence_type,omitempty"`
	CompletedAt        *string `json:"completed_at,omitempty"`
	CompletedById      *string `json:"completed_by_id,omitempty"`
	AssigneeId         *string `json:"assignee_id,omitempty"`
}

//
// Get Subtasks for a Task or List
//
func (s *SubtaskService) All(ctx context.Context, task_or_list interface{}) ([]*Subtask, error) {
	t, id, err := ParseTaskOrList(task_or_list)

	if err != nil {
		return nil, err
	}

	u := fmt.Sprintf("subtasks?%v=%v", t, id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var subtasks []*Subtask

	_, err = s.client.Do(ctx, req, &subtasks)
	if err != nil {
		return nil, err
	}

	return subtasks, nil
}

//
//Get Completed Subtasks
//
func (s *SubtaskService) Completed(ctx context.Context, task_or_list interface{}, completed bool) ([]*Subtask, error) {
	t, id, err := ParseTaskOrList(task_or_list)

	if err != nil {
		return nil, err
	}

	u := fmt.Sprintf("subtasks?%v=%v&completed=%v", t, id, completed)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var subtasks []*Subtask

	_, err = s.client.Do(ctx, req, &subtasks)
	if err != nil {
		return nil, err
	}

	return subtasks, nil
}

//
//Get a specific task
//
func (s *SubtaskService) Get(ctx context.Context, id int) (*Subtask, error) {
	u := fmt.Sprintf("subtasks/%v", id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	newSubtask := new(Subtask)
	_, err = s.client.Do(ctx, req, newSubtask)
	if err != nil {
		return nil, err
	}

	return newSubtask, nil
}

//
// Create a task
//
func (s *SubtaskService) Create(ctx context.Context, subtask *Subtask) (*Subtask, error) {
	u := "subtasks"
	req, err := s.client.NewRequest("POST", u, subtask)
	if err != nil {
		return nil, err
	}

	newSubtask := new(Subtask)

	resp, err := s.client.Do(ctx, req, newSubtask)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusCreated {
		return nil, errors.New(fmt.Sprintf("expected status: %v, got: %v", http.StatusCreated, resp.Status))
	}

	return newSubtask, nil
}

//
// Update a task by overwriting properties
//
func (s *SubtaskService) Update(ctx context.Context, subtask *Subtask) (*Subtask, error) {
	u := fmt.Sprintf("subtasks/%d", subtask.Id)

	req, err := s.client.NewRequest("PATCH", u, subtask)
	if err != nil {
		return nil, err
	}

	newSubtask := new(Subtask)
	_, err = s.client.Do(ctx, req, newSubtask)
	if err != nil {
		return nil, err
	}

	return newSubtask, nil
}

//
//Delete a task
//
func (s *SubtaskService) Delete(ctx context.Context, subtask *Subtask) (error) {
	u := fmt.Sprintf("subtasks/%v", subtask.Id)
	req, err := s.client.NewRequest("DELETE", u, subtask)
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
