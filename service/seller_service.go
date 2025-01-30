package service

import (
	"context"

	"github.com/Users/natza/simple-rest/data/request"
	"github.com/Users/natza/simple-rest/data/response"
)

type SellerService interface {
	Create(ctx context.Context, request request.SellerCreateRequest)
	Update(ctx context.Context, request request.SellerUpdateRequest)
	Delete(ctx context.Context, sellerId int)
	Read(ctx context.Context) []response.SellerRespons
	FindById(ctx context.Context, sellerId int) response.SellerRespons
}
