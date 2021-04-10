package client

import (
	"log"
	"sync"

	"github.com/isaacgr/nmos-node-generator/node"
)

type Data struct {
	Type string      `json:"type"`
	Node interface{} `json:"data"`
}

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

func RegisterDevices(client NmosClient, devices []node.Device, ng *sync.WaitGroup, dg *sync.WaitGroup) {
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

func RegisterRecievers(client NmosClient, receivers []node.Receiver, dg *sync.WaitGroup) {
	dg.Wait()
	for _, r := range receivers {
		data := Data{
			"receiver",
			r,
		}
		RegisterResource(client, data)
	}
}

func RegisterSources(client NmosClient, sources []node.Source, dg *sync.WaitGroup, sg *sync.WaitGroup) {
	dg.Wait()
	defer sg.Done()
	for _, s := range sources {
		data := Data{
			"source",
			s,
		}
		RegisterResource(client, data)
	}
}

func RegisterFlows(client NmosClient, flows []node.Flow, sg *sync.WaitGroup, fg *sync.WaitGroup) {
	sg.Wait()
	defer fg.Done()
	for _, f := range flows {
		data := Data{
			"flow",
			f,
		}
		RegisterResource(client, data)
	}
}
