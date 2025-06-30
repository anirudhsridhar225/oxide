package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

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

func main() {
	router := chi.NewRouter()

	router.Use(corsMiddleware)
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)

	router.Use(middleware.Timeout(60 * time.Second))

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})

	router.Route("/api", func(r chi.Router) {
		r.With(paginate).Get("/hello", func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("hello from backend")
			w.Write([]byte("hello from backend"))
		})

		r.With(paginate).Get("/add", func(w http.ResponseWriter, r *http.Request) {
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
		})
	})

	fmt.Println("Starting server on :8000")
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
}
