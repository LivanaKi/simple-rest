package service

import (
	"context"

	"github.com/Users/natza/simple-rest/internal/model"
	"github.com/Users/natza/simple-rest/internal/repository"
)

type sellerService struct {
	sellerRepository repository.SellerRepository
}

func NewSellerServiceImpl(sellerRepository repository.SellerRepository) SellerService {
	return &sellerService{sellerRepository: sellerRepository}
}

// Delete
func (b *sellerService) Delete(ctx context.Context, sellerID int) error {
	seller, err := b.sellerRepository.FindByID(ctx, sellerID)
	if err != nil {
		return err
	}

	err = b.sellerRepository.Delete(ctx, seller.ID)
	if err != nil {
		return err
	}

	return nil
}

// Create
func (b *sellerService) Create(ctx context.Context, seller *model.Seller) error {
	err := seller.Validation()
	if err != nil {
		return err
	}

	err = b.sellerRepository.Save(ctx, seller)
	if err != nil {
		return err
	}

	return nil
}

// Read
func (b *sellerService) Read(ctx context.Context) ([]model.Seller, error) {
	return b.sellerRepository.Read(ctx)
}

// Update
func (b *sellerService) Update(ctx context.Context, newSeller *model.Seller) error {
	err := newSeller.Validation()
	if err != nil {
		return err
	}

	seller, err := b.sellerRepository.FindByID(ctx, newSeller.ID)
	if err != nil {
		return err
	}

	seller.Name = newSeller.Name
	seller.Phone = newSeller.Phone

	err = b.sellerRepository.Update(ctx, seller)
	if err != nil {
		return err
	}

	return nil
}

// Find by id
func (b *sellerService) FindByID(ctx context.Context, sellerID int) (*model.Seller, error) {
	return b.sellerRepository.FindByID(ctx, sellerID)
}
