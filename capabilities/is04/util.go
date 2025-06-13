package is04

import (
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
