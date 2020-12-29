package main

import (
	"github.com/brandon-julio-t/graph-gongular-backend/factories/chi-router"
	"github.com/brandon-julio-t/graph-gongular-backend/factories/gorm-database"
	"github.com/brandon-julio-t/graph-gongular-backend/factories/secret"
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
	appSecret := new(secret.Factory).Create()
	db := new(gorm_database.Factory).Create()
	router := new(chi_router.Factory).Create(appSecret, db)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func getPort() string {
	if port := os.Getenv("PORT"); port != "" {
		return port
	}
	return "8080"
}
