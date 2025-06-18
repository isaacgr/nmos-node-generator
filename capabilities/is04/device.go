package is04

import "encoding/json"

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
	Senders   []string  `json:"senders"`
	Receivers []string  `json:"receivers"`
	Controls  []Control `json:"controls"`
}

func (d Device) encode() ([]byte, error) {
	e, err := json.Marshal(d)

	if err != nil {
		return nil, err
	}

	return e, nil
}

func (d *Device) SetControls(c Control) {
	d.Controls = append(d.Controls, c)
}

func NewDevice(
	n *Node,
	randomId bool,
) *Device {
	d := Device{
		ResourceCore: NewResourceCore(
			"evDevice",
			"NMOS Test Device",
			randomId,
		),
		Type:   "urn:x-nmos:device:generic",
		NodeId: n.ID,
		Senders: []string{},
		Receivers: []string{},
		Controls: []Control{},
	}
	n.devices = append(n.devices, d)
	return &d
}
