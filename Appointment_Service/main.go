package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
)

func main() {

	//Loead Enviroment Varriable in go Program
	Enverr := godotenv.Load()

	if Enverr != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize router
	r := gin.Default()

	// Simple health check route
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "appointment service is running"})
	})

	// Get appointments (Placeholder)
	r.GET("/api/appointments", func(c *gin.Context) {
		// Here you would interact with DB to get appointments
		c.JSON(http.StatusOK, gin.H{"appointments": []string{"Appointment 1", "Appointment 2"}})
	})

	// Get specific appointment (Placeholder)
	r.GET("/api/appointments/:id", func(c *gin.Context) {
		id := c.Param("id")

		writer := kafka.NewWriter(kafka.WriterConfig{
			Brokers: []string{os.Getenv("KAFKA_BROKER")},
			Topic:   "appointments",
		})

		defer writer.Close()

		err := writer.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte("booking"),
			Value: []byte(id),
		})
		if err != nil {
			log.Printf("Failed to write message: %v", err)
		}
		// Here you would query DB for the specific appointment using `id`
		c.JSON(http.StatusOK, gin.H{"appointment_id": id, "status": "Messeged Puhsed to kafka"})
	})

	// Connect to the database (Postgres)
	// DB connection logic will go here later based on your DB client (e.g., pgx or GORM)

	// Start server on port 4242
	port := os.Getenv("PORT")
	if port == "" {
		port = "4242"
	}
	fmt.Println("Appointment Service running on port " + port)
	r.Run(":" + port) // listen on port 4242
}
