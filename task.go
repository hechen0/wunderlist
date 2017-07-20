package wunderlist

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

// https://developer.wunderlist.com/documentation/endpoints/list
type TaskService service

type Task struct {
	Id                 *int    `json:"id,omitempty"`
	Title              *string `json:"title,omitempty"`
	Type               *string `json:"type,omitempty"`
	Completed          *bool   `json:"completed,omitempty"`
	CreatedAt          *string `json:"created_at,omitempty"`
	CreatedById        *int    `json:"created_by_id,omitempty"`
	CreatedByRequestId *string `json:"created_by_request_id,omitempty"`
	ListId             *int    `json:"list_id,omitempty"`
	Revision           *int    `json:"revision,omitempty"`
	Starred            *bool   `json:"starred,omitempty"`
	DueDate            *string `json:"due_date"`
	RecurrenceCount    *string `json:"recurrence_count,omitempty"`
	RecurrenceType     *string `json:"recurrence_type,omitempty"`
	CompletedAt        *string `json:"completed_at,omitempty"`
	CompletedById      *string `json:"completed_by_id,omitempty"`
	AssigneeId         *string `json:"assignee_id,omitempty"`

	//an array of attributes to delete from the task, e.g. 'due_date'
	Remove []string `json:"remove,omitempty"`
}

//
//Get Tasks for a List
//
func (s *TaskService) All(ctx context.Context, list_id int) ([]*Task, error) {
	u := fmt.Sprintf("tasks?list_id=%v", list_id)
	req, err := s.client.NewRequest("GET", u, nil)
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

//
//Get Completed Tasks
//
func (s *TaskService) Completed(ctx context.Context, list_id int, completed bool) ([]*Task, error) {
	u := fmt.Sprintf("tasks?list_id=%v&completed=%v", list_id, completed)
	req, err := s.client.NewRequest("GET", u, nil)
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

//
//Get a specific task
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
// Create a task
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
// Update a task by overwriting properties
//
func (s *TaskService) Update(ctx context.Context, task *Task) (*Task, error) {
	u := fmt.Sprintf("tasks/%d", task.Id)

	req, err := s.client.NewRequest("PATCH", u, task)
	if err != nil {
		return nil, err
	}

	newTask := new(Task)
	_, err = s.client.Do(ctx, req, newTask)
	if err != nil {
		return nil, err
	}

	return newTask, nil
}

//
//Delete a task
//
func (s *TaskService) Delete(ctx context.Context, id, revision int) error {
	u := fmt.Sprintf("tasks/%v?revision=%v", id, revision)
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
