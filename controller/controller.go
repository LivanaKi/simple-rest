package controller

import (
	"net/http"
	"strconv"

	"github.com/Users/natza/simple-rest/data/request"
	"github.com/Users/natza/simple-rest/data/response"
	"github.com/Users/natza/simple-rest/helper"
	"github.com/Users/natza/simple-rest/service"
	"github.com/julienschmidt/httprouter"
)

type SellerControler struct {
	SellerService service.SellerService
}

func NewSellerController(sellerService service.SellerService) *SellerControler {
	return &SellerControler{SellerService: sellerService}
}

func (controller *SellerControler) Create(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	sellerCreateRequest := request.SellerCreateRequest{}
	helper.ReadRequestBody(requests, &sellerCreateRequest)

	controller.SellerService.Create(requests.Context(), sellerCreateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}
	helper.WriteResponse(writer, webResponse)
}

func (controller *SellerControler) Update(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	sellerUpdateRequest := request.SellerUpdateRequest{}
	helper.ReadRequestBody(requests, &sellerUpdateRequest)

	sellerId := params.ByName("sellerId")
	id, err := strconv.Atoi(sellerId)
	helper.PanicIfError(err)
	sellerUpdateRequest.Id = id

	controller.SellerService.Update(requests.Context(), sellerUpdateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	helper.WriteResponse(writer, webResponse)
}

func (controller *SellerControler) Delete(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	sellerId := params.ByName("sellerId")
	id, err := strconv.Atoi(sellerId)
	helper.PanicIfError(err)

	controller.SellerService.Delete(requests.Context(), id)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	helper.WriteResponse(writer, webResponse)
}

func (controller *SellerControler) Read(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	result := controller.SellerService.Read(requests.Context())
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	helper.WriteResponse(writer, webResponse)
}

func (controller *SellerControler) FindById(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	sellerId := params.ByName("sellerId")
	id, err := strconv.Atoi(sellerId)
	helper.PanicIfError(err)
	result := controller.SellerService.FindById(requests.Context(), id)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	helper.WriteResponse(writer, webResponse)
}
