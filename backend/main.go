// Package main Oxide API
//
// # Oxide backend API with math operations and utilities
//
// Terms Of Service: http://swagger.io/terms/
//
// Schemes: http, https
// Host: localhost:8000
// BasePath: /
// Version: 1.0.0
// Contact: Oxide Team <contact@oxide.io> http://oxide.io
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// swagger:meta
package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"oxide/db"
	_ "oxide/docs"
	"oxide/handler"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

// rootHandler handles the root endpoint
// @Summary Root endpoint
// @Description Returns a simple hello message
// @Tags general
// @Produce plain
// @Success 200 {string} string "hello"
// @Router / [get]
func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

// helloHandler handles the hello endpoint
// @Summary Hello endpoint
// @Description Returns a hello message from backend
// @Tags general
// @Produce plain
// @Success 200 {string} string "hello from backend"
// @Router /api/hello [get]
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello from backend")
	w.Write([]byte("hello from backend"))
}

// addHandler handles the addition endpoint
// @Summary Add two numbers
// @Description Add two integers provided as query parameters
// @Tags math
// @Param a query int true "First number"
// @Param b query int true "Second number"
// @Produce plain
// @Success 200 {string} string "Sum of the two numbers"
// @Failure 400 {string} string "Invalid parameters"
// @Router /api/add [get]
func addHandler(w http.ResponseWriter, r *http.Request) {
	aStr := r.URL.Query().Get("a")
	bStr := r.URL.Query().Get("b")

	a, errA := strconv.Atoi(aStr)
	b, errB := strconv.Atoi(bStr)

	if errA != nil || errB != nil {
		http.Error(w, "invalid params", http.StatusBadRequest)
		return
	}

	sum := a + b
	fmt.Fprintf(w, "%d", sum)
}

// CORS middleware to allow frontend requests
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// paginate is a stub middleware function for pagination
func paginate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Add pagination logic here if needed
		next.ServeHTTP(w, r)
	})
}

// main starts the server
// @title Oxide API
// @version 1.0
// @description Oxide backend API with math operations and utilities
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.oxide.io/support
// @contact.email support@oxide.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8000
// @BasePath /
func main() {
	// Initialize database connection and run migrations
	fmt.Println("Initializing database...")
	db.Init()

	router := chi.NewRouter()

	router.Use(corsMiddleware)
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)

	router.Use(middleware.Timeout(60 * time.Second))

	router.Get("/", rootHandler)

	// Swagger documentation - accessible at both /swagger and /docs
	router.Get("/swagger/*", httpSwagger.WrapHandler)

	router.Route("/api", func(r chi.Router) {
		r.With(paginate).Get("/hello", helloHandler)
		r.With(paginate).Get("/add", addHandler)
		r.With(paginate).Get("/user", handler.UserHandler)
		r.With(paginate).Post("/user/add", handler.AddUserHandler)
		r.With(paginate).Post("/user/delete", handler.DeleteUserHandler)

		// r.With(paginate).Get("fetch-repo", func(w http.ResponseWriter, r *http.Request) {
		// 	url := r.URL.Query().Get("url")

		// })
	})

	fmt.Println("Starting server on :8000")
	serverErr := http.ListenAndServe(":8000", router)
	if serverErr != nil {
		fmt.Printf("Server failed to start: %v\n", serverErr)
	}
}
