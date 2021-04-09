package node

type Sender interface {
	BuildResource(n Node, d Device, f Flow, index int)
}

func BuildBaseSender(n Node, d Device, f Flow) *BaseSender {
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

func (s *SenderVideo) BuildResource(n Node, d Device, f Flow, index int) {
	s.BaseSender = BuildBaseSender(n, d, f)
	label := getResourceLabel("TestReceiverVideo", index)
	s.BaseResource = SetBaseResourceProperties(label, "NMOS Test Video Receiver")
}

func (s *SenderAudio) BuildResource(n Node, d Device, f Flow, index int) {
	s.BaseSender = BuildBaseSender(n, d, f)
	label := getResourceLabel("TestReceiverAudio", index)
	s.BaseResource = SetBaseResourceProperties(label, "NMOS Test Audio Receiver")
}
func (s *SenderData) BuildResource(n Node, d Device, f Flow, index int) {
	s.BaseSender = BuildBaseSender(n, d, f)
	label := getResourceLabel("TestReceiverData", index)
	s.BaseResource = SetBaseResourceProperties(label, "NMOS Test Data Receiver")
}
