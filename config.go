package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	ResourceQuantities ResourceQuantities `json:"resource"`
	Registry           Registry           `json:"registry"`
	Delay              int                `json:"node_post_delay"`
}

type SourceResource struct {
	Generic int `json:"generic"`
	Audio   int `json:"audio"`
	Data    int `json:"data"`
}

type ReceiverResource struct {
	Video int `json:"video"`
	Audio int `json:"audio"`
	Data  int `json:"data"`
}

type SenderResource struct {
	Video int `json:"video"`
	Audio int `json:"audio"`
	Data  int `json:"data"`
}

type ResourceQuantities struct {
	Nodes     int              `json:"nodes"`
	Devices   int              `json:"devices"`
	Senders   SenderResource   `json:"senders"`
	Receivers ReceiverResource `json:"receivers"`
	Flows     int              `json:"flows"`
	Sources   SourceResource   `json:"sources"`
}

type Registry struct {
	IP     string `json:"ip"`
	Port   int    `json:"port"`
	Scheme string `json:"scheme"`
}

func ParseConfig(fname string) Config {
	c := Config{}
	configFile, err := os.Open(fname)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	byteValue, err := ioutil.ReadAll(configFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = json.Unmarshal(byteValue, &c)
	if err != nil {
		fmt.Println("Unable to read config file")
	}

	return c

}
