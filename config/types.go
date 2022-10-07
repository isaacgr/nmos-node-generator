package config

type Config struct {
	ResourceQuantities ResourceQuantities `json:"resource"`
	Registry           Registry           `json:"registry"`
}

type FlowResource struct {
	MediaType string         `json:"media_type"`
	Sender    SenderResource `json:"sender"`
}

type GenericSource struct {
	Count int          `json:"count"`
	Flows FlowResource `json:"flows"`
}

type AudioSource struct {
	Count int          `json:"count"`
	Flows FlowResource `json:"flows"`
}

type DataSource struct {
	Count int          `json:"count"`
	Flows FlowResource `json:"flows"`
}

type SourceResource struct {
	Generic GenericSource `json:"generic"`
	Audio   AudioSource   `json:"audio"`
	Data    DataSource    `json:"data"`
}

type ReceiverDetails struct {
	Count     int    `json:"count"`
	MediaType string `json:"media_type"`
	Iface     []int  `json:"iface"`
}

type ReceiverResource struct {
	Video ReceiverDetails `json:"video"`
	Audio ReceiverDetails `json:"audio"`
	Data  ReceiverDetails `json:"data"`
}

type SenderResource struct {
	Iface []int `json:"iface"`
}

type NodeResource struct {
	Count         int    `json:"count"`
	NumInterfaces int    `json:"num_interfaces"`
	NamePrefix    string `json:"name_prefix"`
}

type ResourceQuantities struct {
	Nodes      NodeResource     `json:"nodes"`
	Devices    int              `json:"devices"`
	Receivers  ReceiverResource `json:"receivers"`
	Sources    SourceResource   `json:"sources"`
	NamePrefix string           `json:"name_prefix"`
}

type Registry struct {
	IP      string `json:"ip"`
	Port    int    `json:"port"`
	Scheme  string `json:"scheme"`
	Version string `json:"version"`
}
