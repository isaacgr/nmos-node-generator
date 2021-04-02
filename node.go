package main

import (
	"log"

	regen "github.com/zach-klippenstein/goregen"
)

func (n *Node) BuildResource(index int) {
	// build out node with some default values
	endpoint := Endpoint{
		"172.16.220.69",
		3000,
		"http",
	}
	gmid, err := regen.Generate("^[0-9a-f]{2}-[0-9a-f]{2}-[0-9a-f]{2}-[0-9a-f]{2}-[0-9a-f]{2}-[0-9a-f]{2}-[0-9a-f]{2}-[0-9a-f]{2}$")
	if err != nil {
		log.Fatal("Unable to generate gmid for clock")
	}
	versions := []string{"v1.3", "v1.2"}
	attachedNetworkDevice := NetworkDevice{
		GenerateMac(),
		GenerateMac(),
	}

	iface1 := NetworkInterface{
		GenerateMac(),
		GenerateMac(),
		"eth0",
		attachedNetworkDevice,
	}
	iface2 := NetworkInterface{
		GenerateMac(),
		GenerateMac(),
		"eth1",
		attachedNetworkDevice,
	}

	internalClock := ClockInternal{
		"clk0",
		"internal",
	}
	ptpClock := ClockPTP{
		"clk1",
		"ptp",
		true,
		"IEEE1588-2008",
		gmid,
		true,
	}

	clock1 := &internalClock
	clock2 := &ptpClock
	label := getResourceLabel("TestNode", index)
	n.BaseResource = SetBaseResourceProperties(label, "NMOS Test Node")
	n.Href = "http://172.16.220.69:4003"
	n.Hostname = "TEST-NODE"
	n.Api.Endpoints = append(n.Api.Endpoints, endpoint)
	n.Api.Versions = versions
	n.Interfaces = append(n.Interfaces, iface1, iface2)
	n.Clocks = append(n.Clocks, clock1, clock2)
}
