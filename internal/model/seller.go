package model

import "errors"

type Seller struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func (s *Seller) Validation() error {
	if s.Name == "" {
		return errors.New("name is empty")
	}

	return nil
}
