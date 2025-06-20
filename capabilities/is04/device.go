package is04

import "github.com/segmentio/encoding/json"

type Control struct {
	Href          string `json:"href"`
	Type          string `json:"type"`
	Authorization bool   `json:"authorization"`
}

type Device struct {
	*ResourceCore
	Type   string `json:"type"`
	NodeId string `json:"node_id"`
	// The inclusion of the senders and receivers lists has been deprecated
	Senders   []string   `json:"senders"`
	Receivers []string   `json:"receivers"`
	Controls  []*Control `json:"controls"`
	Sources   []Source   `json:"-"`
}

func (d Device) Encode() ([]byte, error) {
	e, err := json.Marshal(d)

	if err != nil {
		return nil, err
	}

	return e, nil
}

func (d *Device) SetControls(c *Control) {
	d.Controls = append(d.Controls, c)
}

func NewDevice(
	n *Node,
	resourceCore *ResourceCore,
) *Device {
	d := &Device{
		ResourceCore: resourceCore,
		Type:         "urn:x-nmos:device:generic",
		NodeId:       n.ID,
		Senders:      []string{},
		Receivers:    []string{},
		Controls:     []*Control{},
	}
	n.Devices = append(n.Devices, d)
	return d
}
