package main

import "log"

func (s *Sender) BuildResource(n Node, d Device, f *BaseFlow, t string, index int) {
	for i, iface := range n.Interfaces {
		s.InterfaceBindings = append(s.InterfaceBindings, iface.Name)
		if i == 1 {
			break
		}
	}
	s.DeviceId = d.ID
	s.FlowId = f.ID
	s.Transport = SenderTransport

	switch t {
	case "video":
		label := getResourceLabel("TestReceiverVideo", index)
		s.BaseResource = SetBaseResourceProperties(label, "NMOS Test Video Receiver")
	case "audio":
		label := getResourceLabel("TestReceiverAudio", index)
		s.BaseResource = SetBaseResourceProperties(label, "NMOS Test Audio Receiver")
	case "data":
		label := getResourceLabel("TestReceiverData", index)
		s.BaseResource = SetBaseResourceProperties(label, "NMOS Test Data Receiver")
	default:
		log.Fatal("No valid sender type provided")
	}
}
