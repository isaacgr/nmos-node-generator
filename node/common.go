package node

import (
	"io/ioutil"
	"log"
	"math/rand"
	"reflect"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/helloeave/json"
	regen "github.com/zach-klippenstein/goregen"
)

func SetBaseResourceProperties(label string, description string) *BaseResource {
	r := BaseResource{}
	rnd := rand.New(rand.NewSource(1))
	uuid.SetRand(rnd)
	u, _ := uuid.NewRandomFromReader(rnd)
	r.ID = u.String()
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

func SetBaseFlowProperties(label string, description string, d Device, sId string) *BaseFlow {
	r := SetBaseResourceProperties(label, description)
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
