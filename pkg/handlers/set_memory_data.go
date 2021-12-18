package handlers

import (
	"encoding/json"
	"net/http"
)

func SetMemoryData(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("GetMemoryData")
}
