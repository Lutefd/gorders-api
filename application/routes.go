package application

import (
	"github.com/Lutefd/gorders-api/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func loadRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Route("/orders", loadOrderRoutes)
	return router
}

func loadOrderRoutes(router chi.Router) {
	orderHandler := &handler.Order{}
	router.Post("/", orderHandler.CreateOrder)
	router.Get("/", orderHandler.ListOrders)
	router.Get("/{id}", orderHandler.GetOrderByID)
	router.Put("/{id}", orderHandler.UpdateOrder)
	router.Delete("/{id}", orderHandler.DeleteOrder)
}
