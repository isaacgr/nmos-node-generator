package is04

import (
	"github.com/isaacgr/nmos-node-generator/constants"
	"github.com/segmentio/encoding/json"
	"log"
)

type GrainRate struct {
	Numerator   int `json:"numerator"`
	Denominator int `json:"denominator"`
}

type SourceCore struct {
	*ResourceCore
	GrainRate GrainRate `json:"grain_rate"`
	Caps      struct{}  `json:"caps"`
	DeviceId  string    `json:"device_id"`
	Parents   []string  `json:"parents"`
	ClockName string    `json:"clock_name"`
	Device    Device    `json:"-"`
}

type Channels struct {
	Label  string `json:"label"`
	Symbol string `json:"symbol"`
}

type SourceAudio struct {
	*SourceCore
	Channels []Channels `json:"channels"`
	Format   string     `json:"format"`
}

type SourceData struct {
	*SourceCore
	EventType string `json:"event_type"`
	Format    string `json:"format"`
}

type SourceGeneric struct {
	*SourceCore
	Format string `json:"format"`
}

type Source interface {
	IS04Resource
	GetFormat() string
}

func (sg SourceGeneric) GetFormat() string {
	return sg.Format
}

func (sa SourceAudio) GetFormat() string {
	return sa.Format
}

func (sd SourceData) GetFormat() string {
	return sd.Format
}

func (sg SourceGeneric) Encode() ([]byte, error) {
	e, err := json.Marshal(sg)

	if err != nil {
		return nil, err
	}

	return e, nil
}

func (sa SourceAudio) Encode() ([]byte, error) {
	e, err := json.Marshal(sa)

	if err != nil {
		return nil, err
	}

	return e, nil
}

func (sd SourceData) Encode() ([]byte, error) {
	e, err := json.Marshal(sd)

	if err != nil {
		return nil, err
	}

	return e, nil
}

func newSourceCore(
	d *Device,
	resourceCore *ResourceCore,
) *SourceCore {
	gr := GrainRate{
		Numerator:   1,
		Denominator: 1,
	}
	return &SourceCore{
		ResourceCore: resourceCore,
		GrainRate:    gr,
		ClockName:    "clk0",
		DeviceId:     d.ID,
	}

}

func NewSource(
	d *Device,
	format string,
	resourceCore *ResourceCore,
) Source {
	sourceCore := newSourceCore(
		d,
		resourceCore,
	)

	var source Source

	if format == constants.VideoFormat || format == constants.MuxFormat {
		source = SourceGeneric{
			SourceCore: sourceCore,
			Format:     format,
		}
	} else if format == constants.AudioFormat {
		source = SourceAudio{
			SourceCore: sourceCore,
			Format:     format,
		}
	} else if format == constants.DataFormat {
		source = SourceData{
			SourceCore: sourceCore,
			Format:     format,
		}
	} else {
		log.Fatalf("Invalid source format. Format [%s]", format)
		return nil

	}
	d.Sources = append(d.Sources, source)
	return source
}
