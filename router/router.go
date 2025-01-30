package router

import (
	"fmt"
	"net/http"

	"github.com/Users/natza/simple-rest/controller"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(sellerController *controller.SellerControler) *httprouter.Router {
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Welcome Home ")
	})

	router.GET("/api/seller", sellerController.Read)
	router.POST("/api/seller", sellerController.Create)
	router.PATCH("/api/seller/:sellerId", sellerController.Update)
	router.DELETE("/api/seller/:sellerId", sellerController.Delete)
	router.GET("/api/seller/:sellerId", sellerController.FindById)

	return router
}
