package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/Users/natza/simple-rest/internal/model"
	"github.com/Users/natza/simple-rest/internal/repository"
)

func TestCreate(t *testing.T) {
	mockRepo := new(repository.MockSellerRepository)
	service := NewSellerServiceImpl(mockRepo)

	ctx := context.Background()
	req := &model.Seller{Name: "test", Phone: "+11111"}

	mockRepo.On("Save", ctx, mock.AnythingOfType("*model.Seller")).Return(nil)

	err := service.Create(ctx, req)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDelete(t *testing.T) {
	mockRepo := new(repository.MockSellerRepository)
	service := NewSellerServiceImpl(mockRepo)

	ctx := context.Background()
	seller := &model.Seller{ID: 1, Name: "test", Phone: "+11111"}

	mockRepo.On("FindByID", ctx, 1).Return(seller, nil)
	mockRepo.On("Delete", ctx, 1).Return(nil)

	err := service.Delete(ctx, 1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestFindByID(t *testing.T) {
	mockRepo := new(repository.MockSellerRepository)
	service := NewSellerServiceImpl(mockRepo)

	ctx := context.Background()
	seller := &model.Seller{ID: 1, Name: "test", Phone: "+11111"}

	mockRepo.On("FindByID", ctx, 1).Return(seller, nil)

	result, err := service.FindByID(ctx, 1)

	assert.NoError(t, err)
	assert.Equal(t, seller.ID, result.ID)
	assert.Equal(t, seller.Name, result.Name)
	assert.Equal(t, seller.Phone, result.Phone)
	mockRepo.AssertExpectations(t)
}

func TestUpdate(t *testing.T) {
	mockRepo := new(repository.MockSellerRepository)
	service := NewSellerServiceImpl(mockRepo)

	ctx := context.Background()
	existingSeller := &model.Seller{ID: 1, Name: "test", Phone: "+11111"}
	updateReq := &model.Seller{ID: 1, Name: "test", Phone: "+22222"}

	mockRepo.On("FindByID", ctx, 1).Return(existingSeller, nil)
	mockRepo.On("Update", ctx, mock.AnythingOfType("*model.Seller")).Return(nil)

	err := service.Update(ctx, updateReq)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestRead(t *testing.T) {
	mockRepo := new(repository.MockSellerRepository)
	service := NewSellerServiceImpl(mockRepo)

	ctx := context.Background()
	sellers := []model.Seller{
		{ID: 1, Name: "test", Phone: "+11111"},
		{ID: 2, Name: "test1", Phone: "+22222"},
	}

	mockRepo.On("Read", ctx).Return(sellers, nil)

	result, err := service.Read(ctx)

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, result[0].ID, sellers[0].ID)
	assert.Equal(t, result[1].ID, sellers[1].ID)
	mockRepo.AssertExpectations(t)
}
