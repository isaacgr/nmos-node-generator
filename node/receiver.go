package node

import "github.com/isaacgr/nmos-node-generator/config"

type Receiver interface {
	BuildResource(n Node, d *Device, index int)
}

var ReceiverConfig = config.New().ResourceQuantities.Receivers

func BuildBaseReceiver(n Node, d *Device, b []int) *BaseReceiver {
	r := BaseReceiver{}
	c := ReceiverCaps{}
	for i := range b {
		r.InterfaceBindings = append(r.InterfaceBindings, n.Interfaces[i].Name)
	}
	r.DeviceId = d.ID
	r.Transport = ReceiverTransport
	r.Caps = c
	return &r
}

func (r *ReceiverVideo) BuildResource(n Node, d *Device, index int) {
	r.BaseReceiver = BuildBaseReceiver(n, d, ReceiverConfig.Video.Iface)
	label := getResourceLabel("TestReceiverVideo", index)
	r.BaseResource = SetBaseResourceProperties(label, "NMOS Test Video Receiver")
	r.Format = VideoFormat
	r.Caps.MediaTypes = append(r.Caps.MediaTypes, VideoMediaTypes[ReceiverConfig.Video.MediaType])
	d.Receivers = append(d.Receivers, r.ID)

}

func (r *ReceiverAudio) BuildResource(n Node, d *Device, index int) {
	r.BaseReceiver = BuildBaseReceiver(n, d, ReceiverConfig.Audio.Iface)
	label := getResourceLabel("TestReceiverAudio", index)
	r.BaseResource = SetBaseResourceProperties(label, "NMOS Test Audio Receiver")
	r.Format = AudioFormat
	r.Caps.MediaTypes = append(r.Caps.MediaTypes, AudioMediaTypes[ReceiverConfig.Audio.MediaType])
	d.Receivers = append(d.Receivers, r.ID)

}

func (r *ReceiverData) BuildResource(n Node, d *Device, index int) {
	r.BaseReceiver = BuildBaseReceiver(n, d, ReceiverConfig.Data.Iface)
	label := getResourceLabel("TestReceiverData", index)
	r.BaseResource = SetBaseResourceProperties(label, "NMOS Test Data Receiver")
	r.Format = DataFormat
	r.Caps.MediaTypes = append(r.Caps.MediaTypes, DataMediaTypes[ReceiverConfig.Data.MediaType])
	d.Receivers = append(d.Receivers, r.ID)

}
