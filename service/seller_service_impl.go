package service

import (
	"context"

	"github.com/Users/natza/simple-rest/data/request"
	"github.com/Users/natza/simple-rest/data/response"
	"github.com/Users/natza/simple-rest/helper"
	"github.com/Users/natza/simple-rest/model"
	"github.com/Users/natza/simple-rest/repository"
)

type SellerServiceImpl struct {
	SellerRepository repository.SellerRepository
}

func NewSellerServiceImpl(sellerRepository repository.SellerRepository) SellerService {
	return &SellerServiceImpl{SellerRepository: sellerRepository}
}

// Delete
func (b *SellerServiceImpl) Delete(ctx context.Context, sellerId int) {
	seller, err := b.SellerRepository.FindById(ctx, sellerId)
	helper.PanicIfError(err)
	b.SellerRepository.Delete(ctx, seller.Id)
}

// Create
func (b *SellerServiceImpl) Create(ctx context.Context, request request.SellerCreateRequest) {
	seller := model.Seller{
		Name:  request.Name,
		Phone: request.Phone,
	}
	b.SellerRepository.Create(ctx, seller)
}

// Read
func (b *SellerServiceImpl) Read(ctx context.Context) []response.SellerRespons {
	sellers := b.SellerRepository.Read(ctx)

	var sellerResp []response.SellerRespons

	for _, value := range sellers {
		seller := response.SellerRespons{Id: value.Id, Name: value.Name, Phone: value.Phone}
		sellerResp = append(sellerResp, seller)
	}
	return sellerResp
}

// Update
func (b *SellerServiceImpl) Update(ctx context.Context, request request.SellerUpdateRequest) {
	seller, err := b.SellerRepository.FindById(ctx, request.Id)
	helper.PanicIfError(err)

	seller.Name = request.Name
	seller.Phone = request.Phone
	b.SellerRepository.Update(ctx, seller)
}

// Find by id
func (b *SellerServiceImpl) FindById(ctx context.Context, sellerId int) response.SellerRespons {
	seller, err := b.SellerRepository.FindById(ctx, sellerId)
	helper.PanicIfError(err)
	return response.SellerRespons(seller)
}
