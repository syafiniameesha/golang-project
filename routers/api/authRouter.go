package api

import (
    "user-management/controllers"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

func SetupAuthRouter(router *gin.RouterGroup, db *gorm.DB) {
    authController := controllers.NewAuthController(db)

    router.POST("/signup", authController.Signup)
    router.POST("/login", authController.Login)
    router.POST("/logout", authController.Logout)
    router.POST("/forgotPassword", authController.ForgotPassword)
    router.POST("/resetPassword/:token", authController.ResetPassword)
}