package middleware

import (
    "log"
    "net/http"
    "strings"
    "user-management/utils"

    "github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Get the token from the Authorization header
        authHeader := c.GetHeader("Authorization")
        log.Println("Authorization Header:", authHeader)
        if authHeader == "" {
            log.Println("Authorization header is missing")
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
            c.Abort()
            return
        }

        token := strings.TrimPrefix(authHeader, "Bearer ")
        log.Println("Token:", token)
        if token == "" {
            log.Println("Token is missing")
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is missing"})
            c.Abort()
            return
        }

        // Validate and extract user ID from token
        userID, err := utils.ValidateToken(token)
        if err != nil {
            log.Println("Token validation error:", err)
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        log.Println("UserID from Token:", userID)

        // Set the user ID in the context
        c.Set("userID", userID)
        c.Next()
    }
}
