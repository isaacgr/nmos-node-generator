package node

import regen "github.com/zach-klippenstein/goregen"

type Flow interface {
	BuildResource(d Device, s Source, index int, useRandomResource bool)
	getId() string
}

func (f *FlowVideoRaw) BuildResource(d Device, s Source, index int, useRandomResource bool) {
	fv := FlowVideo{}
	label := getResourceLabel(d.Label+"."+"FlowVideo", index)
	f.BaseFlow = SetBaseFlowProperties(label, "NMOS Video Flow", d, s.getId(), useRandomResource)
	f.Format = VideoFormat
	f.FrameWidth = 1920
	f.FrameHeight = 1080
	f.InterlaceMode = InterlaceModes["progressive"]
	f.Colorspace = ColorSpaces["BT709"]
	f.TransferCharacteristic = TransferCharacteristics["SDR"]
	f.FlowVideo = fv
	f.MediaType = VideoMediaTypes["raw"]
	c := RawVideoComponent{
		RawVideoCompName["Y"],
		1920,
		1080,
		10,
	}
	f.Components = append(f.Components, c)
}

func (f *FlowAudioRaw) BuildResource(d Device, s Source, index int, useRandomResource bool) {
	fa := FlowAudio{}
	label := getResourceLabel(d.Label+"."+"FlowAudio", index)
	f.BaseFlow = SetBaseFlowProperties(label, "NMOS Audio Flow", d, s.getId(), useRandomResource)
	f.Format = AudioFormat
	f.SampleRate = SampleRate{
		1,
		1,
	}
	f.FlowAudio = fa
	f.MediaType = AudioMediaTypes["audio/L16"]
	f.BitDepth = 16
}

func (f *FlowAudioCoded) BuildResource(d Device, s Source, index int, useRandomResource bool) {
	fa := FlowAudio{}
	label := getResourceLabel(d.Label+"."+"FlowAudio", index)
	f.BaseFlow = SetBaseFlowProperties(label, "NMOS Audio Flow", d, s.getId(), useRandomResource)
	f.Format = AudioFormat
	f.SampleRate = SampleRate{
		1,
		1,
	}
	f.FlowAudio = fa
	f.MediaType, _ = regen.Generate("^audio\\/[^\\s\\/]+$")
}

func (f *FlowData) BuildResource(d Device, s Source, index int, useRandomResource bool) {
	label := getResourceLabel(d.Label+"."+"FlowData", index)
	f.BaseFlow = SetBaseFlowProperties(label, "NMOS Data Flow", d, s.getId(), useRandomResource)
	f.Format = DataFormat
	f.MediaType, _ = regen.Generate("^[^\\s\\/]+\\/[^\\s\\/]+$")
}

func (f *FlowJsonData) BuildResource(d Device, s Source, index int, useRandomResource bool) {
	label := getResourceLabel(d.Label+"."+"FlowJsonData", index)
	f.BaseFlow = SetBaseFlowProperties(label, "NMOS JSON Data Flow", d, s.getId(), useRandomResource)
	f.MediaType = DataMediaTypes["json"]
	f.EventType = "boolean"
	f.Format = DataFormat
}

func (f *FlowSdiAncData) BuildResource(d Device, s Source, index int, useRandomResource bool) {
	label := getResourceLabel(d.Label+"."+"FlowSdiAncData", index)
	f.BaseFlow = SetBaseFlowProperties(label, "NMOS SDI Anc Data Flow", d, s.getId(), useRandomResource)
	f.MediaType = DataMediaTypes["smpte291"]
	f.Format = DataFormat
}

func (f *FlowMux) BuildResource(d Device, s Source, index int, useRandomResource bool) {
	label := getResourceLabel(d.Label+"."+"FlowMuxData", index)
	f.BaseFlow = SetBaseFlowProperties(label, "NMOS Mux Data Flow", d, s.getId(), useRandomResource)
	f.Format = MuxFormat
	f.MediaType = MuxMediaTypes["2022-6"]
}

func (f *FlowVideoRaw) getId() string {
	return f.ID
}

func (f *FlowMux) getId() string {
	return f.ID
}

func (f *FlowAudioRaw) getId() string {
	return f.ID
}

func (f *FlowAudioCoded) getId() string {
	return f.ID
}

func (f *FlowData) getId() string {
	return f.ID
}

func (f *FlowJsonData) getId() string {
	return f.ID
}

func (f *FlowSdiAncData) getId() string {
	return f.ID
}
