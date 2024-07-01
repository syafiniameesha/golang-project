package main

import (
    "fmt"
    "log"
    "user-management/database"
    "user-management/models"
    "user-management/routers"

)

func main() {
    // Initialize database connection
    db, err := database.InitDB()
    if err != nil {
        log.Fatalf("Error connecting to database: %v", err)
    }
    defer func() {
        if err := database.CloseDB(); err != nil {
            log.Fatalf("Error closing database: %v", err)
        }
    }()

    // Auto-migrate models
    // add more models
    if err := database.AutoMigrate(db, &models.User{},  &models.Project{}, &models.Status{}, &models.Type{}); err != nil {
        log.Fatalf("Error migrating database: %v", err)
    }

    // Initialize Gin router
    router := routers.SetupRouter(db)

    // Start server
    port := ":8080"
    fmt.Printf("Server running on %s\n", port)
    if err := router.Run(port); err != nil {
        log.Fatalf("Error starting server: %v", err)
    }
}
