package main

import (
	"crypto/tls"
	"flag"
	"log"
	"math"
	"net/http"
	"sync"
	"time"

	"github.com/isaacgr/nmos-node-generator/client"
	"github.com/isaacgr/nmos-node-generator/config"
	"github.com/isaacgr/nmos-node-generator/util"
)

var configFile = flag.String(
	"config",
	"config.json",
	"Conifg file containing resource generation info",
)
var useRandomNode = flag.Bool(
	"random-node-id",
	true,
	"Pass this flag to use a non-random UUID for the node",
)
var useRandomDevice = flag.Bool(
	"random-device-id",
	true,
	"Pass this flag to use a non-random UUID for the device",
)
var useRandomResource = flag.Bool(
	"random-resource-id",
	true,
	"Pass this flag to use a non-random UUID for the device's resources",
)
var requestTimeout = flag.Int(
	"request-timeout",
	20,
	"Set the timeout for HTTP requests",
)
var noKeepalive = flag.Bool(
	"connection-keepalive",
	true,
	"Pass this flag to use a persistent HTTP connection",
)

func main() {

	flag.Parse()
	config.ConfigFilename = configFile
	randomDeviceUUID := *useRandomDevice
	randomResourceUUID := *useRandomResource
	randomNodeUUID := *useRandomNode
	httpRequestTimeout := time.Duration(*requestTimeout)
	connectionKeepalive := *noKeepalive

	config := config.New()
	baseUrl := config.Registry.Scheme + "://" + config.Registry.IP
	port := config.Registry.Port

	transport := &http.Transport{
		DisableKeepAlives:   false,
		MaxIdleConns:        0,
		MaxIdleConnsPerHost: math.MaxInt64,
		IdleConnTimeout:     300 * time.Second,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	if connectionKeepalive == false {
		transport = &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}
	}
	httpclient := &http.Client{
		Transport: transport,
		Timeout:   httpRequestTimeout * time.Second,
	}
	c := client.NmosClient{
		baseUrl,
		port,
		config.Registry.Version,
		httpclient,
		transport,
	}
	var KEEPALIVE_URL = "/x-nmos/registration/" + c.RegistryVersion + "/health/nodes/"
	var ng sync.WaitGroup
	var dg sync.WaitGroup
	var sg sync.WaitGroup
	var fg sync.WaitGroup

	numNodes := config.ResourceQuantities.Nodes.Count
	numInterfaces := config.ResourceQuantities.Nodes.NumInterfaces
	if numNodes == 0 {
		log.Fatal("Must define at least one node")
	}
	if numInterfaces == 0 {
		log.Fatal("Must define at least one interface")
	}
	numDevices := config.ResourceQuantities.Devices.Count
	deviceIp := config.ResourceQuantities.Devices.IpAddress
	devicePortStart := config.ResourceQuantities.Devices.PortStart
	if numDevices == 0 {
		log.Fatal("Must define at least one device")
	}
	numVideoReceivers := config.ResourceQuantities.Receivers.Video.Count
	numAudioReceivers := config.ResourceQuantities.Receivers.Audio.Count
	numDataReceivers := config.ResourceQuantities.Receivers.Data.Count
	numGenericSources := config.ResourceQuantities.Sources.Generic.Count
	numAudioSources := config.ResourceQuantities.Sources.Audio.Count
	numDataSources := config.ResourceQuantities.Sources.Data.Count
	videoFlowType := config.ResourceQuantities.Sources.Generic.Flows.MediaType
	audioFlowType := config.ResourceQuantities.Sources.Audio.Flows.MediaType
	dataFlowType := config.ResourceQuantities.Sources.Data.Flows.MediaType

	nodes := util.BuildNodes(
		numNodes,
		numInterfaces,
		config.ResourceQuantities.Nodes.NamePrefix,
		config.ResourceQuantities.Nodes.AttachedNetworkDevices,
		randomNodeUUID,
	)
	devices := util.BuildDevices(
		nodes,
		numDevices,
		config.ResourceQuantities.NamePrefix,
		deviceIp,
		devicePortStart,
		randomDeviceUUID,
	)
	receivers := util.BuildReceivers(
		nodes,
		devices,
		numVideoReceivers,
		numAudioReceivers,
		numDataReceivers,
		randomResourceUUID,
	)
	sources := util.BuildSources(
		devices,
		numGenericSources,
		numAudioSources,
		numDataSources,
		randomResourceUUID,
	)
	flows := util.BuildFlows(
		devices,
		sources,
		videoFlowType,
		audioFlowType,
		dataFlowType,
		randomResourceUUID,
	)
	senders := util.BuildSenders(nodes, devices, flows, randomResourceUUID)

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
	sg.Add(1)
	fg.Add(1)
	go client.RegisterDevices(c, devices, &ng, &dg)
	go client.RegisterRecievers(c, receivers, &dg)
	go client.RegisterSources(c, sources, &dg, &sg)
	go client.RegisterFlows(c, flows, &sg, &fg)
	go client.RegisterSenders(c, senders, &fg)

	for n := range k {
		go func(n string) {
			if n == "404" {
				for _, n := range nodes {
					data := client.Data{
						"node",
						n,
					}
					ng.Add(1)
					go client.RegisterNodes(c, data, &ng)
				}
				dg.Add(1)
				sg.Add(1)
				fg.Add(1)
				go client.RegisterDevices(c, devices, &ng, &dg)
				go client.RegisterRecievers(c, receivers, &dg)
				go client.RegisterSources(c, sources, &dg, &sg)
				go client.RegisterFlows(c, flows, &sg, &fg)
				go client.RegisterSenders(c, senders, &fg)

			}
		}(n)

	}

}

func NodeKeepalive(
	client client.NmosClient,
	url string,
	k chan string,
	ng *sync.WaitGroup,
) {
	ng.Wait()
	time.Sleep(5 * time.Second)
	client.Keepalive(url, k)
}
