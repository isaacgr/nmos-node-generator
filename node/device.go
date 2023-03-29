package node

func (d *Device) BuildResource(n Node, index int, nameprefix string, useRandom bool) {
	label := getResourceLabel(n.Label+"."+nameprefix, index)
	d.BaseResource = SetBaseResourceProperties(label, "NMOS Test Device", useRandom)
	d.NodeId = n.ID
	c1 := Controls{
		n.Href,
		"urn:x-nmos:control:sr-ctrl/v1.0",
		false,
	}
	c2 := Controls{
		"http://172.16.169.169:4003",
		"urn:x-nmos:control:sr-ctrl/v1.0",
		false,
	}
	d.Type = "urn:x-nmos:device:generic"
	d.Controls = append(d.Controls, c1, c2)
}
