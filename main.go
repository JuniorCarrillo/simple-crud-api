package main

import (
	"fmt"
	"github.com/JuniorCarrillo/simple-crud-api/api/router"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = "8080"
	}

	e := router.New()
	e.Use(middleware.CORS())
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())

	listenOnPort := fmt.Sprintf(":%s", port)
	e.Logger.Fatal(e.Start(listenOnPort))
}
