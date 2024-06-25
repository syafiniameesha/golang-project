package utils

import (
    "errors"
    "github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("your-secret-key")

// ValidateToken validates the JWT token and returns the user ID
func ValidateToken(tokenString string) (int, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, errors.New("unexpected signing method")
        }
        return secretKey, nil
    })

    if err != nil {
        return 0, err
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        userID, ok := claims["user_id"].(float64)
        if !ok {
            return 0, errors.New("user_id not found in token claims")
        }
        return int(userID), nil
    } else {
        return 0, errors.New("invalid token")
    }
}
