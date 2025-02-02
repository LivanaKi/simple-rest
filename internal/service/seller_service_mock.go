package service

import (
	"context"

	"github.com/stretchr/testify/mock"

	"github.com/Users/natza/simple-rest/internal/model"
)

type MockSellerService struct {
	mock.Mock
}

func (m *MockSellerService) Create(ctx context.Context, req *model.Seller) error {
	args := m.Called(ctx, req)
	return args.Error(0)
}

func (m *MockSellerService) Update(ctx context.Context, req *model.Seller) error {
	args := m.Called(ctx, req)
	return args.Error(0)
}

func (m *MockSellerService) Delete(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockSellerService) Read(ctx context.Context) ([]model.Seller, error) {
	args := m.Called(ctx)
	return args.Get(0).([]model.Seller), args.Error(1)
}

func (m *MockSellerService) FindByID(ctx context.Context, id int) (*model.Seller, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*model.Seller), args.Error(1)
}
