package service

import (
	"context"

	"github.com/Users/natza/simple-rest/model"
	"github.com/stretchr/testify/mock"
)

type MockSellerRepository struct {
	mock.Mock
}

func (m *MockSellerRepository) Save(ctx context.Context, seller model.Seller) error {
	args := m.Called(ctx, seller)
	return args.Error(0)
}

func (m *MockSellerRepository) FindByID(ctx context.Context, id int) (model.Seller, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(model.Seller), args.Error(1)
}

func (m *MockSellerRepository) Delete(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockSellerRepository) Read(ctx context.Context) ([]model.Seller, error) {
	args := m.Called(ctx)
	return args.Get(0).([]model.Seller), args.Error(1)
}

func (m *MockSellerRepository) Update(ctx context.Context, seller model.Seller) error {
	args := m.Called(ctx, seller)
	return args.Error(0)
}
