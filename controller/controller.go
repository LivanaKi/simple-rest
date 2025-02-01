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

type SellerController struct {
	SellerService service.SellerService
}

func NewSellerController(sellerService service.SellerService) *SellerController {
	return &SellerController{SellerService: sellerService}
}

func (controller *SellerController) Create(writer http.ResponseWriter, requests *http.Request, _ httprouter.Params) {
	sellerCreateRequest := request.SellerCreateRequest{}
	helper.ReadRequestBody(requests, &sellerCreateRequest)

	err := controller.SellerService.Create(requests.Context(), sellerCreateRequest)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	webResponse := response.WebResponse{
		Code:   http.StatusCreated,
		Status: "Created",
		Data:   nil,
	}
	helper.WriteResponse(writer, webResponse)
}

func (controller *SellerController) Update(
	writer http.ResponseWriter,
	requests *http.Request,
	params httprouter.Params,
) {
	sellerUpdateRequest := request.SellerUpdateRequest{}
	helper.ReadRequestBody(requests, &sellerUpdateRequest)

	sellerID := params.ByName("sellerID")
	id, err := strconv.Atoi(sellerID)
	if err != nil {
		http.Error(writer, "Invalid seller ID", http.StatusBadRequest)
		return
	}
	sellerUpdateRequest.ID = id

	if err := controller.SellerService.Update(requests.Context(), sellerUpdateRequest); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "Updated",
		Data:   nil,
	}
	helper.WriteResponse(writer, webResponse)
}

func (controller *SellerController) Delete(
	writer http.ResponseWriter,
	requests *http.Request,
	params httprouter.Params,
) {
	sellerID := params.ByName("sellerID")
	id, err := strconv.Atoi(sellerID)
	if err != nil {
		http.Error(writer, "Invalid seller ID", http.StatusBadRequest)
		return
	}

	if err := controller.SellerService.Delete(requests.Context(), id); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "Deleted",
		Data:   nil,
	}
	helper.WriteResponse(writer, webResponse)
}

func (controller *SellerController) Read(writer http.ResponseWriter, requests *http.Request, _ httprouter.Params) {
	result, err := controller.SellerService.Read(requests.Context())
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	if result == nil {
		result = nil
	}

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   result,
	}
	helper.WriteResponse(writer, webResponse)
}

func (controller *SellerController) FindByID(
	writer http.ResponseWriter,
	requests *http.Request,
	params httprouter.Params,
) {
	sellerID := params.ByName("sellerID")
	id, err := strconv.Atoi(sellerID)
	if err != nil {
		http.Error(writer, "Invalid seller ID", http.StatusBadRequest)
		return
	}

	result, err := controller.SellerService.FindByID(requests.Context(), id)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusNotFound)
		return
	}

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "Found",
		Data:   result,
	}
	helper.WriteResponse(writer, webResponse)
}
