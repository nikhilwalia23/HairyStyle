package main

import (
	"log"
	"net/http"
	"os"

	"context"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

var db *pgxpool.Pool

func initDB() {
	dbUrl := os.Getenv("DATABASE_URL")
	var err error
	db, err = pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
}

func main() {
	initDB()
	defer db.Close()

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "user service running"})
	})

	r.Run(":4202")
}
