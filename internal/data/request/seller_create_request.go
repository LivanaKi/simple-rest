package request

import "github.com/Users/natza/simple-rest/internal/model"

type SellerCreateRequest struct {
	Name  string `validate:"required min=1, max=100" json:"name"`
	Phone string `validate:"required min=1, max=20" json:"phone"`
}

func (s *SellerCreateRequest) ToSeller() *model.Seller {
	return &model.Seller{
		Name:  s.Name,
		Phone: s.Phone,
	}
}
