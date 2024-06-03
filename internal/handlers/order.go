package handlers

import (
    "net/http"
    "encoding/json"
    "github.com/yourusername/foodtruck-cms/internal/models"
    "github.com/yourusername/foodtruck-cms/internal/services"
)

func OrderHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        w.Header().Set("Content-Type", "application/json")
        orders, err := services.GetOrders()
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            json.NewEncoder(w).Encode(map[string]string{"message": "Error retrieving orders"})
            return
        }
        json.NewEncoder(w).Encode(orders)
    case "POST":
        var order models.Order
        json.NewDecoder(r.Body).Decode(&order)
        err := services.CreateOrder(order)
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            json.NewEncoder(w).Encode(map[string]string{"message": "Error creating order"})
            return
        }
        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(order)
    }
}
