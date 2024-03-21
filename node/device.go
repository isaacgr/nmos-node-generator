package node

import "fmt"

func (d *Device) BuildResource(n Node, numdevices int, nodeidx int, index int, nameprefix string, deviceip string, devicePortStart int, useRandom bool) {
	label := getResourceLabel(n.Label+"."+nameprefix, index)
	d.BaseResource = SetBaseResourceProperties(label, "NMOS Test Device", useRandom)
	d.NodeId = n.ID
	deviceport := devicePortStart + ((nodeidx - 1) * numdevices) + (index - 1)
	deviceHref := fmt.Sprintf("http://%s:%d/", deviceip, deviceport)

	c1 := Controls{
		n.Href,
		"urn:x-nmos:control:sr-ctrl/v1.1",
		false,
	}
	c2 := Controls{
		deviceHref,
		"urn:x-nmos:control:sr-ctrl/v1.1",
		false,
	}
	d.Type = "urn:x-nmos:device:generic"
	d.Controls = append(d.Controls, c1, c2)
}
