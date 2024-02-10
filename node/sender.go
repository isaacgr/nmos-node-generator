package node

import "github.com/isaacgr/nmos-node-generator/config"

type Sender interface {
	BuildResource(n Node, d *Device, f Flow, index int, useRandomResource bool)
}

func getSenderConfig() config.SourceResource {
	return config.New().ResourceQuantities.Sources
}

func BuildBaseSender(n Node, d *Device, f Flow, b []int) *BaseSender {
	s := BaseSender{}
	for _, iface := range b {
		s.InterfaceBindings = append(s.InterfaceBindings, n.Interfaces[iface-1].Name)
	}
	s.DeviceId = d.ID
	s.FlowId = f.getId()
	s.Transport = SenderTransport
	s.Manifest = &n.Href
	return &s
}

func (s *SenderVideo) BuildResource(n Node, d *Device, f Flow, index int, useRandomResource bool) {
	s.BaseSender = BuildBaseSender(n, d, f, getSenderConfig().Generic.Flows.Sender.Iface)
	label := getResourceLabel(d.Label+"."+"SenderVideo", index)
	s.BaseResource = SetBaseResourceProperties(label, "NMOS Test Video Sender", useRandomResource)
	d.Senders = append(d.Senders, s.ID)
}

func (s *SenderAudio) BuildResource(n Node, d *Device, f Flow, index int, useRandomResource bool) {
	s.BaseSender = BuildBaseSender(n, d, f, getSenderConfig().Audio.Flows.Sender.Iface)
	label := getResourceLabel(d.Label+"."+"SenderAudio", index)
	s.BaseResource = SetBaseResourceProperties(label, "NMOS Test Audio Sender", useRandomResource)
	d.Senders = append(d.Senders, s.ID)
}
func (s *SenderData) BuildResource(n Node, d *Device, f Flow, index int, useRandomResource bool) {
	s.BaseSender = BuildBaseSender(n, d, f, getSenderConfig().Data.Flows.Sender.Iface)
	label := getResourceLabel(d.Label+"."+"SenderData", index)
	s.BaseResource = SetBaseResourceProperties(label, "NMOS Test Data Sender", useRandomResource)
	d.Senders = append(d.Senders, s.ID)
}
