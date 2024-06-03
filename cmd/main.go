package main

import (
    "log"
    "net/http"
    "github.com/yourusername/foodtruck-cms/internal/handlers"
    "github.com/yourusername/foodtruck-cms/internal/middleware"
)

func main() {
    http.Handle("/api/v1/menus", middleware.JWTMiddleware(http.HandlerFunc(handlers.MenuHandler)))
    http.Handle("/api/v1/orders", middleware.JWTMiddleware(http.HandlerFunc(handlers.OrderHandler)))
    http.HandleFunc("/api/v1/auth", handlers.AuthHandler)

    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("Could not start server: %s\n", err.Error())
    }
}
