package handlers

import (
	redis "github.com/ahmetaniltokatli/golang_rest_api/pkg/db"
	"log"
	"net/http"
	"os"
)

func FlushMemoryData(w http.ResponseWriter, r *http.Request) {
	redisClient := redis.Initialize()
	redisClient.FlushMemoryData()
	e := os.Remove("tmp/data.json")
	if e != nil {
		log.Fatal(e)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
