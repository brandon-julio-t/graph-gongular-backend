package main

import (
	"github.com/brandon-julio-t/graph-gongular-backend/factories"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("error loading .env file")
	}

	port := getPort()
	secret := new(factories.SecretFactory).NewSecret()
	db := new(factories.GormDatabaseFactory).NewGormDB()
	router := new(factories.ChiRouterFactory).NewRouter(secret, db)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func getPort() string {
	if port := os.Getenv("PORT"); port != "" {
		return port
	}
	return "8080"
}
