package main

import (
	"fmt"

	"example.com/orders/internal/config"
	server "example.com/orders/internal/handler"
	repo "example.com/orders/internal/repository"
	"example.com/orders/internal/service"
)

func main() {
	repo := repo.NewMemory()
	orders := service.NewService(repo)
	handler := server.NewHandler(orders)

	address := config.GetAddress()

	order, err := handler.Create("Покупатель")
	if err != nil {
		fmt.Println("ошибка:", err)
		return
	}

	fmt.Printf("сервис слушал бы %s, создан заказ: %+v\n", address, order)
}
