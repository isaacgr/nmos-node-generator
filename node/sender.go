package node

type Sender interface {
	BuildResource(n Node, d *Device, f Flow, index int)
}

func BuildBaseSender(n Node, d *Device, f Flow) *BaseSender {
	s := BaseSender{}
	for i, iface := range n.Interfaces {
		s.InterfaceBindings = append(s.InterfaceBindings, iface.Name)
		if i == 1 {
			break
		}
	}
	s.DeviceId = d.ID
	s.FlowId = f.getId()
	s.Transport = SenderTransport
	return &s
}

func (s *SenderVideo) BuildResource(n Node, d *Device, f Flow, index int) {
	s.BaseSender = BuildBaseSender(n, d, f)
	label := getResourceLabel("TestSenderVideo", index)
	s.BaseResource = SetBaseResourceProperties(label, "NMOS Test Video Sender")
	d.Senders = append(d.Senders, s.ID)
}

func (s *SenderAudio) BuildResource(n Node, d *Device, f Flow, index int) {
	s.BaseSender = BuildBaseSender(n, d, f)
	label := getResourceLabel("TestSenderAudio", index)
	s.BaseResource = SetBaseResourceProperties(label, "NMOS Test Audio Sender")
	d.Senders = append(d.Senders, s.ID)
}
func (s *SenderData) BuildResource(n Node, d *Device, f Flow, index int) {
	s.BaseSender = BuildBaseSender(n, d, f)
	label := getResourceLabel("TestSenderData", index)
	s.BaseResource = SetBaseResourceProperties(label, "NMOS Test Data Sender")
	d.Senders = append(d.Senders, s.ID)
}
