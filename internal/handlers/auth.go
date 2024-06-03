package handlers

import (
    "net/http"
    "encoding/json"
    "github.com/yourusername/foodtruck-cms/internal/models"
    "github.com/yourusername/foodtruck-cms/internal/services"
)

func AuthHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var user models.User
    json.NewDecoder(r.Body).Decode(&user)
    token, err := services.AuthenticateUser(user)
    if err != nil {
        w.WriteHeader(http.StatusUnauthorized)
        json.NewEncoder(w).Encode(map[string]string{"message": "Invalid credentials"})
        return
    }
    json.NewEncoder(w).Encode(map[string]string{"token": token})
}
