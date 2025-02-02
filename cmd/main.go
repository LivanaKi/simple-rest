package main

import (
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"

	"github.com/Users/natza/simple-rest/internal/controller"
	"github.com/Users/natza/simple-rest/internal/repository"
	"github.com/Users/natza/simple-rest/internal/router"
	"github.com/Users/natza/simple-rest/internal/service"
	"github.com/Users/natza/simple-rest/pkg/auth"
	"github.com/Users/natza/simple-rest/pkg/helper"
	"github.com/Users/natza/simple-rest/pkg/pg"
)

func main() {
	log.Printf("Server start")

	db := pg.InitDB()
	sellerRepository := repository.NewSeller(db)
	sellerService := service.NewSellerServiceImpl(sellerRepository)
	sellerController := controller.NewSellerController(sellerService)
	routes := router.NewRouter(sellerController)
	securedRoutes := auth.BasicAuthMiddleware(routes)

	server := &http.Server{
		Addr:              ":8080",
		Handler:           securedRoutes,
		ReadHeaderTimeout: 5 * time.Second,
	}
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
