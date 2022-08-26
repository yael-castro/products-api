// Package handler is the presentation layer, contains everything related to the layer that will be exposed to the end user (in this case an HTTP Client)
package handler

import (
	"encoding/json"
	"fmt"
	"github.com/yael-castro/layered-architecture/internal/model"
	"log"
	"net/http"
)

// Configuration are the settings for initialize an instance of Handler
type Configuration struct {
	// NotFound handle requests made to non exist paths
	NotFound http.Handler
	// HealthCheck handle requests made to know the path
	HealthCheck http.Handler
	// MethodNotAllowed handle requests made to existing paths but with wrong method
	MethodNotAllowed http.Handler
}

// New builds an instance of *Handler using the Configuration
func New(config Configuration) *Handler {
	h := &Handler{
		NotFound:         config.NotFound,
		MethodNotAllowed: config.MethodNotAllowed,
	}

	h.handlers = make(map[string]map[string]http.Handler)

	h.SetHandler("/", http.MethodGet, config.HealthCheck)

	return h
}

// Handler contains all instances of http.Handler to setting the endpoints
type Handler struct {
	NotFound         http.Handler
	MethodNotAllowed http.Handler
	handlers         map[string]map[string]http.Handler
}

// SetHandler configure a Handler to path and method
func (h *Handler) SetHandler(path, method string, handler http.Handler) {
	_, ok := h.handlers[path]
	if !ok {
		h.handlers[path] = make(map[string]http.Handler)
	}

	h.handlers[path][method] = handler
}

// ServeHTTP handle every
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	log.Println(path)

	methods, ok := h.handlers[path]
	if !ok {
		h.NotFound.ServeHTTP(w, r)
		return
	}

	handler, ok := methods[r.Method]
	if !ok {
		h.MethodNotAllowed.ServeHTTP(w, r)
		return
	}

	handler.ServeHTTP(w, r)
}

// NotFound default http.Handler to handle requests made to no exists path
func NotFound(w http.ResponseWriter, r *http.Request) {
	JSON(w, http.StatusOK, model.Map{"message": fmt.Sprintf(`path '%s' does not exists`, r.URL.Path)})
}

// HealthCheck default http.Handler to make a health check status
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write(nil)
}

// JSON sends the data encoded in JSON format via HTTP
func JSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}
