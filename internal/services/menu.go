package services

import (
    "github.com/yourusername/foodtruck-cms/internal/models"
)

var menuItems = []models.MenuItem{
    {ID: 1, Name: "Burger", Price: 5.99, Category: "Main"},
    {ID: 2, Name: "Fries", Price: 2.99, Category: "Side"},
}

func GetMenuItems() ([]models.MenuItem, error) {
    return menuItems, nil
}
