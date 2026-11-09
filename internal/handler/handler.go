package handler

import "example.com/orders/internal/model"

type Handler struct {
	orders Service
}

type Service interface {
	Create(customer string) (model.Order, error)
}

func NewHandler(orders Service) *Handler {
	return &Handler{orders: orders}
}

func (h *Handler) Create(customer string) (model.Order, error) {
	return h.orders.Create(customer)
}
