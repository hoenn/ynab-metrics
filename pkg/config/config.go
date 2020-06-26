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
		panic("Could not open config file")
	}
	defer jsonFile.Close()
	jsonBytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic("Could not read config file")
	}
	var c Config
	err = json.Unmarshal(jsonBytes, &c)
	if err != nil {
		panic("Could not unmarshal config file")
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
