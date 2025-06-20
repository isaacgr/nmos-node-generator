package is04

type SampleRate struct {
	Numerator   int `json:"numerator"`
	Denominator int `json:"denominator"`
}

type FlowCore struct {
	*ResourceCore
	SourceID  string        `json:"source_id"`
	DeviceID  string        `json:"device_id"`
	Parents   []string      `json:"parents"`
	GrainRate GrainRate     `json:"grain_rate"`
	Source    *IS04Resource `json:"-"`
	Device    *Device       `json:"-"`
}

type FlowVideo struct {
	*FlowCore
	Format                 string `json:"format"`
	FrameWidth             int    `json:"frame_width"`
	FrameHeight            int    `json:"frame_height"`
	InterlaceMode          string `json:"interlace_mode"`
	Colorspace             string `json:"colorspace"`
	TransferCharacteristic string `json:"transfer_characteristic"`
}

type FlowAudio struct {
	*FlowCore
	Format     string     `json:"format"`
	SampleRate SampleRate `json:"sample_rate"`
}

type RawVideoComponent struct {
	Name     string `json:"name"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	BitDepth int    `json:"bit_depth"`
}

type FlowVideoRaw struct {
	FlowVideo
	MediaType  string              `json:"media_type"`
	Components []RawVideoComponent `json:"components"`
}

type FlowAudioRaw struct {
	FlowAudio
	MediaType string `json:"media_type"`
	BitDepth  int    `json:"bit_depth"`
}

type FlowAudioCoded struct {
	FlowAudio
	MediaType string `json:"media_type"`
}

type FlowData struct {
	*FlowCore
	Format    string `json:"format"`
	MediaType string `json:"media_type"`
}

type FlowJsonData struct {
	*FlowCore
	Format    string `json:"format"`
	MediaType string `json:"media_type"`
	EventType string `json:"event_type"`
}

type DidSdid struct {
	DID  string `json:"DID"`
	SDID string `json:"SDID"`
}

type FlowSdiAncData struct {
	*FlowCore
	Format    string    `json:"format"`
	MediaType string    `json:"media_type"`
	DidSdid   []DidSdid `json:"DID_SDID"`
}

type FlowMux struct {
	*FlowCore
	Format    string `json:"format"`
	MediaType string `json:"media_type"`
}
