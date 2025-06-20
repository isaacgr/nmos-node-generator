package is04

import (
	"errors"
	"fmt"
	"github.com/segmentio/encoding/json"
	"log"
	"strconv"

	regen "github.com/zach-klippenstein/goregen"
)

type Endpoint struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Protocol string `json:"protocol"`
}

type Api struct {
	Versions  []string   `json:"versions"`
	Endpoints []Endpoint `json:"endpoints"`
}

type NetworkDevice struct {
	ChassisId string `json:"chassis_id"`
	PortId    string `json:"port_id"`
}

type NetworkInterface struct {
	ChassisId             string         `json:"chassis_id"`
	PortId                string         `json:"port_id"`
	Name                  string         `json:"name"`
	AttachedNetworkDevice *NetworkDevice `json:"attached_network_device,omitempty"`
}

type ClockInternal struct {
	Name    string `json:"name"`
	RefType string `json:"ref_type"`
}

type ClockPTP struct {
	Name      string `json:"name"`
	RefType   string `json:"ref_type"`
	Traceable bool   `json:"traceable"`
	Version   string `json:"version"`
	Gmid      string `json:"gmid"`
	Locked    bool   `json:"locked"`
}

type Service struct {
	Href          string `json:"href"`
	Type          string `json:"type"`
	Authorization bool   `json:"authorization"`
}

type Capabilities struct{}

type Node struct {
	*ResourceCore
	Href       string             `json:"href"`
	Hostname   string             `json:"hostname"`
	Caps       Capabilities       `json:"caps"`
	Api        Api                `json:"api"`
	Services   []Service          `json:"services"`
	Clocks     []interface{}      `json:"clocks"`
	Interfaces []NetworkInterface `json:"interfaces"`
	Devices    []*Device          `json:"-"` // prevents json marshalling
}

func (n Node) Encode() ([]byte, error) {
	e, err := json.Marshal(n)

	if err != nil {
		return nil, err
	}

	return e, nil
}

func (n Node) getDevice(id string) (*Device, error) {
	for _, device := range n.Devices {
		if device.ID == id {
			return device, nil
		}
	}
	return &Device{}, errors.New("Device not found")
}

func NewNode(
	host string,
	port int,
	numInterfaces int,
	attachedNetworkDevices []NetworkDevice,
	resourceCore *ResourceCore,
) *Node {
	n := Node{
		ResourceCore: resourceCore,
	}

	endpoint := Endpoint{
		Host: host,
		Port: port,
		// TODO: Support more protocols
		Protocol: "http",
	}

	for i := range numInterfaces {
		var attachedNetworkDevice *NetworkDevice
		if i < len(attachedNetworkDevices) {
			attachedNetworkDevice = &NetworkDevice{
				ChassisId: attachedNetworkDevices[i].ChassisId,
				PortId:    attachedNetworkDevices[i].PortId,
			}
		}
		n.Interfaces = append(n.Interfaces, NetworkInterface{
			ChassisId:             GenerateMac(),
			PortId:                GenerateMac(),
			Name:                  "eth" + strconv.Itoa(i),
			AttachedNetworkDevice: attachedNetworkDevice,
		})
	}

	gmid, err := regen.Generate(
		"^[0-9a-f]{2}-[0-9a-f]{2}-[0-9a-f]{2}-[0-9a-f]{2}-[0-9a-f]{2}-[0-9a-f]{2}-[0-9a-f]{2}-[0-9a-f]{2}$",
	)

	if err != nil {
		log.Fatalf("Unable to generate gmid for clock. Error [%s]", err)
	}

	versions := []string{"v1.3"}

	internalClock := ClockInternal{
		Name:    "clk0",
		RefType: "internal",
	}
	ptpClock := ClockPTP{
		Name:      "clk1",
		RefType:   "ptp",
		Traceable: true,
		Version:   "IEEE1588-2008",
		Gmid:      gmid,
		Locked:    true,
	}
	// Hostname was deprecated
	// TODO: Provide a way to modify this
	n.Hostname = "TEST-NODE"

	// Href was deprecated
	n.Href = fmt.Sprintf(
		"%s://%s:%d",
		endpoint.Protocol,
		endpoint.Host,
		endpoint.Port,
	)
	n.Clocks = append(n.Clocks, internalClock, ptpClock)
	n.Api.Versions = versions
	n.Api.Endpoints = append(n.Api.Endpoints, endpoint)

	return &n
}
