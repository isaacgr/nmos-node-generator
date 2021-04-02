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
	genericSources := BuildGenericSources(devices, numGenericSources)
	audioSources := BuildAudioSources(devices, numAudioSources)
	dataSources := BuildDataSources(devices, numDataSources)

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
	go RegisterGenericSources(client, genericSources, &dg)
	go RegisterAudioSources(client, audioSources, &dg)
	go RegisterDataSources(client, dataSources, &dg)

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
			receiver := Receiver{}
			receiver.BuildResource(n[i], &d[i], "video", i+1)
			receivers = append(receivers, receiver)
		}
		for i := 0; i < nar; i++ {
			receiver := Receiver{}
			receiver.BuildResource(n[i], &d[i], "audio", i+1)
			receivers = append(receivers, receiver)
		}
		for i := 0; i < ndr; i++ {
			receiver := Receiver{}
			receiver.BuildResource(n[i], &d[i], "data", i+1)
			receivers = append(receivers, receiver)
		}
	}
	return receivers
}

func BuildGenericSources(d []Device, ns int) []SourceGeneric {
	sources := []SourceGeneric{}
	for i := 0; i < ns; i++ {
		for j := 0; j < len(d); j++ {
			source := SourceGeneric{}
			source.BuildResource(d[i], i+1)
			sources = append(sources, source)
		}
	}
	return sources
}

func BuildAudioSources(d []Device, ns int) []SourceAudio {
	sources := []SourceAudio{}
	for i := 0; i < ns; i++ {
		for j := 0; j < len(d); j++ {
			source := SourceAudio{}
			source.BuildResource(d[i], i+1)
			sources = append(sources, source)
		}
	}
	return sources
}

func BuildDataSources(d []Device, ns int) []SourceData {
	sources := []SourceData{}
	for i := 0; i < ns; i++ {
		for j := 0; j < len(d); j++ {
			source := SourceData{}
			source.BuildResource(d[i], i+1)
			sources = append(sources, source)
		}
	}
	return sources
}
