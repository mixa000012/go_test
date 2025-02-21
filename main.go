package main

import (
	_ "awesomeProject/docs" // Подключение Swagger-документации
	"awesomeProject/handlers"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
)

func GetCORSConfig() cors.Options {
	return cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080", "http://127.0.0.1:8080"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/redis/incr", handlers.HandleIncrement)
	mux.HandleFunc("/sign/hmacsha512", handlers.HandleHMACSHA512)
	mux.HandleFunc("/postgres/users", handlers.HandleCreateUser)
	mux.Handle("/swagger/", httpSwagger.Handler(httpSwagger.URL("http://localhost:8080/swagger/doc.json")))

	corsHandler := cors.New(GetCORSConfig())

	log.Println("API запущен на :8080")
	log.Fatal(http.ListenAndServe(":8080", corsHandler.Handler(mux)))
}
