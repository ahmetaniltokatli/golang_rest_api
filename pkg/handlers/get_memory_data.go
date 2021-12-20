package handlers

import (
	"encoding/json"
	redis "github.com/ahmetaniltokatli/golang_rest_api/pkg/db"
	"github.com/gorilla/mux"
	"net/http"
)

func GetMemoryData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	redisClient := redis.Initialize()
	redisData := redisClient.GetKey(key)
	json.NewEncoder(w).Encode(redisData)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
