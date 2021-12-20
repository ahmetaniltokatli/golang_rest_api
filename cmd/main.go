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
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/getMemoryData/{key}", handlers.GetMemoryData).Methods(http.MethodGet)
	router.HandleFunc("/setMemoryData", handlers.SetMemoryData).Methods(http.MethodPost)
	router.HandleFunc("/flushMemoryData", handlers.FlushMemoryData).Methods(http.MethodGet)

	c := cron.New()
	c.AddFunc("@every 1m", func() {
		fmt.Println("Every 1 min writes data to json file")
		redisClient := redis.Initialize()
		allKeys := redisClient.GetAllKeys()

		var data []models.RedisData
		filePath := "tmp/data.json"
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
