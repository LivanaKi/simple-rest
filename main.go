package main

import (
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/Users/natza/simple-rest/auth"
	"github.com/Users/natza/simple-rest/controller"
	"github.com/Users/natza/simple-rest/database"
	"github.com/Users/natza/simple-rest/helper"
	"github.com/Users/natza/simple-rest/repository"
	"github.com/Users/natza/simple-rest/router"
	"github.com/Users/natza/simple-rest/service"
)

func main() {

	log.Printf("Server start")

	db := database.InitDB()
	sellerRepository := repository.NewSeller(db)
	sellerService := service.NewSellerServiceImpl(sellerRepository)
	sellerController := controller.NewSellerController(sellerService)
	routes := router.NewRouter(sellerController)
	securedRoutes := auth.BasicAuthMiddleware(routes)

	server := http.Server{Addr: ":8080", Handler: securedRoutes}
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
