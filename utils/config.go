package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var config map[string]string

func init() {
	file, _ := os.Open("config.json")
	bytestring, _ := ioutil.ReadAll(file)
	defer file.Close()
	json.Unmarshal(bytestring, &config)
}

// GetConfig ...
func GetConfig(key string) string {
	return config[key]
}
