package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

//ParseConfig parses the configuration json
func ParseConfig(cfgFile string) *Config {
	jsonFile, err := os.Open(cfgFile)
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()
	jsonBytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}
	var c Config
	err = json.Unmarshal(jsonBytes, &c)
	if err != nil {
		panic(err)
	}
	return &c
}

//Config represents the configuration for the exporter
type Config struct {
	Port            string `json:"port"`
	GetTrans        bool   `json:"include_transactions"`
	AccessToken     string `json:"access_token"`
	IntervalSeconds uint64 `json:"interval_seconds"`
}
