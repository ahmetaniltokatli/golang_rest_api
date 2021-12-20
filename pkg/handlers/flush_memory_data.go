package handlers

import (
	"encoding/json"
	redis "github.com/ahmetaniltokatli/golang_rest_api/pkg/db"
	"github.com/ahmetaniltokatli/golang_rest_api/pkg/models"
	"log"
	"net/http"
	"os"
)

func FlushMemoryData(w http.ResponseWriter, r *http.Request) {
	var response models.Response
	response.Success = ExistFile("tmp/data.json")

	redisClient := redis.Initialize()
	redisClient.FlushMemoryData()
	existFile := ExistFile("tmp/data.json")

	if existFile {
		e := os.Remove("tmp/data.json")
		if e != nil {
			response.Success = false
			response.ErrorMessage = e.Error()
			log.Fatal(e)
		}
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
}
