package node

// Common resource

type Tags struct{}

type BaseResource struct {
	ID          string `json:"id"`
	Version     string `json:"version"`
	Label       string `json:"label"`
	Description string `json:"description"`
	Tags        Tags   `json:"tags"`
}

// Node

type Endpoint struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Protocol string `json:"protocol"`
}

type Api struct {
	Versions  []string   `json:"versions"`
	Endpoints []Endpoint `json:"endpoints"`
}

type NetworkDevice struct {
	ChassisId string `json:"chassis_id"`
	PortId    string `json:"port_id"`
}

type NetworkInterface struct {
	ChassisId             string        `json:"chassis_id"`
	PortId                string        `json:"port_id"`
	Name                  string        `json:"name"`
	AttachedNetworkDevice NetworkDevice `json:"attached_network_device"`
}

type ClockInternal struct {
	Name    string `json:"name"`
	RefType string `json:"ref_type"`
}

type ClockPTP struct {
	Name      string `json:"name"`
	RefType   string `json:"ref_type"`
	Traceable bool   `json:"traceable"`
	Version   string `json:"version"`
	Gmid      string `json:"gmid"`
	Locked    bool   `json:"locked"`
}

type Service struct {
	Href          string `json:"href"`
	Type          string `json:"type"`
	Authorization bool   `json:"authorization"`
}

type Capabilities struct{}

type Node struct {
	*BaseResource
	Href       string             `json:"href"`
	Hostname   string             `json:"hostname"`
	Caps       Capabilities       `json:"caps"`
	Api        Api                `json:"api"`
	Services   []Service          `json:"services"`
	Clocks     []interface{}      `json:"clocks"`
	Interfaces []NetworkInterface `json:"interfaces"`
}

// Device

type Controls struct {
	Href          string `json:"href"`
	Type          string `json:"type"`
	Authorization bool   `json:"authorization"`
}

type Device struct {
	*BaseResource
	Type      string     `json:"type"`
	NodeId    string     `json:"node_id"`
	Senders   []string   `json:"senders"`
	Receivers []string   `json:"receivers"`
	Controls  []Controls `json:"controls"`
}

// Sender

type SenderSubscription struct {
	ReceiverId *string `json:"receiver_id"`
	Active     bool    `json:"active"`
}

type BaseSender struct {
	*BaseResource
	FlowId            string             `json:"flow_id"`
	Caps              Capabilities       `json:"caps"`
	Transport         string             `json:"transport"`
	DeviceId          string             `json:"device_id"`
	Manifest          string             `json:"manifest"`
	InterfaceBindings []string           `json:"interface_bindings"`
	Subscription      SenderSubscription `json:"subscription"`
}

type SenderVideo struct {
	*BaseSender
}

type SenderAudio struct {
	*BaseSender
}

type SenderData struct {
	*BaseSender
}

// Receiver

type ReceiverSubscription struct {
	SenderId *string `json:"sender_id"`
	Active   bool    `json:"active"`
}

type ReceiverCaps struct {
	MediaTypes []string `json:"media_types"`
}

type BaseReceiver struct {
	*BaseResource
	Transport         string               `json:"transport"`
	DeviceId          string               `json:"device_id"`
	InterfaceBindings []string             `json:"interface_bindings"`
	Subscription      ReceiverSubscription `json:"subscription"`
	Format            string               `json:"format"`
	Caps              ReceiverCaps         `json:"caps"`
}

type ReceiverVideo struct {
	*BaseReceiver
}

type ReceiverAudio struct {
	*BaseReceiver
}

type ReceiverData struct {
	*BaseReceiver
}

// Source

type GrainRate struct {
	Numerator   int `json:"numerator"`
	Denominator int `json:"denominator"`
}

type BaseSource struct {
	*BaseResource
	GrainRate GrainRate    `json:"grain_rate"`
	Caps      Capabilities `json:"caps"`
	DeviceId  string       `json:"device_id"`
	Parents   []string     `json:"parents"`
	ClockName string       `json:"clock_name"`
}

type SourceChannels struct {
	Label  string `json:"label"`
	Symbol string `json:"symbol"`
}

type SourceAudio struct {
	*BaseSource
	Channels []SourceChannels `json:"channels"`
	Format   string           `json:"format"`
}

type SourceData struct {
	*BaseSource
	EventType string `json:"event_type"`
	Format    string `json:"format"`
}

type SourceGeneric struct {
	*BaseSource
	Format string `json:"format"`
}

// Flow

type SampleRate struct {
	Numerator   int `json:"numerator"`
	Denominator int `json:"denominator"`
}

type BaseFlow struct {
	*BaseResource
	SourceID  string    `json:"source_id"`
	DeviceId  string    `json:"device_id"`
	Parents   []string  `json:"parents"`
	GrainRate GrainRate `json:"grain_rate"`
}

type FlowVideo struct {
	*BaseFlow
	Format                 string `json:"format"`
	FrameWidth             int    `json:"frame_width"`
	FrameHeight            int    `json:"frame_height"`
	InterlaceMode          string `json:"interlace_mode"`
	Colorspace             string `json:"colorspace"`
	TransferCharacteristic string `json:"transfer_characteristic"`
}

type FlowAudio struct {
	*BaseFlow
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
	*BaseFlow
	Format    string `json:"format"`
	MediaType string `json:"media_type"`
}

type FlowJsonData struct {
	*BaseFlow
	Format    string `json:"format"`
	MediaType string `json:"media_type"`
	EventType string `json:"event_type"`
}

type DidSdid struct {
	DID  string `json:"DID"`
	SDID string `json:"SDID"`
}

type FlowSdiAncData struct {
	*BaseFlow
	Format    string    `json:"format"`
	MediaType string    `json:"media_type"`
	DidSdid   []DidSdid `json:"DID_SDID"`
}

type FlowMux struct {
	*BaseFlow
	Format    string `json:"format"`
	MediaType string `json:"media_type"`
}
