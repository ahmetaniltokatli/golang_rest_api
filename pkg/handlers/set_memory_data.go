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
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var redisDataModel models.RedisData
	json.Unmarshal(body, &redisDataModel)

	redisClient := redis.Initialize()
	redisData := redisClient.SetKey(redisDataModel.Key, redisDataModel.Value, time.Minute*REDIS_DATA_TTL)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(redisDataModel.Value)
	json.NewEncoder(w).Encode(redisData)
}
