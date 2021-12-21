package handlers

import (
	"encoding/json"
	redis "github.com/ahmetaniltokatli/golang_rest_api/pkg/db"
	"github.com/ahmetaniltokatli/golang_rest_api/pkg/models"
	"net/http"
)

func FlushMemoryData(w http.ResponseWriter, r *http.Request) {
	var response models.Response
	response.Success = true

	redisClient := redis.Initialize()
	err := redisClient.FlushMemoryData()

	if err != nil {
		response.Success = false
		response.ErrorMessage = err.Error()
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
}
