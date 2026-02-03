package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors" // Added for CORS
)

func main() {
	r := chi.NewRouter()

	// CORS Middleware
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4321"}, // Allow requests from Astro dev server
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// A good set of default middleware
	r.Use(middleware.Logger)    // Log requests to console
	r.Use(middleware.Recoverer) // Panic recovery
	r.Use(middleware.Timeout(60 * time.Second))

	// Basement Bloomberg route
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "Basement Bloomberg"}`))
	})

	// Health check for Docker/Cloudflare
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	r.Get("/api/stocks/{ticker}", func(w http.ResponseWriter, r *http.Request) {
		ticker := chi.URLParam(r, "ticker")
		
		// Mock data - later you'll get this from a real API or DB
		responseData := map[string]interface{}{
			"symbol": ticker,
			"price":  150.25,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
	})

	// Start server
	println("Server starting on :8080...")
	http.ListenAndServe(":8080", r)
}