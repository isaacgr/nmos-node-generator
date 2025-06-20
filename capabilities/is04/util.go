package is04

import (
	regen "github.com/zach-klippenstein/goregen"
	"log"
	"strconv"
	"time"
)

func ResourceVersion() string {
	return strconv.FormatInt(
		time.Now().UTC().Unix(),
		10,
	) + ":" + strconv.FormatInt(
		time.Now().UTC().UnixNano(),
		10,
	)
}

func GenerateMac() string {
	mac, err := regen.Generate("^([0-9a-f]{2}-){5}([0-9a-f]{2})$")
	if err != nil {
		log.Fatalf("Unable to generate mac address. Error [%s]", err)
	}
	return mac
}
