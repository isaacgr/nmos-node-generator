package util

import "github.com/isaacgr/nmos-node-generator/node"

func BuildNodes(nn int) []node.Node {
	nodes := []node.Node{}
	for i := 0; i < nn; i++ {
		node := node.Node{}
		node.BuildResource(i + 1)
		nodes = append(nodes, node)
	}
	return nodes
}

func BuildDevices(n []node.Node, nd int) []node.Device {
	devices := []node.Device{}
	for i := 0; i < nd; i++ {
		for j := 0; j < len(n); j++ {
			device := node.Device{}
			device.BuildResource(n[j], i+1)
			devices = append(devices, device)
		}
	}
	return devices
}

func BuildReceivers(n []node.Node, d []node.Device, nvr int, nar int, ndr int) []node.Receiver {
	receivers := []node.Receiver{}
	for j := 0; j < len(d); j++ {
		for i := 0; i < nvr; i++ {
			receiver := node.ReceiverVideo{}
			receiver.BuildResource(n[i], &d[i], i+1)
			receivers = append(receivers, &receiver)
		}
		for i := 0; i < nar; i++ {
			receiver := node.ReceiverAudio{}
			receiver.BuildResource(n[i], &d[i], i+1)
			receivers = append(receivers, &receiver)
		}
		for i := 0; i < ndr; i++ {
			receiver := node.ReceiverData{}
			receiver.BuildResource(n[i], &d[i], i+1)
			receivers = append(receivers, &receiver)
		}
	}
	return receivers
}

func BuildSources(d []node.Device, ngs int, nas int, nds int) []node.Source {
	sources := []node.Source{}
	for j := 0; j < len(d); j++ {
		for i := 0; i < ngs; i++ {
			source := node.SourceGeneric{}
			source.BuildResource(d[i], i+1)
			sources = append(sources, &source)
		}
		for i := 0; i < nas; i++ {
			source := node.SourceAudio{}
			source.BuildResource(d[i], i+1)
			sources = append(sources, &source)
		}
		for i := 0; i < nds; i++ {
			source := node.SourceData{}
			source.BuildResource(d[i], i+1)
			sources = append(sources, &source)
		}
	}
	return sources
}
