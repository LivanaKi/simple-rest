package request

type SellerUpdateRequest struct {
	ID    int
	Name  string `validate:"required min=1, max=100" json:"name"`
	Phone string `validate:"required min=1, max=20" json:"phone"`
}
