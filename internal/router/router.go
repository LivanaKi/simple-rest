package router

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/Users/natza/simple-rest/internal/controller"
)

func NewRouter(sellerController *controller.SellerController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, "Hello")
	})

	router.GET("/api/seller", sellerController.Read)
	router.POST("/api/seller", sellerController.Create)
	router.PATCH("/api/seller/:sellerID", sellerController.Update)
	router.DELETE("/api/seller/:sellerID", sellerController.Delete)
	router.GET("/api/seller/:sellerID", sellerController.FindByID)

	return router
}
