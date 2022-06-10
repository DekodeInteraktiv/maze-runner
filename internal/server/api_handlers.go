package server

import (
	"encoding/json"
	"log"
	"net/http"
)

// error holds data about an error
type error struct {
	Err string `json:"error"`
}

func writeJSON(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Panicf("Failed to encode JSON: %v\n", err)
	}
}
