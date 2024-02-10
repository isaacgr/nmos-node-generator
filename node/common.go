package node

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

func SetBaseResourceProperties(label string, description string, random bool) *BaseResource {
	r := BaseResource{}
	if !random {
		namespace := uuid.MustParse("00000000-0000-0000-0000-000000000000")
		u := uuid.NewMD5(namespace, []byte(label+description))
		r.ID = u.String()
	} else {
		u, _ := uuid.NewRandom()
		r.ID = u.String()
	}
	r.Version = GenerateVersion()
	r.Label = label
	r.Description = description
	return &r
}

func SetBaseSourceProperties(label string, description string, d Device, useRandomResource bool) *BaseSource {
	r := SetBaseResourceProperties(label, description, useRandomResource)
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

func SetBaseFlowProperties(label string, description string, d Device, sId string, useRandomResource bool) *BaseFlow {
	r := SetBaseResourceProperties(label, description, useRandomResource)
	bf := BaseFlow{}
	bf.BaseResource = r
	bf.GrainRate = GrainRate{
		1,
		2,
	}
	bf.DeviceId = d.ID
	bf.SourceID = sId
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
}
