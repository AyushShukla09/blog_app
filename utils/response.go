package utils

import (
	"encoding/json"
	"net/http"
)

// Response implements standard JSON response payload structure.
type Response struct {
	Status string          `json:"status"`
	Error  *ResponseError  `json:"error,omitempty"`
	Result json.RawMessage `json:"result,omitempty"`
}

type ResponseError struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Details []string `json:"details"`
}

// Fail sends an unsuccessful JSON response
func Fail(w http.ResponseWriter, status int, details ...string) {
	r := &Response{
		Status: StatusFail,
		Error: &ResponseError{
			Code:    status,
			Details: details,
		},
	}
	j, err := json.Marshal(r)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(j)
}

// Send sends a successful JSON response.
func Send(w http.ResponseWriter, status int, result any) {
	rj, err := json.Marshal(result)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	r := &Response{
		Status: StatusOK,
		Result: rj,
	}
	j, err := json.Marshal(r)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(j)
}

// ResponseStatus constants
const (
	StatusOK   = "ok"
	StatusFail = "nok"
)
