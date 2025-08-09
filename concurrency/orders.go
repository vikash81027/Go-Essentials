package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Order struct {
	ID     int
	Status string
}

func ProcessOrders(orders []*Order) {
	for _, order := range orders {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		fmt.Printf("Processing order %d\n", order.ID)
	}
}

func GenerateOrders(count int) []*Order {
	orders := make([]*Order, count)

	for idx := range orders {
		orders[idx] = &Order{
			ID:     idx + 1,
			Status: "Pending",
		}
	}

	return orders
}