package handlers

import (
    "net/http"
    "encoding/json"
    "github.com/yourusername/foodtruck-cms/internal/services"
)

func MenuHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    menuItems, err := services.GetMenuItems()
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{"message": "Error retrieving menu items"})
        return
    }
    json.NewEncoder(w).Encode(menuItems)
}
