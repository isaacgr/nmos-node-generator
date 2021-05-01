package config

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

var ConfigFilename *string
var config *Config = nil
var once sync.Once

func New() *Config {
	once.Do(func() {
		config = parseConfig()
	})
	return config
}

func parseConfig() *Config {
	flag.Parse()
	c := &Config{}
	configFile, err := os.Open(*ConfigFilename)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	byteValue, err := ioutil.ReadAll(configFile)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	err = json.Unmarshal(byteValue, &c)
	if err != nil {
		log.Fatal("Unable to parse config file")
	}
	return c
}
