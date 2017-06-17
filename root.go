package wunderlist

import "context"

// api: https://developer.wunderlist.com/documentation/endpoints/user
// Root is the top-level entity in the sync hierarchy.

type RootService service

type Root struct {
	Id       *int    `json:"id,omitempty"`
	Type     *string `json:"type,omitempty"`
	Revision *int `json:"revision,omitempty"`
	UserId   *int `json:"user_id,omitempty"`
}

//Fetch the Root for the current User
func (s *RootService) Current(ctx context.Context) (*Root, error) {
	req, err := s.client.NewRequest("GET", "root", nil)
	if err != nil {
		return nil, err
	}

	newRoot := new(Root)
	_, err = s.client.Do(ctx, req, newRoot)
	if err != nil {
		return nil, err
	}
	return newRoot, nil
}
