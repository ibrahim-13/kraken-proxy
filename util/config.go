package util

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	KrakenApiKey       string
	KrakenPrivateKey   string
	Host               string
	Port               string
	ServerCertPath     string
	ServerKeyPath      string
	EnableSsl          bool
	EnableOtherRequest bool
}

var configData Config

func init() {
	dir, _ := os.Getwd()
	configFile := filepath.Join(dir, "config.json")
	if !IsFileExist(configFile) {
		bytes, _ := json.MarshalIndent(&configData, "", "\t")
		os.WriteFile(configFile, bytes, 0755)
		panic("config.json file not found, sample output emitted")
	}
	bytes, err := os.ReadFile(configFile)
	PanicIfNotNil(err)
	err = json.Unmarshal(bytes, &configData)
	PanicIfNotNil(err)
}

func GetConfig() Config {
	return configData
}
