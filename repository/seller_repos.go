package repository

import (
	"context"

	"github.com/Users/natza/simple-rest/model"
)

type SellerRepository interface {
	Create(ctx context.Context, seller model.Seller)
	Read(ctx context.Context) []model.Seller
	Update(ctx context.Context, seller model.Seller)
	Delete(ctx context.Context, sellerId int)
	FindById(ctx context.Context, sellerId int) (model.Seller, error)
}
