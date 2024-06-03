package services

import (
    "errors"
    "github.com/yourusername/foodtruck-cms/internal/models"
)

var orders = []models.Order{
    {ID: 1, MenuItemID: 1, Quantity: 2, Status: "Pending"},
}

func GetOrders() ([]models.Order, error) {
    return orders, nil
}

func CreateOrder(order models.Order) error {
    if order.Quantity <= 0 {
        return errors.New("invalid quantity")
    }
    order.ID = len(orders) + 1
    orders = append(orders, order)
    return nil
}
