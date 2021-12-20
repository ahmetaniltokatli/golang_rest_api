package handlers

import (
	"encoding/json"
	redis "github.com/ahmetaniltokatli/golang_rest_api/pkg/db"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func GetMemoryData(w http.ResponseWriter, r *http.Request) {

	plan, _ := ioutil.ReadFile("./tmp/test.json")
	var data interface{}
	err := json.Unmarshal(plan, &data)
	json.NewEncoder(w).Encode(err)
	json.NewEncoder(w).Encode(data)

	vars := mux.Vars(r)
	key := vars["key"]

	redisClient := redis.Initialize()
	redisData := redisClient.GetKey(key)
	json.NewEncoder(w).Encode(redisData)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
