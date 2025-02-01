package controller

import (
	"context"

	"github.com/Users/natza/simple-rest/data/request"
	"github.com/Users/natza/simple-rest/data/response"
	"github.com/stretchr/testify/mock"
)

type MockSellerService struct {
	mock.Mock
}

func (m *MockSellerService) Create(ctx context.Context, req request.SellerCreateRequest) error {
	args := m.Called(ctx, req)
	return args.Error(0)
}

func (m *MockSellerService) Update(ctx context.Context, req request.SellerUpdateRequest) error {
	args := m.Called(ctx, req)
	return args.Error(0)
}

func (m *MockSellerService) Delete(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockSellerService) Read(ctx context.Context) ([]response.SellerResponse, error) {
	args := m.Called(ctx)
	return args.Get(0).([]response.SellerResponse), args.Error(1)
}

func (m *MockSellerService) FindByID(ctx context.Context, id int) (response.SellerResponse, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(response.SellerResponse), args.Error(1)
}
