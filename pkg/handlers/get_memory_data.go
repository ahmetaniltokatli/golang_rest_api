package handlers

import (
	"encoding/json"
	redis "github.com/ahmetaniltokatli/golang_rest_api/pkg/db"
	"github.com/gorilla/mux"
	"net/http"
)

func GetMemoryData(w http.ResponseWriter, r *http.Request) {

	redis.Initialize()
	//redis.SetKey("anil", test)
	// Read dynamic id parameter
	vars := mux.Vars(r)
	key := vars["key"]
	json.NewEncoder(w).Encode(key)
	json.NewEncoder(w).Encode(anil)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
