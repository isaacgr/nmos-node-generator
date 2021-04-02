package main

func (d *Device) BuildResource(n Node, index int) {
	label := getResourceLabel("TestDevice", index)
	d.BaseResource = SetBaseResourceProperties(label, "NMOS Test Device")
	d.NodeId = n.ID
	c := Controls{
		n.Href,
		"urn:x-manufacturer:control:generic",
		false,
	}
	d.Type = "urn:x-nmos:device:generic"
	d.Controls = append(d.Controls, c)
}
