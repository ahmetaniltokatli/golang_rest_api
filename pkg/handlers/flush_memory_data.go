package handlers

import (
	"encoding/json"
	redis "github.com/ahmetaniltokatli/golang_rest_api/pkg/db"
	"net/http"
)

func FlushMemoryData(w http.ResponseWriter, r *http.Request) {
	redisClient := redis.Initialize()
	redisData := redisClient.FlushMemoryData()
	json.NewEncoder(w).Encode(redisData)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
