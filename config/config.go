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
		config = loadConfig()
	})
	return config
}

//func parseConfig() *Config {
//	flag.Parse()
//	c := &Config{}
//	configFile, err := os.Open(*ConfigFilename)
//	if err != nil {
//		log.Fatal(err)
//		os.Exit(1)
//	}
//
//	byteValue, err := ioutil.ReadAll(configFile)
//	if err != nil {
//		log.Fatal(err)
//		os.Exit(1)
//	}
//	err = json.Unmarshal(byteValue, &c)
//	if err != nil {
//		log.Fatal("Unable to parse config file")
//	}
//	return c
//}

func loadConfig() *Config {
	flag.Parse()
	c := &Config{}
    var configData []byte
    var err error

    if *ConfigFilename == "-" {
        // Read from standard input
        configData, err = ioutil.ReadAll(os.Stdin)
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
        configData, err = ioutil.ReadAll(configFile)
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
