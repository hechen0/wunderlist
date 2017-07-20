package wunderlist

import (
	"context"
	"fmt"
)

//https://developer.wunderlist.com/documentation/endpoints/membership

type MembershipService service

type Membership struct {
	Id       *string `json:"id,omitempty"`
	ListId   *string `json:"list_id,omitemptby"`
	Owner    *bool   `json:"owner,omitemptby"`
	Revision *string `json:"revision,omitempty"`
	State    *string `json:"state,omitempty"`
	Type     *string `json:"type,omitempty"`
	UserId   *string `json:"user_id,omitempty"`
	Email    *string `json:"email,omitempty"` // use to add member
	Muted    *bool   `json:"muted,omitempty"`
}

//Get Memberships for a List or the current User
//
//GET a.wunderlist.com/api/v1/memberships
//
func (s *MembershipService) All(ctx context.Context) ([]*Membership, error) {
	u := "memberships"
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var memberships []*Membership
	_, err = s.client.Do(ctx, req, &memberships)
	if err != nil {
		return nil, err
	}

	return memberships, nil
}

//
//Add a Member to a List
//
//POST a.wunderlist.com/api/v1/memberships
//
func (s *MembershipService) AddMember(ctx context.Context, membership *Membership) (*Membership, error) {
	u := "memberships"
	req, err := s.client.NewRequest("POST", u, membership)
	if err != nil {
		return nil, err
	}

	m := new(Membership)
	_, err = s.client.Do(ctx, req, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

//Mark a Member as accepted
//
//PATCH a.wunderlist.com/api/v1/memberships/:id
//
func (s *MembershipService) Accept(ctx context.Context, membership *Membership) (*Membership, error) {
	u := fmt.Sprintf("memberships/%v", membership.Id)
	req, err := s.client.NewRequest("PATCH", u, membership)
	if err != nil {
		return nil, err
	}

	m := new(Membership)
	_, err = s.client.Do(ctx, req, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

//Remove a Member from a List
//
//DELETE a.wunderlist.com/api/v1/memberships/:id
//
//Reject an invite to a List
//
//DELETE a.wunderlist.com/api/v1/memberships/:id
//
func (s *MembershipService) Reject(ctx context.Context, id int) error {
	u := fmt.Sprintf("memberships/%v", id)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(ctx, req, nil)
	if err != nil {
		return err
	}

	return nil
}
