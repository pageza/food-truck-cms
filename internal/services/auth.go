package services

import (
    "errors"
    "github.com/yourusername/foodtruck-cms/internal/models"
    "github.com/yourusername/foodtruck-cms/internal/middleware"
)

var users = []models.User{
    {Username: "admin", Password: "adminpass", Role: "admin"},
    {Username: "vendor", Password: "vendorpass", Role: "vendor"},
}

func AuthenticateUser(user models.User) (string, error) {
    for _, u := range users {
        if u.Username == user.Username && u.Password == u.Password {
            token, err := middleware.GenerateJWT(user.Username)
            if err != nil {
                return "", err
            }
            return token, nil
        }
    }
    return "", errors.New("invalid credentials")
}
