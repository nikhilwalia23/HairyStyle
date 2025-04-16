package main

import (
	"fmt"
	"os"

	"log"
	"net/http"

	"github.com/quic-go/quic-go/http3"

	_ "github.com/jackc/pgx/v5/stdlib" // PostgreSQL driver
	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"

	"my-go-server/config"
	"my-go-server/routes"
)

func main() {

	//Loead Enviroment Varriable in go Program
	Enverr := godotenv.Load()

	// Initialize database
	config.DB_Connnection()

	if Enverr != nil {
		log.Fatal("Error loading .env file")
	}

	APP_PORT := os.Getenv("APP_PORT_LOCAL")

	PORT_ADDRESS := ":" + APP_PORT

	//Set Up server
	r := gin.Default()

	// Register your user routes
	routes.RegisterUserRoutes(r)

	// Log the server start
	fmt.Println("Starting HTTP/3 server on https://localhost" + PORT_ADDRESS)

	// Start listening and serving using HTTP/3
	err := http3.ListenAndServeTLS(PORT_ADDRESS, "./cert.pem", "./key.pem", r)

	if err != nil {
		log.Fatalf("Failed to configure HTTP/3: %v", err)
	}

}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Proto)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "pong"}`))

}
