package routers

import (
    "user-management/routers/api"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"

    "github.com/gin-contrib/cors"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
    router := gin.Default()

    // Configure CORS middleware
    router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:5173"}, // Frontend URL
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Authorization", "Content-Type", "Accept"},
        AllowCredentials: true,
    }))

    apiGroup := router.Group("/api/v1")
    {
        api.SetupUserRouter(apiGroup, db)
        api.SetupAuthRouter(apiGroup, db)
        api.SetupProjectRouter(apiGroup, db)
    }

    return router
}
