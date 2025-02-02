package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/Users/natza/simple-rest/internal/data/response"
	"github.com/Users/natza/simple-rest/internal/model"
	"github.com/Users/natza/simple-rest/internal/service"
)

func TestCreateSeller(t *testing.T) {
	mockService := new(service.MockSellerService)
	controller := NewSellerController(mockService)

	sellerReq := &model.Seller{Name: "Test Seller"}
	jsonBody, _ := json.Marshal(sellerReq)

	req := httptest.NewRequest(http.MethodPost, "/sellers", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	mockService.On("Create", mock.Anything, sellerReq).Return(nil)

	controller.Create(w, req, httprouter.Params{})

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusCreated, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestCreateSeller_Error(t *testing.T) {
	mockService := new(service.MockSellerService)
	controller := NewSellerController(mockService)

	sellerReq := &model.Seller{Name: "Test Seller"}
	jsonBody, _ := json.Marshal(sellerReq)

	req := httptest.NewRequest(http.MethodPost, "/sellers", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	mockService.On("Create", mock.Anything, sellerReq).Return(errors.New("DB error"))

	controller.Create(w, req, httprouter.Params{})

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}
func TestFindByID_Success(t *testing.T) {
	mockService := new(service.MockSellerService)
	controller := NewSellerController(mockService)

	expectedSeller := &model.Seller{ID: 1, Name: "Test Seller"}
	mockService.On("FindByID", mock.Anything, 1).Return(expectedSeller, nil)

	req := httptest.NewRequest(http.MethodGet, "/sellers/1", http.NoBody)
	w := httptest.NewRecorder()
	params := httprouter.Params{{Key: "sellerID", Value: "1"}}

	controller.FindByID(w, req, params)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestFindByID_NotFound(t *testing.T) {
	mockService := new(service.MockSellerService)
	controller := NewSellerController(mockService)

	mockService.On("FindByID", mock.Anything, 1).Return(&model.Seller{}, errors.New("not found"))

	req := httptest.NewRequest(http.MethodGet, "/sellers/1", http.NoBody)
	w := httptest.NewRecorder()
	params := httprouter.Params{{Key: "sellerID", Value: "1"}}

	controller.FindByID(w, req, params)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestFindByID_InvalidID(t *testing.T) {
	mockService := new(service.MockSellerService)
	controller := NewSellerController(mockService)

	req := httptest.NewRequest(http.MethodGet, "/sellers/abc", http.NoBody)
	w := httptest.NewRecorder()
	params := httprouter.Params{{Key: "sellerID", Value: "abc"}}

	controller.FindByID(w, req, params)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}
func TestUpdate_Success(t *testing.T) {
	mockService := new(service.MockSellerService)
	controller := NewSellerController(mockService)

	sellerReq := &model.Seller{ID: 1, Name: "Updated Seller"}
	jsonBody, _ := json.Marshal(sellerReq)

	req := httptest.NewRequest(http.MethodPut, "/sellers/1", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	params := httprouter.Params{{Key: "sellerID", Value: "1"}}

	mockService.On("Update", mock.Anything, sellerReq).Return(nil)

	controller.Update(w, req, params)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestDelete_Success(t *testing.T) {
	mockService := new(service.MockSellerService)
	controller := NewSellerController(mockService)

	mockService.On("Delete", mock.Anything, 1).Return(nil)

	req := httptest.NewRequest(http.MethodDelete, "/sellers/1", http.NoBody)
	w := httptest.NewRecorder()
	params := httprouter.Params{{Key: "sellerID", Value: "1"}}

	controller.Delete(w, req, params)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestDelete_Error(t *testing.T) {
	mockService := new(service.MockSellerService)
	controller := NewSellerController(mockService)

	mockService.On("Delete", mock.Anything, 1).Return(errors.New("DB error"))

	req := httptest.NewRequest(http.MethodDelete, "/sellers/1", http.NoBody)
	w := httptest.NewRecorder()
	params := httprouter.Params{{Key: "sellerID", Value: "1"}}

	controller.Delete(w, req, params)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
}
func TestRead_Success(t *testing.T) {
	mockService := new(service.MockSellerService)
	controller := NewSellerController(mockService)

	expectedSellers := []model.Seller{
		{ID: 1, Name: "Seller 1"},
		{ID: 2, Name: "Seller 2"},
	}

	mockService.On("Read", mock.Anything).Return(expectedSellers, nil)

	req := httptest.NewRequest(http.MethodGet, "/sellers", http.NoBody)
	w := httptest.NewRecorder()

	controller.Read(w, req, httprouter.Params{})

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var webResp response.WebResponse
	err := json.NewDecoder(resp.Body).Decode(&webResp)
	if err != nil {
		log.Printf("Error decoding JSON: %v", err)
	}

	assert.Equal(t, http.StatusOK, webResp.Code)
	assert.Equal(t, "Success", webResp.Status)
	assert.Len(t, webResp.Data, 2)

	mockService.AssertExpectations(t)
}

func TestRead_Error(t *testing.T) {
	mockService := new(service.MockSellerService)
	controller := NewSellerController(mockService)

	mockService.On("Read", mock.Anything).Return([]model.Seller{}, errors.New("DB error"))

	req := httptest.NewRequest(http.MethodGet, "/sellers", http.NoBody)
	w := httptest.NewRecorder()

	controller.Read(w, req, httprouter.Params{})

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}
