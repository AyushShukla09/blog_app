package api

import (
	response "blog_app/utils"

	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

// Request Logging Middleware
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("%s %s %s", r.Method, r.RequestURI, time.Since(start))
		next.ServeHTTP(w, r)
	})
}

// Create/Update Blog Request Body Validation Middleware
func ValidatePostMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost || r.Method == http.MethodPut {
			body, err := io.ReadAll(r.Body)
			if err != nil {
				response.Fail(w, http.StatusBadRequest, err.Error())
				return
			}
			var post Blog
			if err := json.Unmarshal(body, &post); err != nil {
				response.Fail(w, http.StatusBadRequest, err.Error())
				return
			}
			if post.Title == "" || post.Description == "" || post.Body == "" {
				response.Fail(w, http.StatusBadRequest, "All fields (title, description, body) are required")
				return
			}

			r.Body = io.NopCloser(io.MultiReader(bytes.NewReader(body)))
		}

		next.ServeHTTP(w, r)
	})
}
