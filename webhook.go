package wunderlist

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

// api: https://developer.wunderlist.com/documentation/endpoints/webhooks

type WebhookService service

type Webhook struct {
	Id            *int    `json:"id,omitempty"`
	ListId        *int    `json:"list_id,omitempty"`
	Url           *string `json:"url,omitempty"`
	ProcessorType *string `json:"processor_type,omitempty"`
	Configuration *string `json:"configuration,omitempty"`
}

//Get all webhooks for a list
//
//GET a.wunderlist.com/api/v1/webhooks
func (s *WebhookService) Get(ctx context.Context, listId int) ([]*Webhook, error) {
	u := fmt.Sprintf("webhooks?list_id=%v", listId)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var webhooks []*Webhook

	_, err = s.client.Do(ctx, req, &webhooks)
	if err != nil {
		return nil, err
	}

	return webhooks, nil
}

//Create a Webhook
//
//POST a.wunderlist.com/api/v1/webhooks
//
func (s *WebhookService) Create(ctx context.Context, webhook *Webhook) (*Webhook, error) {
	u := "webhooks"
	req, err := s.client.NewRequest("POST", u, webhook)
	if err != nil {
		return nil, err
	}

	newWebhook := new(Webhook)

	resp, err := s.client.Do(ctx, req, newWebhook)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusCreated {
		return nil, errors.New(fmt.Sprintf("expected status: %v, got: %v", http.StatusCreated, resp.Status))
	}

	return newWebhook, nil
}

//Delete a webhook permanently
//
//DELETE a.wunderlist.com/api/v1/webhooks/:id
//
func (s *WebhookService) Delete(ctx context.Context, id, revision int) error {
	u := fmt.Sprintf("webhooks/%v?revision=%v", id, revision)
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
