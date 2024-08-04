package application

import (
	"github.com/Lutefd/gorders-api/handler"
	"github.com/Lutefd/gorders-api/repository/order"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (a *App) loadRoutes() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Route("/orders", a.loadOrderRoutes)
	a.router = router
}

func (a *App) loadOrderRoutes(router chi.Router) {
	orderHandler := &handler.Order{
		Repo: &order.RedisRepo{
			Client: a.rdb,
		},
	}
	router.Post("/", orderHandler.CreateOrder)
	router.Get("/", orderHandler.ListOrders)
	router.Get("/{id}", orderHandler.GetOrderByID)
	router.Put("/{id}", orderHandler.UpdateOrder)
	router.Delete("/{id}", orderHandler.DeleteOrder)
}
