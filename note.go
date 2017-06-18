package wunderlist

import (
	"fmt"
	"net/http"
	"errors"
	"context"
)

// https://developer.wunderlist.com/documentation/endpoints/note
type NoteService service

type Note struct {
	Id                 *int `json:"id,omitempty"`
	Revision           *int `json:"revision,omitempty"`
	Content            *string `json:"Content,omitempty"`
	TaskId             *int `json:"task_id,omitempty"`
	Type               *string `json:"type,omitempty"`
	CreatedByRequestId *string `json:"created_by_request_id,omitempty"`
}

//
// Get the Notes for a Task or List
//
// TODO: why list_id can also work?
func (s *NoteService) All(ctx context.Context, task_id int) ([]*Note, error) {
	u := fmt.Sprintf("notes?task_id=%v", task_id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var notes []*Note
	_, err = s.client.Do(ctx, req, &notes)
	if err != nil {
		return nil, err
	}
	return notes, nil
}

//
// Get a specific note
//
func (s *NoteService) Get(ctx context.Context, id int) (*Note, error) {
	u := fmt.Sprintf("notes/%v", id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	note := new(Note)
	_, err = s.client.Do(ctx, req, note)
	if err != nil {
		return nil, err
	}
	return note, nil
}

//
// Create a note
//
func (s *NoteService) Create(ctx context.Context, note *Note) (*Note, error) {
	u := "notes"
	req, err := s.client.NewRequest("POST", u, note)
	if err != nil {
		return nil, err
	}

	n := new(Note)

	resp, err := s.client.Do(ctx, req, n)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusCreated {
		return nil, errors.New(fmt.Sprintf("expected status: %v, got: %v", http.StatusCreated, resp.Status))
	}

	return n, nil
}

//
// Update a note by overwriting properties
//
func (s *NoteService) Update(ctx context.Context, note *Note) (*Note, error) {
	u := fmt.Sprintf("notes/%v", note.Id)
	req, err := s.client.NewRequest("PATCH", u, note)
	if err != nil {
		return nil, err
	}

	f := new(Note)
	_, err = s.client.Do(ctx, req, f)
	if err != nil {
		return nil, err
	}

	return f, nil
}

//
// Delete a note
//
func (s *NoteService) Delete(ctx context.Context, id int) (error) {
	u := fmt.Sprintf("notes/%v", id)
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
