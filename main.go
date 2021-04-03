package main

import (
	"flag"
	"log"
	"sync"
	"time"
)

var configFilename = flag.String("config", "configs/config.json", "Conifg file containing resource generation info")

type Data struct {
	Type string      `json:"type"`
	Node interface{} `json:"data"`
}

func main() {

	flag.Parse()
	config := ParseConfig(*configFilename)
	baseUrl := config.Registry.Scheme + "://" + config.Registry.IP
	port := config.Registry.Port
	client := NmosClient{
		baseUrl,
		port,
		"v1.2",
	}
	var KEEPALIVE_URL = "/x-nmos/registration/" + client.RegistryVersion + "/health/nodes/"
	// var KEEPALIVE_URL = "/"
	var ng sync.WaitGroup
	var dg sync.WaitGroup

	numNodes := config.ResourceQuantities.Nodes
	if numNodes == 0 {
		log.Fatal("Must define at least one node")
	}
	numDevices := config.ResourceQuantities.Devices
	numVideoReceivers := config.ResourceQuantities.Receivers.Video
	numAudioReceivers := config.ResourceQuantities.Receivers.Audio
	numDataReceivers := config.ResourceQuantities.Receivers.Data
	numGenericSources := config.ResourceQuantities.Sources.Generic
	numAudioSources := config.ResourceQuantities.Sources.Audio
	numDataSources := config.ResourceQuantities.Sources.Data

	nodes := BuildNodes(numNodes)
	devices := BuildDevices(nodes, numDevices)
	receivers := BuildReceivers(nodes, devices, numVideoReceivers, numAudioReceivers, numDataReceivers)
	sources := BuildSources(devices, numGenericSources, numAudioSources, numDataSources)

	k := make(chan string)
	for _, n := range nodes {
		data := Data{
			"node",
			n,
		}
		ng.Add(1)
		go RegisterNodes(client, data, &ng)
		go NodeKeepalive(client, KEEPALIVE_URL+n.ID, k, &ng)
	}
	dg.Add(1)
	go RegisterDevices(client, devices, &ng, &dg)
	go RegisterRecievers(client, receivers, &dg)
	go RegisterSources(client, sources, &dg)

	for n := range k {
		go func(n string) {
			log.Printf("Keepalive [%s]", n)
			time.Sleep(5 * time.Second)
			client.Keepalive(KEEPALIVE_URL+n, k)
		}(n)
	}
}

func NodeKeepalive(client NmosClient, url string, k chan string, ng *sync.WaitGroup) {
	ng.Wait()
	client.Keepalive(url, k)
}

func BuildNodes(nn int) []Node {
	nodes := []Node{}
	for i := 0; i < nn; i++ {
		node := Node{}
		node.BuildResource(i + 1)
		nodes = append(nodes, node)
	}
	return nodes
}

func BuildDevices(n []Node, nd int) []Device {
	devices := []Device{}
	for i := 0; i < nd; i++ {
		for j := 0; j < len(n); j++ {
			device := Device{}
			device.BuildResource(n[j], i+1)
			devices = append(devices, device)
		}
	}
	return devices
}

func BuildReceivers(n []Node, d []Device, nvr int, nar int, ndr int) []Receiver {
	receivers := []Receiver{}
	for j := 0; j < len(d); j++ {
		for i := 0; i < nvr; i++ {
			receiver := ReceiverVideo{}
			receiver.BuildResource(n[i], &d[i], i+1)
			receivers = append(receivers, &receiver)
		}
		for i := 0; i < nar; i++ {
			receiver := ReceiverAudio{}
			receiver.BuildResource(n[i], &d[i], i+1)
			receivers = append(receivers, &receiver)
		}
		for i := 0; i < ndr; i++ {
			receiver := ReceiverData{}
			receiver.BuildResource(n[i], &d[i], i+1)
			receivers = append(receivers, &receiver)
		}
	}
	return receivers
}

func BuildSources(d []Device, ngs int, nas int, nds int) []Source {
	sources := []Source{}
	for j := 0; j < len(d); j++ {
		for i := 0; i < ngs; i++ {
			source := SourceGeneric{}
			source.BuildResource(d[i], i+1)
			sources = append(sources, &source)
		}
		for i := 0; i < nas; i++ {
			source := SourceAudio{}
			source.BuildResource(d[i], i+1)
			sources = append(sources, &source)
		}
		for i := 0; i < nds; i++ {
			source := SourceData{}
			source.BuildResource(d[i], i+1)
			sources = append(sources, &source)
		}
	}
	return sources
}
