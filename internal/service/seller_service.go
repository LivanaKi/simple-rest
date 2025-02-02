package service

import (
	"context"

	"github.com/Users/natza/simple-rest/internal/model"
)

type SellerService interface {
	Create(ctx context.Context, request *model.Seller) error
	Update(ctx context.Context, request *model.Seller) error
	Delete(ctx context.Context, sellerID int) error
	Read(ctx context.Context) ([]model.Seller, error)
	FindByID(ctx context.Context, sellerID int) (*model.Seller, error)
}
