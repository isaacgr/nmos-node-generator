package node

type Receiver interface {
	BuildResource(n Node, d *Device, index int)
}

func BuildBaseReceiver(n Node, d *Device) *BaseReceiver {
	r := BaseReceiver{}
	c := ReceiverCaps{}
	for i, iface := range n.Interfaces {
		r.InterfaceBindings = append(r.InterfaceBindings, iface.Name)
		if i == 1 {
			break
		}
	}
	r.DeviceId = d.ID
	r.Transport = ReceiverTransport
	r.Caps = c
	return &r
}

func (r *ReceiverVideo) BuildResource(n Node, d *Device, index int) {
	r.BaseReceiver = BuildBaseReceiver(n, d)
	label := getResourceLabel("TestReceiverVideo", index)
	r.BaseResource = SetBaseResourceProperties(label, "NMOS Test Video Receiver")
	r.Format = VideoFormat
	r.Caps.MediaTypes = append(r.Caps.MediaTypes, VideoMediaTypes["raw"])
	d.Receivers = append(d.Receivers, r.ID)

}

func (r *ReceiverAudio) BuildResource(n Node, d *Device, index int) {
	r.BaseReceiver = BuildBaseReceiver(n, d)
	label := getResourceLabel("TestReceiverAudio", index)
	r.BaseResource = SetBaseResourceProperties(label, "NMOS Test Audio Receiver")
	r.Format = AudioFormat
	r.Caps.MediaTypes = append(r.Caps.MediaTypes, AudioMediaTypes[16])
	d.Receivers = append(d.Receivers, r.ID)

}

func (r *ReceiverData) BuildResource(n Node, d *Device, index int) {
	r.BaseReceiver = BuildBaseReceiver(n, d)
	label := getResourceLabel("TestReceiverData", index)
	r.BaseResource = SetBaseResourceProperties(label, "NMOS Test Data Receiver")
	r.Format = DataFormat
	r.Caps.MediaTypes = append(r.Caps.MediaTypes, DataMediaTypes["json"])
	d.Receivers = append(d.Receivers, r.ID)

}
