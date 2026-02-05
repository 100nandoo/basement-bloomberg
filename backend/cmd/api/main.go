package main

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	rest "github.com/100nandoo/basement-bloomberg/backend/internal/rest"
)

func main() {
	r := chi.NewRouter()

	// CORS Middleware (Still useful for dev, though less critical in single container)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4321", "http://localhost:8080"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	// API Routes
	r.Route("/api", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"message": "Basement Bloomberg API"}`))
		})

		r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		r.Get("/stocks/{ticker}", func(w http.ResponseWriter, r *http.Request) {
			ticker := chi.URLParam(r, "ticker")

			// Use the rest package to get quote summary
			params := rest.QuoteSummaryQuery{
				Modules:    "assetProfile,summaryDetail,price", // Default modules
				CorsDomain: "finance.yahoo.com",
				Formatted:  false,
				Symbol:     ticker,
			}

			summary, err := rest.GetQuoteSummary(ticker, params)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(summary)
		})
	})

	// Serve Static Files (Frontend)
	// We assume index.html and assets are in "./dist"
	workDir, _ := os.Getwd()
	filesDir := filepath.Join(workDir, "dist")
	FileServer(r, "/", http.Dir(filesDir))

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	println("Server starting on :" + port + "...")
	http.ListenAndServe(":"+port, r)
}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
