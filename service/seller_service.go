package service

import (
	"context"

	"github.com/Users/natza/simple-rest/data/request"
	"github.com/Users/natza/simple-rest/data/response"
)

type SellerService interface {
	Create(ctx context.Context, request request.SellerCreateRequest) error
	Update(ctx context.Context, request request.SellerUpdateRequest) error
	Delete(ctx context.Context, sellerID int) error
	Read(ctx context.Context) ([]response.SellerResponse, error)
	FindByID(ctx context.Context, sellerID int) (response.SellerResponse, error)
}
