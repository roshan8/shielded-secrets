package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	startServer()
}

func startServer() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Get("/api/secrets", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]interface{}{
			"data": []map[string]interface{}{
				{
					"id":          1,
					"name":        "Apple MacBook Pro 17",
					"description": "Laptop",
				},
				{
					"id":          2,
					"name":        "Microsoft Surface Pro",
					"description": "Tab",
				},
				{
					"id":          3,
					"name":        "Magic Mouse 2",
					"description": "Accessary",
				},
			},
			"meta": map[string]interface{}{
				"total": 3,
			},
		}
		jData, _ := json.Marshal(data)

		w.Header().Set("Content-Type", "application/json")
		w.Write(jData)
	})

	fmt.Println("Server started at port 8080...")
	http.ListenAndServe(":8080", r)
}
