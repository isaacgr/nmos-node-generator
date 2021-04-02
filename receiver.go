package main

import "log"

func (r *Receiver) BuildResource(n Node, d *Device, resourceType string, index int) {
	c := ReceiverCaps{}
	for i, iface := range n.Interfaces {
		r.InterfaceBindings = append(r.InterfaceBindings, iface.Name)
		if i == 1 {
			break
		}
	}
	r.DeviceId = d.ID
	r.Transport = ReceiverTransport
	switch resourceType {
	case "video":
		label := getResourceLabel("TestReceiverVideo", index)
		r.BaseResource = SetBaseResourceProperties(label, "NMOS Test Video Receiver")
		r.Format = VideoFormat
		c.MediaTypes = append(c.MediaTypes, VideoMediaTypes["raw"])
	case "audio":
		label := getResourceLabel("TestReceiverAudio", index)
		r.BaseResource = SetBaseResourceProperties(label, "NMOS Test Audio Receiver")
		r.Format = AudioFormat
		c.MediaTypes = append(c.MediaTypes, AudioMediaTypes[16])
	case "data":
		label := getResourceLabel("TestReceiverData", index)
		r.BaseResource = SetBaseResourceProperties(label, "NMOS Test Data Receiver")
		r.Format = DataFormat
		c.MediaTypes = append(c.MediaTypes, DataMediaTypes["json"])
	default:
		log.Fatal("No valid receiver type provided")
	}
	r.Caps = c
	d.Receivers = append(d.Receivers, r.ID)
}
