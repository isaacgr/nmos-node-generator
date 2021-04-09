package util

import (
	"log"

	"github.com/isaacgr/nmos-node-generator/node"
)

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

func BuildFlows(d []node.Device, s []node.Source, vf string, af string, df string) []node.Flow {
	flows := []node.Flow{}
	for j := 0; j < len(d); j++ {
		for i := 0; i < len(s); i++ {
			switch s[i].(type) {
			case *node.SourceGeneric:
				if vf == "raw" {
					f := node.FlowVideoRaw{}
					f.BuildResource(d[i], s[i], i+1)
					flows = append(flows, &f)
				} else if vf == "mux" {
					f := node.FlowMux{}
					f.BuildResource(d[i], s[i], i+1)
					flows = append(flows, &f)
				} else {
					log.Fatal("Invalid video flow type")
				}
			case *node.SourceAudio:
				f := node.FlowAudioRaw{}
				f.BuildResource(d[i], s[i], i+1)
				f.MediaType = node.AudioMediaTypes[af]
				flows = append(flows, &f)
			case *node.SourceData:
				if df == "smpte291" {
					f := node.FlowSdiAncData{}
					f.BuildResource(d[i], s[i], i+1)
					flows = append(flows, &f)
				} else if df == "json" {
					f := node.FlowJsonData{}
					f.BuildResource(d[i], s[i], i+1)
					flows = append(flows, &f)
				} else {
					log.Fatal("Invalid data flow type")
				}
			}
		}
	}
	return flows
}
