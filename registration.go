package main

import (
	"log"
	"sync"
)

func RegisterResource(client NmosClient, r interface{}) {
	var REGISTER_URL = "/x-nmos/registration/" + client.RegistryVersion + "/resource"
	// var REGISTER_URL = "/"
	request, err := client.PostWith(REGISTER_URL, r)
	if err != nil {
		log.Fatal(err.Error())
	}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println(string(response.Body))
}

func RegisterNodes(client NmosClient, r interface{}, ng *sync.WaitGroup) {
	defer ng.Done()
	RegisterResource(client, r)
}

func RegisterDevices(client NmosClient, devices []Device, ng *sync.WaitGroup, dg *sync.WaitGroup) {
	ng.Wait()
	defer dg.Done()
	for _, d := range devices {
		data := Data{
			"device",
			d,
		}
		RegisterResource(client, data)
	}
}

func RegisterRecievers(client NmosClient, receivers []Receiver, dg *sync.WaitGroup) {
	dg.Wait()
	for _, r := range receivers {
		data := Data{
			"receiver",
			r,
		}
		RegisterResource(client, data)
	}
}

func RegisterSources(client NmosClient, sources []Source, dg *sync.WaitGroup) {
	dg.Wait()
	for _, s := range sources {
		data := Data{
			"source",
			s,
		}
		RegisterResource(client, data)
	}
}
