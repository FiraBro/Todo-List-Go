package main

import (
	"encoding/json"
	"net/http"
)

// writeJSON sends a JSON response with the given status code and data
func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

// readJSON reads JSON data from the request body and decodes it into the given destination
func readJSON(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := 1_048_576 // 1MB limit
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	return decoder.Decode(data)
}

// handleJSONError provides detailed, user-friendly error messages for JSON decoding
func writeJSONError(w http.ResponseWriter,status int, message string) error {
	type envelop struct{
		Error string `json:error`
	}
	return writeJSON(w,status, &envelop{Error: message})
}
