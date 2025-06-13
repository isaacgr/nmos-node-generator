package main

import (
	"encoding/json"
	"flag"
	"io"
	"log"
	"os"
	"sync"
)

var ConfigFilename *string
var nmosConfig *Config = nil
var once sync.Once

func New() *Config {
	once.Do(func() {
		nmosConfig = loadConfig()
	})
	return nmosConfig
}

func loadConfig() *Config {
	flag.Parse()
	c := &Config{}
	var configData []byte
	var err error

	if *ConfigFilename == "-" {
		// Read from standard input
		configData, err = io.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	} else {
		// Read from file
		configFile, err := os.Open(*ConfigFilename)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		configData, err = io.ReadAll(configFile)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}

	//var config Config
	if err := json.Unmarshal(configData, &c); err != nil {
		log.Fatal("Unable to parse config file")
	}

	return c
}
