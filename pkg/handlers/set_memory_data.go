package handlers

import (
	"encoding/json"
	redis "github.com/ahmetaniltokatli/golang_rest_api/pkg/db"
	"github.com/ahmetaniltokatli/golang_rest_api/pkg/models"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const REDIS_DATA_TTL = 60

func SetMemoryData(w http.ResponseWriter, r *http.Request) {
	var response models.Response
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		response.Success = false
		response.ErrorMessage = err.Error()
		log.Fatalln(err)
	}

	var redisDataModel models.RedisData
	response.Success = true
	json.Unmarshal(body, &redisDataModel)

	redisClient := redis.Initialize()
	redisData := redisClient.SetKey(redisDataModel.Key, redisDataModel.Value, time.Minute*REDIS_DATA_TTL)

	if redisData != nil {
		response.Success = false
		response.ErrorMessage = redisData.Error()
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
}
