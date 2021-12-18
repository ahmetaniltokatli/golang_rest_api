package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func GetMemoryData(w http.ResponseWriter, r *http.Request) {
	// Read dynamic id parameter
	vars := mux.Vars(r)
	key := vars["key"]
	json.NewEncoder(w).Encode(key)
	redisClient := initialize()

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(key)
}
