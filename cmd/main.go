package main

import (
	"fmt"
	redis "github.com/ahmetaniltokatli/golang_rest_api/pkg/db"
	"github.com/ahmetaniltokatli/golang_rest_api/pkg/handlers"
	"github.com/ahmetaniltokatli/golang_rest_api/pkg/models"
	"github.com/gorilla/mux"
	"github.com/robfig/cron"
	"log"
	"net/http"
	"time"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/getMemoryData/{key}", handlers.GetMemoryData).Methods(http.MethodGet)
	router.HandleFunc("/setMemoryData", handlers.SetMemoryData).Methods(http.MethodPost)
	router.HandleFunc("/flushMemoryData", handlers.FlushMemoryData).Methods(http.MethodDelete)

	filePath := "tmp/data.json"
	WriteCacheDataFromFile(filePath)

	c := cron.New()
	c.AddFunc("@every 1m", func() {
		fmt.Println("Every 1 min writes data to json file")
		redisClient := redis.Initialize()
		allKeys := redisClient.GetAllKeys()

		var data []models.RedisData
		for _, element := range allKeys {
			newStruct := models.RedisData{
				Key:   element,
				Value: redisClient.GetKey(element),
			}

			data = append(data, newStruct)
		}

		handlers.WriteJsonFile(filePath, data)

	})
	c.Start()

	log.Println("API is running!")
	http.ListenAndServe(":4000", router)
}

func WriteCacheDataFromFile(fileName string) {
	existFile := handlers.ExistFile(fileName)

	if existFile {
		var data []models.RedisData
		data = handlers.ReadJsonFile(fileName)
		redisClient := redis.Initialize()
		for _, element := range data {
			redisClient.SetKey(element.Key, element.Value, time.Minute*60)
		}
	}
}
