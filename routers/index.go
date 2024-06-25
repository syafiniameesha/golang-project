package routers

import (
    "user-management/routers/api"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
    router := gin.Default()

    apiGroup := router.Group("/api/v1")
    {
        api.SetupUserRouter(apiGroup, db)
        api.SetupAuthRouter(apiGroup, db)
    }

    return router
}
