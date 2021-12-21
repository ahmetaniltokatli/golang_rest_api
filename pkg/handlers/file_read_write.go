package handlers

import (
	"encoding/json"
	"github.com/ahmetaniltokatli/golang_rest_api/pkg/models"
	"io/ioutil"
	"os"
)

func ExistFile(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func CheckFile(filename string) error {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		_, err := os.Create(filename)
		if err != nil {
			return err
		}
	}
	return nil
}

func ReadJsonFile(filename string) []models.RedisData {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return []models.RedisData{}
	}

	data := []models.RedisData{}

	json.Unmarshal(file, &data)

	return data
}

func WriteJsonFile(filename string, data []models.RedisData) {
	CheckFile(filename)
	file, _ := json.MarshalIndent(data, "", " ")
	_ = ioutil.WriteFile(filename, file, 0644)
}

func AppendJsonFile(filename string, key string, value string) error {
	CheckFile(filename)
	data := ReadJsonFile(filename)
	newStruct := &models.RedisData{
		Key:   key,
		Value: value,
	}

	data = append(data, *newStruct)

	// Preparing the data to be marshalled and written.
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, dataBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}
