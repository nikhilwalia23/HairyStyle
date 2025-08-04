package utils

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"database/sql"
	"errors"
	"my-go-server/config"

	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

var jwtSecret []byte

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Unable to load enviroment file")
	}

	secret := os.Getenv("JWT_SECRET")

	jwtSecret = []byte(secret)
}

// GenerateJWT creates a signed token with user info
func GenerateJWT(userID, email string) (string, error) {
	claims := jwt.MapClaims{
		"user_id":      userID,
		"emailOrPhone": email,
		"exp":          time.Now().Add(time.Hour * 24).Unix(), // expires in 24h
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func AuthenticateUser(emailOrPhone, password string) (string, error) {
	var userID, storedPassword string

	// Simple regex to check if it's an email
	isEmail := regexp.MustCompile(`^[^@\s]+@[^@\s]+\.[^@\s]+$`).MatchString(emailOrPhone)

	var query string
	if isEmail {
		query = `SELECT id, password FROM auth_application_schema.users WHERE email = $1`
	} else {
		query = `SELECT id, password FROM auth_application_schema.users WHERE phone = $1`
	}

	err := config.DB.QueryRow(query, emailOrPhone).Scan(&userID, &storedPassword)

	if err == sql.ErrNoRows {
		return "", errors.New("invalid email/phone or password")
	} else if err != nil {
		return "", err
	}

	if storedPassword != password {
		return "", errors.New("invalid email/phone or password")
	}

	return userID, nil
}

func LoggingMiddleware(c *gin.Context) {
	fmt.Println("Request received")
	c.Next() // Call the next middleware or handler
}

func JWTAuthMiddleware(c *gin.Context) {
	// Get the token from Authorization header
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token not provided"})
		c.Abort()
		return
	}

	// Extract the token string
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		// Secret key from env
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		c.Abort()
		return
	}

	// Optional: extract claims and set in context
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		c.Set("user_id", claims["user_id"])
		c.Set("email", claims["email"])
		// Add more claims as needed
	}

	c.Next()
}
