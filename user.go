package wunderlist

import (
	"context"
)

// api: https://developer.wunderlist.com/documentation/endpoints/user

type UserService service

type User struct {
	Id        *int    `json:"id,omitempty"`
	Email     *string `json:"email,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	Name      *string `json:"name,omitempty"`
	Type      *string `json:"type,omitempty"`
	Revision  *int `json:"revision,omitempty"`
}

//Fetch the currently logged in user
func (s *UserService) Current(ctx context.Context) (*User, error) {
	req, err := s.client.NewRequest("GET", "user", nil)
	if err != nil {
		return nil, err
	}

	newUser := new(User)
	_, err = s.client.Do(ctx, req, newUser)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

//Fetch the users this user can access
func (s *UserService) All(ctx context.Context) ([]*User, error) {
	req, err := s.client.NewRequest("GET", "users", nil)
	if err != nil {
		return nil, err
	}

	var users []*User
	_, err = s.client.Do(ctx, req, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}
