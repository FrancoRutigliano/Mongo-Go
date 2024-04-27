package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Data estructure json -> Decode -> Data estructure Go
func ParseJSON(r *http.Request, payload any) error {
	// empty body
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}

	return json.NewDecoder(r.Body).Decode(payload)
}

func WriteJson(w http.ResponseWriter, status int, v any) error {
	// agregar a cabeceras content -> json
	w.Header().Add("Content-Type", "application/json")
	// add status
	w.WriteHeader(status)

	// encode data --> json
	return json.NewEncoder(w).Encode(v)
}

func WriteErr(w http.ResponseWriter, status int, err error) {
	WriteJson(w, status, map[string]string{"error": err.Error()})
}
