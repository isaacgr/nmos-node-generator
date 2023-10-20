package node

import (
	"log"
	"strconv"

	"github.com/isaacgr/nmos-node-generator/config"
	regen "github.com/zach-klippenstein/goregen"
)

func (n *Node) BuildResource(index int, numInterfaces int, namePrefix string, attachedNetworkDevices []config.AttachedNetworkDevices) {
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
	// attachedNetworkDevice :=

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

	for i := 0; i < numInterfaces; i++ {
		n.Interfaces = append(n.Interfaces, NetworkInterface{
			GenerateMac(),
			GenerateMac(),
			"eth" + strconv.Itoa(i),
			NetworkDevice{
				attachedNetworkDevices[i].ChassisID,
				attachedNetworkDevices[i].PortID,
			},
		})
	}

	clock1 := &internalClock
	clock2 := &ptpClock
	label := getResourceLabel(namePrefix, index)
	n.BaseResource = SetBaseResourceProperties(label, "NMOS Test Node")
	n.Href = "http://172.16.169.69:4003/"
	n.Hostname = "TEST-NODE"
	n.Api.Endpoints = append(n.Api.Endpoints, endpoint)
	n.Api.Versions = versions
	n.Clocks = append(n.Clocks, clock1, clock2)
}
