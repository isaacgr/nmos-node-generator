package node

import "github.com/isaacgr/nmos-node-generator/config"

type Receiver interface {
	BuildResource(n Node, d *Device, index int, useRandomResource bool)
}

func getReceiverConfig() config.ReceiverResource {
	return config.New().ResourceQuantities.Receivers
}

func BuildBaseReceiver(n Node, d *Device, b []int) *BaseReceiver {
	r := BaseReceiver{}
	c := ReceiverCaps{}
	for _, iface := range b {
		r.InterfaceBindings = append(r.InterfaceBindings, n.Interfaces[iface-1].Name)
	}
	r.DeviceId = d.ID
	r.Transport = ReceiverTransport
	r.Caps = c
	return &r
}

func (r *ReceiverVideo) BuildResource(n Node, d *Device, index int, useRandomResource bool) {
	r.BaseReceiver = BuildBaseReceiver(n, d, getReceiverConfig().Video.Iface)
	label := getResourceLabel(d.Label+"."+"ReceiverVideo", index)
	r.BaseResource = SetBaseResourceProperties(label, "NMOS Test Video Receiver", useRandomResource)
	r.Format = VideoFormat
	r.Caps.MediaTypes = append(r.Caps.MediaTypes, VideoMediaTypes[getReceiverConfig().Video.MediaType])
	d.Receivers = append(d.Receivers, r.ID)

}

func (r *ReceiverAudio) BuildResource(n Node, d *Device, index int, useRandomResource bool) {
	r.BaseReceiver = BuildBaseReceiver(n, d, getReceiverConfig().Audio.Iface)
	label := getResourceLabel(d.Label+"."+"ReceiverAudio", index)
	r.BaseResource = SetBaseResourceProperties(label, "NMOS Test Audio Receiver", useRandomResource)
	r.Format = AudioFormat
	r.Caps.MediaTypes = append(r.Caps.MediaTypes, AudioMediaTypes[getReceiverConfig().Audio.MediaType])
	d.Receivers = append(d.Receivers, r.ID)

}

func (r *ReceiverData) BuildResource(n Node, d *Device, index int, useRandomResource bool) {
	r.BaseReceiver = BuildBaseReceiver(n, d, getReceiverConfig().Data.Iface)
	label := getResourceLabel(d.Label+"."+"ReceiverData", index)
	r.BaseResource = SetBaseResourceProperties(label, "NMOS Test Data Receiver", useRandomResource)
	r.Format = DataFormat
	r.Caps.MediaTypes = append(r.Caps.MediaTypes, DataMediaTypes[getReceiverConfig().Data.MediaType])
	d.Receivers = append(d.Receivers, r.ID)

}
