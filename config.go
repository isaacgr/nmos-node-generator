package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	ResourceQuantities ResourceQuantities `json:"resource"`
	Registry           Registry           `json:"registry"`
	Delay              int                `json:"node_post_delay"`
}

type FlowResource struct {
	MediaType string         `json:"media_type"`
	Sender    SenderResource `json:"sender"`
}

type GenericSource struct {
	Count int          `json:"count"`
	Flows FlowResource `json:"flows"`
}

type AudioSource struct {
	Count int          `json:"count"`
	Flows FlowResource `json:"flows"`
}

type DataSource struct {
	Count int          `json:"count"`
	Flows FlowResource `json:"flows"`
}

type SourceResource struct {
	Generic GenericSource `json:"generic"`
	Audio   AudioSource   `json:"audio"`
	Data    DataSource    `json:"data"`
}

type ReceiverDetails struct {
	Count     int    `json:"count"`
	MediaType string `json:"media_type"`
	Iface     []int  `json:"iface"`
}

type ReceiverResource struct {
	Video ReceiverDetails `json:"video"`
	Audio ReceiverDetails `json:"audio"`
	Data  ReceiverDetails `json:"data"`
}

type SenderResource struct {
	Iface []int `json:"iface"`
}

type ResourceQuantities struct {
	Nodes     int              `json:"nodes"`
	Devices   int              `json:"devices"`
	Receivers ReceiverResource `json:"receivers"`
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
		log.Fatal("Unable to read config file")
	}

	return c

}
