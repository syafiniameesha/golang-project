package middleware

import (
    "net/http"
    "strings"
    "user-management/helpers"

    "github.com/gin-gonic/gin"
)

// AuthMiddleware is the middleware to verify JWT tokens.
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
            c.Abort()
            return
        }

        tokenString := strings.TrimSpace(authHeader)
        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is empty"})
            c.Abort()
            return
        }

        claims, err := helpers.VerifyToken(tokenString)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        // Set the user ID from the token into the context
        c.Set("userID", claims.UserID)
        c.Next()
    }
}
