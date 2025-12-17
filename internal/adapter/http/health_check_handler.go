package http

import (
	"encoding/json"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data := struct {
			Status string `json:"status"`
	}{
			Status: "ok",
	}

	json.NewEncoder(w).Encode(data)
}