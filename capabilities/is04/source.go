package is04

type GrainRate struct {
	Numerator   int `json:"numerator"`
	Denominator int `json:"denominator"`
}

type SourceCore struct {
	ResourceCore
	GrainRate GrainRate    `json:"grain_rate"`
	Caps      Capabilities `json:"caps"`
	DeviceId  string       `json:"device_id"`
	Parents   []string     `json:"parents"`
	ClockName string       `json:"clock_name"`
	Device    Device       `json:"-"`
}

type SourceChannels struct {
	Label  string `json:"label"`
	Symbol string `json:"symbol"`
}

type SourceAudio struct {
	SourceCore
	Channels []SourceChannels `json:"channels"`
	Format   string           `json:"format"`
}

type SourceData struct {
	SourceCore
	EventType string `json:"event_type"`
	Format    string `json:"format"`
}

type SourceGeneric struct {
	SourceCore
	Format string `json:"format"`
}

func NewSource(
	d Device,
	format string
) {

}
