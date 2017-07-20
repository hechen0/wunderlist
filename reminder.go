package wunderlist

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

// https://developer.wunderlist.com/documentation/endpoints/reminder
type ReminderService service

type Reminder struct {
	Id                  *int    `json:"id,omitempty"`
	Type                *string `json:"type,omitempty"`
	CreatedAt           *string `json:"created_at,omitempty"`
	CreatedByRequestId  *string `json:"created_by_request_id,omitempty"`
	Revision            *int    `json:"revision,omitempty"`
	TaskId              *int    `json:"task_id,omitempty"`
	Date                *string `json:"date,omitempty"`
	UpdatedAt           *string `json:"updated_at,omitempty"`
	CreatedByDeviceUdid *string `json:"created_by_device_udid,omitempty"`
}

//
// Get Reminders for a Task or List
//
func (s *ReminderService) All(ctx context.Context, task_or_list interface{}) ([]*Reminder, error) {
	t, id, err := ParseTaskOrList(task_or_list)

	if err != nil {
		return nil, err
	}

	u := fmt.Sprintf("reminders?%v=%v", t, id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var reminders []*Reminder

	_, err = s.client.Do(ctx, req, &reminders)
	if err != nil {
		return nil, err
	}

	return reminders, nil
}

//
// Create a Reminder
//
func (s *ReminderService) Create(ctx context.Context, reminder *Reminder) (*Reminder, error) {
	req, err := s.client.NewRequest("POST", "reminders", reminder)
	if err != nil {
		return nil, err
	}

	newReminder := new(Reminder)

	resp, err := s.client.Do(ctx, req, newReminder)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusCreated {
		return nil, errors.New(fmt.Sprintf("expected status: %v, got: %v", http.StatusCreated, resp.Status))
	}

	return newReminder, nil
}

//
// Update a Reminder
//
func (s *ReminderService) Update(ctx context.Context, reminder *Reminder) (*Reminder, error) {
	u := fmt.Sprintf("reminders/%d", reminder.Id)

	req, err := s.client.NewRequest("PATCH", u, reminder)
	if err != nil {
		return nil, err
	}

	newReminder := new(Reminder)
	_, err = s.client.Do(ctx, req, newReminder)
	if err != nil {
		return nil, err
	}

	return newReminder, nil
}

//
// Delete a Reminder
//
func (s *ReminderService) Delete(ctx context.Context, reminder *Reminder) error {
	u := fmt.Sprintf("reminders/%v", reminder.Id)
	req, err := s.client.NewRequest("DELETE", u, reminder)
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
