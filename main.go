package main

import (
	"flag"
	"log"
	"sync"
	"time"

	"github.com/isaacgr/nmos-node-generator/client"
	"github.com/isaacgr/nmos-node-generator/util"
)

var configFilename = flag.String("config", "config.json", "Conifg file containing resource generation info")

func main() {

	flag.Parse()
	config := ParseConfig(*configFilename)
	baseUrl := config.Registry.Scheme + "://" + config.Registry.IP
	port := config.Registry.Port
	c := client.NmosClient{
		baseUrl,
		port,
		"v1.2",
	}
	var KEEPALIVE_URL = "/x-nmos/registration/" + c.RegistryVersion + "/health/nodes/"
	// var KEEPALIVE_URL = "/"
	var ng sync.WaitGroup
	var dg sync.WaitGroup

	numNodes := config.ResourceQuantities.Nodes
	if numNodes == 0 {
		log.Fatal("Must define at least one node")
	}
	numDevices := config.ResourceQuantities.Devices
	if numDevices == 0 {
		log.Fatal("Must define at least one device")
	}
	numVideoReceivers := config.ResourceQuantities.Receivers.Video
	numAudioReceivers := config.ResourceQuantities.Receivers.Audio
	numDataReceivers := config.ResourceQuantities.Receivers.Data
	numGenericSources := config.ResourceQuantities.Sources.Generic
	numAudioSources := config.ResourceQuantities.Sources.Audio
	numDataSources := config.ResourceQuantities.Sources.Data

	nodes := util.BuildNodes(numNodes)
	devices := util.BuildDevices(nodes, numDevices)
	receivers := util.BuildReceivers(nodes, devices, numVideoReceivers, numAudioReceivers, numDataReceivers)
	sources := util.BuildSources(devices, numGenericSources, numAudioSources, numDataSources)

	k := make(chan string)
	for _, n := range nodes {
		data := client.Data{
			"node",
			n,
		}
		ng.Add(1)
		go client.RegisterNodes(c, data, &ng)
		go NodeKeepalive(c, KEEPALIVE_URL+n.ID, k, &ng)
	}
	dg.Add(1)
	go client.RegisterDevices(c, devices, &ng, &dg)
	go client.RegisterRecievers(c, receivers, &dg)
	go client.RegisterSources(c, sources, &dg)

	for n := range k {
		go func(n string) {
			log.Printf("Keepalive [%s]", n)
			time.Sleep(5 * time.Second)
			c.Keepalive(KEEPALIVE_URL+n, k)
		}(n)
	}
}

func NodeKeepalive(client client.NmosClient, url string, k chan string, ng *sync.WaitGroup) {
	ng.Wait()
	client.Keepalive(url, k)
}
