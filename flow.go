package main

type Flow interface {
	BuildResource(d Device, s Source, index int)
}

func (f *FlowVideo) BuildResource(d Device, s SourceGeneric, index int) {
	label := getResourceLabel("TestFlowVideo", index)
	f.BaseFlow = SetBaseFlowProperties(label, "NMOS Video Flow", d, *s.BaseSource)
	f.Format = VideoFormat
	f.FrameWidth = 1920
	f.FrameHeight = 1080
	f.InterlaceMode = InterlaceModes["progressive"]
	f.Colorspace = ColorSpaces["BT709"]
	f.TransferCharacteristic = TransferCharacteristics["SDR"]
}

func (f *FlowVideoRaw) BuildResource(d Device, s SourceGeneric, index int) {
	fv := FlowVideo{}
	fv.BuildResource(d, s, index)
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

func (f *FlowAudio) BuildResource(d Device, s SourceAudio, index int) {
	label := getResourceLabel("TestFlowAudio", index)
	f.BaseFlow = SetBaseFlowProperties(label, "NMOS Audio Flow", d, *s.BaseSource)
	f.Format = AudioFormat
	f.SampleRate = SampleRate{
		1,
		1,
	}
}

func (f *FlowAudioRaw) BuildResource(d Device, s SourceAudio, index int) {
	fa := FlowAudio{}
	fa.BuildResource(d, s, index)
	f.FlowAudio = fa
	f.MediaType = AudioMediaTypes[16]
	f.BitDepth = 16
}

func (f *FlowAudioCoded) BuildResource(d Device, s SourceAudio, index int) {
	fa := FlowAudio{}
	fa.BuildResource(d, s, index)
	f.FlowAudio = fa
	f.MediaType = AudioMediaTypes[16]
}

func (f *FlowData) BuildResource(d Device, s SourceData, index int) {
	label := getResourceLabel("TestFlowData", index)
	f.BaseFlow = SetBaseFlowProperties(label, "NMOS Data Flow", d, *s.BaseSource)
	f.Format = DataFormat
	f.MediaType = DataMediaTypes["json"]
}

func (f *FlowJsonData) BuildResource(d Device, s SourceData, index int) {
	label := getResourceLabel("TestFlowJsonData", index)
	f.BaseFlow = SetBaseFlowProperties(label, "NMOS JSON Data Flow", d, *s.BaseSource)
	f.MediaType = DataMediaTypes["json"]
	f.EventType = "boolean"
}

func (f *FlowSdiAncData) BuildResource(d Device, s SourceData, index int) {
	label := getResourceLabel("TestFlowSdiAncData", index)
	f.BaseFlow = SetBaseFlowProperties(label, "NMOS SDI Anc Data Flow", d, *s.BaseSource)
	f.MediaType = DataMediaTypes["smpte291"]
}

func (f *FlowMux) BuildResource(d Device, s SourceGeneric, index int) {
	label := getResourceLabel("TestFlowMuxData", index)
	f.BaseFlow = SetBaseFlowProperties(label, "NMOS Mux Data Flow", d, *s.BaseSource)
	f.Format = MuxFormat
	f.MediaType = MuxMediaTypes["2022-6"]
}
