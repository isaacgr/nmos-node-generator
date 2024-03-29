package util

import (
	"log"

	"github.com/isaacgr/nmos-node-generator/config"
	"github.com/isaacgr/nmos-node-generator/node"
)

func BuildNodes(nn int, ni int, nameprefix string, attachedNetworkDevices []config.AttachedNetworkDevices, randomNodeUUID bool) []node.Node {
	nodes := []node.Node{}
	for i := 0; i < nn; i++ {
		node := node.Node{}
		node.BuildResource(i+1, ni, nameprefix, attachedNetworkDevices, randomNodeUUID)
		nodes = append(nodes, node)
	}
	return nodes
}

func BuildDevices(n []node.Node, nd int, nameprefix string, deviceip string, devicePortStart int, useRandom bool) []node.Device {
	devices := []node.Device{}
	for j := 0; j < len(n); j++ {
		for i := 0; i < nd; i++ {
			device := node.Device{}
			device.BuildResource(n[j], nd, j+1, i+1, nameprefix, deviceip, devicePortStart, useRandom)
			devices = append(devices, device)
		}
	}
	return devices
}

func BuildReceivers(n []node.Node, d []node.Device, nvr int, nar int, ndr int, useRandomResource bool) []node.Receiver {
	receivers := []node.Receiver{}
	for k := 0; k < len(n); k++ {
		for j := (len(d) / len(n)) * k; j < (len(d)/len(n))*(k+1); j++ {
			for i := 0; i < nvr; i++ {
				receiver := node.ReceiverVideo{}
				receiver.BuildResource(n[k], &d[j], i+1, useRandomResource)
				receivers = append(receivers, &receiver)
			}
			for i := 0; i < nar; i++ {
				receiver := node.ReceiverAudio{}
				receiver.BuildResource(n[k], &d[j], i+1, useRandomResource)
				receivers = append(receivers, &receiver)
			}
			for i := 0; i < ndr; i++ {
				receiver := node.ReceiverData{}
				receiver.BuildResource(n[k], &d[j], i+1, useRandomResource)
				receivers = append(receivers, &receiver)
			}
		}
	}
	return receivers
}

func BuildSenders(n []node.Node, d []node.Device, f []node.Flow, useRandomResource bool) []node.Sender {
	senders := []node.Sender{}
	for k := 0; k < len(n); k++ {
		for j := (len(d) / len(n)) * k; j < (len(d)/len(n))*(k+1); j++ {
			for x := (len(f) / len(d)) * j; x < (len(f)/len(d))*(j+1); x++ {
				switch f[x].(type) {
				case *node.FlowVideoRaw:
					sender := node.SenderVideo{}
					sender.BuildResource(n[k], &d[j], f[x], x+1, useRandomResource)
					senders = append(senders, &sender)
				case *node.FlowMux:
					sender := node.SenderVideo{}
					sender.BuildResource(n[k], &d[j], f[x], x+1, useRandomResource)
					senders = append(senders, &sender)
				case *node.FlowAudioRaw:
					sender := node.SenderAudio{}
					sender.BuildResource(n[k], &d[j], f[x], x+1, useRandomResource)
					senders = append(senders, &sender)
				default:
					sender := node.SenderData{}
					sender.BuildResource(n[k], &d[j], f[x], x+1, useRandomResource)
					senders = append(senders, &sender)
				}
			}
		}
	}
	return senders
}

func BuildSources(d []node.Device, ngs int, nas int, nds int, useRandomResource bool) []node.Source {
	sources := []node.Source{}
	for j := 0; j < len(d); j++ {
		for i := 0; i < ngs; i++ {
			source := node.SourceGeneric{}
			source.BuildResource(d[j], i+1, useRandomResource)
			sources = append(sources, &source)
		}
		for i := 0; i < nas; i++ {
			source := node.SourceAudio{}
			source.BuildResource(d[j], i+1, useRandomResource)
			sources = append(sources, &source)
		}
		for i := 0; i < nds; i++ {
			source := node.SourceData{}
			source.BuildResource(d[j], i+1, useRandomResource)
			sources = append(sources, &source)
		}
	}
	return sources
}

func BuildFlows(d []node.Device, s []node.Source, vf string, af string, df string, useRandomResource bool) []node.Flow {
	flows := []node.Flow{}
	for j := 0; j < len(d); j++ {
		for i := (len(s) / len(d)) * j; i < (len(s)/len(d))*(j+1); i++ {
			switch s[i].(type) {
			case *node.SourceGeneric:
				if vf == "raw" {
					f := node.FlowVideoRaw{}
					f.BuildResource(d[j], s[i], i+1, useRandomResource)
					flows = append(flows, &f)
				} else if vf == "mux" {
					f := node.FlowMux{}
					f.BuildResource(d[j], s[i], i+1, useRandomResource)
					flows = append(flows, &f)
				} else {
					log.Fatal("Invalid video flow type")
				}
			case *node.SourceAudio:
				f := node.FlowAudioRaw{}
				f.BuildResource(d[j], s[i], i+1, useRandomResource)
				f.MediaType = node.AudioMediaTypes[af]
				flows = append(flows, &f)
			case *node.SourceData:
				if df == "smpte291" {
					f := node.FlowSdiAncData{}
					f.BuildResource(d[j], s[i], i+1, useRandomResource)
					flows = append(flows, &f)
				} else if df == "json" {
					f := node.FlowJsonData{}
					f.BuildResource(d[j], s[i], i+1, useRandomResource)
					flows = append(flows, &f)
				} else {
					log.Fatal("Invalid data flow type")
				}
			}
		}
	}
	return flows
}
