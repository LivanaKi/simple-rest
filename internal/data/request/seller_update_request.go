package request

import "github.com/Users/natza/simple-rest/internal/model"

type SellerUpdateRequest struct {
	ID    int
	Name  string `validate:"required min=1, max=100" json:"name"`
	Phone string `validate:"required min=1, max=20" json:"phone"`
}

func (s *SellerUpdateRequest) ToSeller() *model.Seller {
	return &model.Seller{
		ID:    s.ID,
		Name:  s.Name,
		Phone: s.Phone,
	}
}
