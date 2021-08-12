package main

import (
	"log"
	"os"
	"product-api/inmem"
	"product-api/rest"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func createRouter() chi.Router {
	r := chi.NewRouter()

	l := log.New(os.Stdout, "products-api", log.LstdFlags)
	ps := inmem.NewProductStore()
	ph := rest.NewProductHandler(ps, l)

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/products", ph.Routes)
	return r
}
