package helpers

import (
    "time"
    "log"
    "github.com/dgrijalva/jwt-go"
    "gopkg.in/gomail.v2"
)

var jwtSecret = []byte("your_jwt_secret_key")
var passwordResetSecret = []byte("your_password_reset_secret_key")

type Claims struct {
    UserID uint `json:"user_id"`
    jwt.StandardClaims
}

// GenerateAccessToken generates an access token with a specified expiration time.
func GenerateAccessToken(userID uint) (string, error) {
    claims := &Claims{
        UserID: userID,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: time.Now().Add(time.Hour * 2).Unix(), // Token expires in 2 hours
            IssuedAt:  time.Now().Unix(),
            Issuer:    "your_application_name",
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

// GenerateRefreshToken generates a refresh token with a specified expiration time.
func GenerateRefreshToken(userID uint) (string, error) {
    claims := &Claims{
        UserID: userID,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: time.Now().AddDate(0, 0, 7).Unix(), // Refresh token expires in 7 days
            IssuedAt:  time.Now().Unix(),
            Issuer:    "your_application_name",
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

func GeneratePasswordResetToken(userID uint) (string, error) {
    claims := jwt.MapClaims{
        "sub": userID,
        "exp": time.Now().Add(time.Hour).Unix(), // Example: Token expires in 1 hour
        "iat": time.Now().Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(passwordResetSecret)
}

func VerifyPasswordResetToken(tokenString string) (*jwt.Token, error) {
    return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return passwordResetSecret, nil
    })
}

func SendEmail(to string, subject string, body string) error {
    from := "etesting053@gmail.com"
    password := "ruyonsnglowjqmpp" //use app password instead of login password

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