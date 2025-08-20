package controllers

import (
	"fmt"
	"my-go-server/config"
	"my-go-server/kafka"
	"my-go-server/models"
	"my-go-server/utils"

	"context"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Event model
type UserRegisteredEvent struct {
	EventType string    `json:"eventType"`
	UserID    string    `json:"userId"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Timestamp time.Time `json:"timestamp"`
}

func SendUserCreatedEvent() {
	// Kafka writer from your connection package
	writer := kafka.NewKafkaWriter([]string{"localhost:9092"}, "user-events")

	msg := kafka.PrepareMessage("user1", "User registered successfully")
	err := kafka.SendMessage(context.Background(), writer, msg)
	if err != nil {
		panic(err)
	}

	fmt.Println("✅ Message sent successfully!")
}

// Function to register user
func RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Validation to check Details
	if len(user.AadharCard) != 12 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Aadhar card must be 12 digits"})
		return
	}

	if len(user.Phone) < 10 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Phone must be of at least 10 Digit"})
		return
	}

	if len(user.Password) < 8 || len(user.Password) > 20 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password must be 8 to 20 characters long"})
		return
	}

	err := config.DB.Ping()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user: " + err.Error()})
		return
	}

	query := `
		INSERT INTO auth_application_schema.users (id, name, role, email, phone, is_verified, aadhar_card, password)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	id := uuid.New().String()

	_, err = config.DB.Exec(
		query,
		id, user.Name, user.Role, user.Email, user.Phone,
		user.IsVerified, user.AadharCard, user.Password,
	)

	if err != nil {
		//Check Phone Number and Email Already Registered
		if strings.Contains(err.Error(), "users_phone_unique") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Phone number already registered"})
			return
		}
		if strings.Contains(err.Error(), "users_email_key") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email already registered"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, id)
}
func ProtectedTesting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"token": "working Ok"})
	SendUserCreatedEvent()
	c.Next()
}

func Login(c *gin.Context) {
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	userId, errr := utils.AuthenticateUser(input.Email, input.Password)

	if errr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errr.Error()})
		return
	}

	token, err := utils.GenerateJWT(userId, input.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token generation failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
