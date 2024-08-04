package handler

import "net/http"

type Order struct {
}

func (o *Order) CreateOrder(w http.ResponseWriter, r *http.Request) {
}

func (o *Order) ListOrders(w http.ResponseWriter, r *http.Request) {
}

func (o *Order) GetOrderByID(w http.ResponseWriter, r *http.Request) {
}

func (o *Order) UpdateOrder(w http.ResponseWriter, r *http.Request) {
}

func (o *Order) DeleteOrder(w http.ResponseWriter, r *http.Request) {}
