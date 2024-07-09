package helpers

import (
    "time"
    "log"
    "github.com/dgrijalva/jwt-go"
    "gopkg.in/gomail.v2"
)

var jwtSecret = []byte("123456")
var passwordResetSecret = []byte("123456")

type Claims struct {
    UserID uint `json:"ID"`
    jwt.StandardClaims
}

// GenerateAccessToken generates an access token 
func GenerateAccessToken(userID uint) (string, error) {
    claims := &Claims{
        UserID: userID,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: time.Now().Add(time.Hour * 1).Unix(), // Token expires in 1 hour
            IssuedAt:  time.Now().Unix(),
            Issuer:    "User_Management",
            Subject:   "access_token",
        },
    }
    accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := accessToken.SignedString(jwtSecret)
    if err != nil {
        return "", err
    }
    return tokenString, nil
}

// VerifyToken parses and verifies a JWT token.
func VerifyToken(tokenString string) (*Claims, error) {
    claims := &Claims{}
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtSecret, nil
    })

    if err != nil || !token.Valid {
        return nil, err
    }
    return claims, nil
}

// GenerateRefreshToken generates a refresh token  
func GenerateRefreshToken(userID uint) (string, error) {
    claims := &Claims{
        UserID: userID,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: time.Now().AddDate(0, 0, 7).Unix(), // Refresh token expires in 7 days
            IssuedAt:  time.Now().Unix(),
            Issuer:    "User_Management",
            Subject:   "refresh_token",
        },
    }
    refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := refreshToken.SignedString(jwtSecret)
    if err != nil {
        return "", err
    }
    return tokenString, nil
}

// generate token for reset password
func GeneratePasswordResetToken(userID uint) (string, error) {
    claims := jwt.MapClaims{
        "sub": userID,
        "exp": time.Now().Add(time.Hour).Unix(), // Token expires in 1 hour
        "iat": time.Now().Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(passwordResetSecret)
}

// verify reset password token
func VerifyPasswordResetToken(tokenString string) (*jwt.Token, error) {
    return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return passwordResetSecret, nil
    })
}

// SendEmail sends an email to the specified domain
func SendEmail(to string, subject string, body string) error {
    from := "etesting053@gmail.com"
    password := "" //use app password instead of login password

    // Setup SMTP server
    smtpHost := "smtp.gmail.com"
    smtpPort := 587

    // Message
    m := gomail.NewMessage()
    m.SetHeader("From", from)
    m.SetHeader("To", to)
    m.SetHeader("Subject", subject)
    m.SetBody("text/html", body)

    // Authentication and sending email
    d := gomail.NewDialer(smtpHost, smtpPort, from, password)

    log.Printf("Sending email to: %s with subject: %s", to, subject)

    if err := d.DialAndSend(m); err != nil {
        log.Printf("Failed to send email: %v", err)
        return err
    }

    log.Println("Email sent successfully")
    return nil
}
