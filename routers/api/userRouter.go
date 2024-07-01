package api

import (
    "user-management/controllers"
    "user-management/middleware"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

func SetupUserRouter(router *gin.RouterGroup, db *gorm.DB) {
    userController := controllers.NewUserController(db)
    
    // Add other user routes here
    router.POST("/users", userController.CreateUser)
    router.PUT("/users/:id", middleware.AuthMiddleware(), userController.UpdateUser)
    router.GET("/userDetails", middleware.AuthMiddleware(), userController.GetUserDetails)
    router.GET("/userList", middleware.AuthMiddleware(), userController.GetUserList)
    // router.POST("/userDetails", userController.GetUserDetailsByToken)
}


