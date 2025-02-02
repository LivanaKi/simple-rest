package repository

import (
	"context"

	"github.com/Users/natza/simple-rest/internal/model"
)

type SellerRepository interface {
	Save(ctx context.Context, seller *model.Seller) error
	Read(ctx context.Context) ([]model.Seller, error)
	Update(ctx context.Context, seller *model.Seller) error
	Delete(ctx context.Context, sellerID int) error
	FindByID(ctx context.Context, sellerID int) (*model.Seller, error)
}
