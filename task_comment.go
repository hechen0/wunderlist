package wunderlist

import (
	"fmt"
	"net/http"
	"errors"
	"context"
)

// https://developer.wunderlist.com/documentation/endpoints/task_comment
type TaskCommentService service

type CommentAuthor struct {
	Avatar *string `json:"avatar,omitempty"`
	Id     *string `json:"id,omitempty"`
	Name   *string `json:"name,omitempty"`
}

type TaskComment struct {
	Author             *CommentAuthor `json:"author,omitempty"`
	CreatedAt          *string `json:"created_at,omitempty"`
	CreatedByRequestId *string `json:"created_by_request_id,omitempty"`
	Id                 *string `json:"id,omitempty"`
	LocalCreatedAt     *string `json:"local_created_at,omitempty"`
	Revision           *string `json:"revision,omitempty"`
	TaskId             *string `json:"task_id,omitempty"`
	Text               *string `json:"text,omitempty"`
	Type               *string `json:"type,omitempty"`
}

//
// Get the Comments for a Task or a List
//
func (s *TaskCommentService) All(ctx context.Context, task_or_list interface{}) ([]*TaskComment, error) {

	param, id, err := ParseTaskOrList(task_or_list)

	if err != nil {
		return nil, err
	}

	u := fmt.Sprintf("task_comments?%v=%v", param, id)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var taskComments []*TaskComment

	_, err = s.client.Do(ctx, req, &taskComments)
	if err != nil {
		return nil, err
	}

	return taskComments, nil
}

//
// Get a specific comment
//
func (s *TaskCommentService) Get(ctx context.Context, id int) (*TaskComment, error) {
	u := fmt.Sprintf("task_comments/%v", id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	taskComment := new(TaskComment)
	_, err = s.client.Do(ctx, req, taskComment)
	if err != nil {
		return nil, err
	}

	return taskComment, nil
}

//
// Create a Comment
//
func (s *TaskCommentService) Create(ctx context.Context, taskComment *TaskComment) (*TaskComment, error) {
	u := "task_comments"
	req, err := s.client.NewRequest("POST", u, taskComment)
	if err != nil {
		return nil, err
	}

	newTaskComment := new(TaskComment)

	resp, err := s.client.Do(ctx, req, newTaskComment)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusCreated {
		return nil, errors.New(fmt.Sprintf("expected status: %v, got: %v", http.StatusCreated, resp.Status))
	}

	return newTaskComment, nil
}
