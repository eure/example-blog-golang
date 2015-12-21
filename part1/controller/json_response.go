package controller

import (
	"encoding/json"
	"net/http"
)

// status codes
const (
	StatusOK      = 200
	StatusCreated = 201
)

// JSONResponse is alias of map for JSON response
type JSONResponse struct {
	data   map[string]interface{}
	status int
}

// NewResponse creates a new JSONResponse
func NewResponse() *JSONResponse {
	r := &JSONResponse{
		data:   make(map[string]interface{}),
		status: StatusOK,
	}
	return r
}

// Merge adds multiple map data to the response
func (r *JSONResponse) Merge(data map[string]interface{}) {
	for k, v := range data {
		r.data[k] = v
	}
}

// Add adds a single key-value to the response
func (r *JSONResponse) Add(key string, value interface{}) {
	r.data[key] = value
}

// SetCreated set http status code to 201
func (r *JSONResponse) SetCreated() {
	r.status = StatusCreated
}

// RenderJSON render json response
func RenderJSON(w http.ResponseWriter, msg interface{}) {
	switch v := msg.(type) {
	case *JSONResponse:
		if _, ok := v.data["error"]; !ok {
			v.data["error"] = nil
		}
		w.WriteHeader(v.status)
		msg = v.data
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msg)
}
