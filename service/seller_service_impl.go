package service

import (
	"context"

	"github.com/Users/natza/simple-rest/data/request"
	"github.com/Users/natza/simple-rest/data/response"
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
func (b *SellerServiceImpl) Delete(ctx context.Context, sellerID int) error {
	seller, err := b.SellerRepository.FindByID(ctx, sellerID)
	if err != nil {
		return err
	}
	err = b.SellerRepository.Delete(ctx, seller.ID)
	if err != nil {
		return err
	}

	return nil
}

// Create
func (b *SellerServiceImpl) Create(ctx context.Context, request request.SellerCreateRequest) error {
	seller := model.Seller{
		Name:  request.Name,
		Phone: request.Phone,
	}
	err := b.SellerRepository.Save(ctx, seller)
	if err != nil {
		return err
	}
	return nil
}

// Read
func (b *SellerServiceImpl) Read(ctx context.Context) ([]response.SellerResponse, error) {
	sellers, err := b.SellerRepository.Read(ctx)
	if err != nil {
		return nil, err
	}

	sellerResp := make([]response.SellerResponse, 0, len(sellers))

	for _, value := range sellers {
		seller := response.SellerResponse{ID: value.ID, Name: value.Name, Phone: value.Phone}
		sellerResp = append(sellerResp, seller)
	}
	return sellerResp, nil
}

// Update
func (b *SellerServiceImpl) Update(ctx context.Context, request request.SellerUpdateRequest) error {
	seller, err := b.SellerRepository.FindByID(ctx, request.ID)
	if err != nil {
		return err
	}

	seller.Name = request.Name
	seller.Phone = request.Phone
	err = b.SellerRepository.Update(ctx, seller)
	if err != nil {
		return err
	}

	return nil
}

// Find by id
func (b *SellerServiceImpl) FindByID(ctx context.Context, sellerID int) (response.SellerResponse, error) {
	seller, err := b.SellerRepository.FindByID(ctx, sellerID)
	if err != nil {
		return response.SellerResponse{}, err
	}
	return response.SellerResponse{ID: seller.ID, Name: seller.Name, Phone: seller.Phone}, nil
}
