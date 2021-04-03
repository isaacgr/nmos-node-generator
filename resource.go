package main

import (
	"io/ioutil"
	"log"
	"reflect"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/helloeave/json"
	regen "github.com/zach-klippenstein/goregen"
)

func SetBaseResourceProperties(label string, description string) *BaseResource {
	r := BaseResource{}
	r.ID = uuid.NewString()
	r.Version = GenerateVersion()
	r.Label = label
	r.Description = description
	return &r
}

func SetBaseSourceProperties(label string, description string, d Device) *BaseSource {
	r := SetBaseResourceProperties(label, description)
	bs := BaseSource{}
	bs.BaseResource = r
	bs.GrainRate = GrainRate{
		1,
		2,
	}
	bs.DeviceId = d.ID
	bs.ClockName = "clk0"
	return &bs
}

func SetBaseFlowProperties(label string, description string, d Device, s BaseSource) *BaseFlow {
	r := SetBaseResourceProperties(label, description)
	bf := BaseFlow{}
	bf.BaseResource = r
	bf.GrainRate = GrainRate{
		1,
		2,
	}
	bf.DeviceId = d.ID
	bf.SourceID = s.ID
	return &bf
}

func WriteResourceToFile(t interface{}) error {
	r := reflect.ValueOf(t)
	data, err := json.MarshalSafeCollections(t)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(r.FieldByName("Label").String()+".json", data, 0644)
}

func getResourceLabel(label string, index int) string {
	label = label + strconv.Itoa(index)
	return label
}

func GenerateVersion() string {
	return strconv.FormatInt(time.Now().UTC().Unix(), 10) + ":" + strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
}

func GenerateMac() string {
	mac, err := regen.Generate("^([0-9a-f]{2}-){5}([0-9a-f]{2})$")
	if err != nil {
		log.Fatal("Unable to generate mac address")
	}
	return mac
	// buf := make([]byte, 6)
	// var mac net.HardwareAddr

	// _, err := rand.Read(buf)
	// if err != nil {
	// 	log.Fatal("Unable to generate mac address")
	// 	os.Exit(1)
	// }

	// // Set the local bit
	// buf[0] |= 2

	// mac = append(mac, buf[0], buf[1], buf[2], buf[3], buf[4], buf[5])

	// return mac.String()
}
