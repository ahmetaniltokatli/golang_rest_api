package handlers

import (
	"encoding/json"
	redis "github.com/ahmetaniltokatli/golang_rest_api/pkg/db"
	"github.com/ahmetaniltokatli/golang_rest_api/pkg/models"
	"github.com/gorilla/mux"
	"net/http"
)

func GetMemoryData(w http.ResponseWriter, r *http.Request) {
	var response models.Response
	vars := mux.Vars(r)
	key := vars["key"]

	redisClient := redis.Initialize()
	redisData := redisClient.GetKey(key)

	response.Success = true
	response.Data = redisData

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
}
