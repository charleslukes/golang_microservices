package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

const DefaultDatabase = "shopResource"
var mh *MongoHandler

func main() {
	mongoDbConnection := "mongodb://localhost:27017"
	godotenv.Load()
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("Could not find port in the environment")
	}
	
    mh = NewHandler(mongoDbConnection) 


	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))	

	// mount router
	r.Mount("/v1", shoppingRouter())

	srv := &http.Server{
		Handler: r,
		Addr:    ":" + port,
	}

	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Running Shopping server")
}
