package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoggingMiddleware(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/blogs", nil)
	recorder := httptest.NewRecorder()

	handler := LoggingMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	handler.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestValidatePostMiddleware(t *testing.T) {
	routes := http.NewServeMux()
	routes.Handle("/blog", ValidatePostMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})))

	validBlog := Blog{
		Title:       "Valid Title",
		Description: "Valid Description",
		Body:        "Valid Body",
	}

	body, _ := json.Marshal(validBlog)
	req, _ := http.NewRequest(http.MethodPost, "/blog", bytes.NewBuffer(body))
	recorder := httptest.NewRecorder()
	routes.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusOK, recorder.Code)

	// Test case: Malformed Body (read error)
	req.Body = io.NopCloser(errReader(0))
	recorder = httptest.NewRecorder()
	routes.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusBadRequest, recorder.Code)

	// Test case: Invalid JSON
	req, _ = http.NewRequest(http.MethodPost, "/blog", bytes.NewBuffer([]byte("Invalid JSON")))
	recorder = httptest.NewRecorder()
	routes.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusBadRequest, recorder.Code)

	// Test case: Missing Required Fields
	invalidBlog := Blog{Title: "", Description: "", Body: ""}
	body, _ = json.Marshal(invalidBlog)
	req, _ = http.NewRequest(http.MethodPost, "/blog", bytes.NewBuffer(body))
	recorder = httptest.NewRecorder()
	routes.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

type errReader int

func (errReader) Read(p []byte) (int, error) {
	return 0, io.ErrUnexpectedEOF
}
