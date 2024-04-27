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
