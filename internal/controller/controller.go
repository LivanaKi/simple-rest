package controller

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"github.com/Users/natza/simple-rest/internal/data/request"
	"github.com/Users/natza/simple-rest/internal/data/response"
	"github.com/Users/natza/simple-rest/internal/service"
	"github.com/Users/natza/simple-rest/pkg/helper"
)

type SellerController struct {
	SellerService service.SellerService
}

func NewSellerController(sellerService service.SellerService) *SellerController {
	return &SellerController{SellerService: sellerService}
}

func (controller *SellerController) Create(writer http.ResponseWriter, requests *http.Request, _ httprouter.Params) {
	var sellerCreateRequest request.SellerCreateRequest

	helper.ReadRequestBody(requests, &sellerCreateRequest)

	err := controller.SellerService.Create(requests.Context(), sellerCreateRequest.ToSeller())
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
	var sellerUpdateRequest request.SellerUpdateRequest

	helper.ReadRequestBody(requests, &sellerUpdateRequest)

	sellerID := params.ByName("sellerID")
	id, err := strconv.Atoi(sellerID)
	if err != nil {
		http.Error(writer, "Invalid seller ID", http.StatusBadRequest)
		return
	}
	sellerUpdateRequest.ID = id

	if err = controller.SellerService.Update(requests.Context(), sellerUpdateRequest.ToSeller()); err != nil {
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

	if err = controller.SellerService.Delete(requests.Context(), id); err != nil {
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
	sellers, err := controller.SellerService.Read(requests.Context())
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	sellerResp := make([]response.SellerResponse, 0, len(sellers))

	for _, value := range sellers {
		seller := response.SellerResponse{ID: value.ID, Name: value.Name, Phone: value.Phone}
		sellerResp = append(sellerResp, seller)
	}

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   sellerResp,
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

	resp := response.SellerResponse{
		ID:    result.ID,
		Name:  result.Name,
		Phone: result.Phone,
	}

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "Found",
		Data:   resp,
	}

	helper.WriteResponse(writer, webResponse)
}
