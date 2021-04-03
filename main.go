package main

import (
	"flag"
	"log"
	"sync"
	"time"

	"github.com/isaacgr/nmos-node-generator/client"
	"github.com/isaacgr/nmos-node-generator/configuration"
	"github.com/isaacgr/nmos-node-generator/node"
)

var configFilename = flag.String("config", "configs/config.json", "Conifg file containing resource generation info")

func main() {

	flag.Parse()
	config := configuration.ParseConfig(*configFilename)
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

func BuildNodes(nn int) []node.Node {
	nodes := []node.Node{}
	for i := 0; i < nn; i++ {
		node := node.Node{}
		node.BuildResource(i + 1)
		nodes = append(nodes, node)
	}
	return nodes
}

func BuildDevices(n []node.Node, nd int) []node.Device {
	devices := []node.Device{}
	for i := 0; i < nd; i++ {
		for j := 0; j < len(n); j++ {
			device := node.Device{}
			device.BuildResource(n[j], i+1)
			devices = append(devices, device)
		}
	}
	return devices
}

func BuildReceivers(n []node.Node, d []node.Device, nvr int, nar int, ndr int) []node.Receiver {
	receivers := []node.Receiver{}
	for j := 0; j < len(d); j++ {
		for i := 0; i < nvr; i++ {
			receiver := node.ReceiverVideo{}
			receiver.BuildResource(n[i], &d[i], i+1)
			receivers = append(receivers, &receiver)
		}
		for i := 0; i < nar; i++ {
			receiver := node.ReceiverAudio{}
			receiver.BuildResource(n[i], &d[i], i+1)
			receivers = append(receivers, &receiver)
		}
		for i := 0; i < ndr; i++ {
			receiver := node.ReceiverData{}
			receiver.BuildResource(n[i], &d[i], i+1)
			receivers = append(receivers, &receiver)
		}
	}
	return receivers
}

func BuildSources(d []node.Device, ngs int, nas int, nds int) []node.Source {
	sources := []node.Source{}
	for j := 0; j < len(d); j++ {
		for i := 0; i < ngs; i++ {
			source := node.SourceGeneric{}
			source.BuildResource(d[i], i+1)
			sources = append(sources, &source)
		}
		for i := 0; i < nas; i++ {
			source := node.SourceAudio{}
			source.BuildResource(d[i], i+1)
			sources = append(sources, &source)
		}
		for i := 0; i < nds; i++ {
			source := node.SourceData{}
			source.BuildResource(d[i], i+1)
			sources = append(sources, &source)
		}
	}
	return sources
}
