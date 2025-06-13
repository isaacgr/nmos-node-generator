package is04

import (
	"github.com/google/uuid"
	"github.com/isaacgr/nmos-node-generator/capabilities/is04"
	"github.com/isaacgr/nmos-node-generator/types"
)

func ResourceCore(
	label string,
	description string,
	random bool,
) *types.ResourceCore {
	r := types.ResourceCore{}
	if !random {
		namespace := uuid.MustParse("00000000-0000-0000-0000-000000000000")
		u := uuid.NewMD5(namespace, []byte(label+description))
		r.ID = u.String()
	} else {
		u, _ := uuid.NewRandom()
		r.ID = u.String()
	}
	r.Version = is04.ResourceVersion()
	r.Label = label
	r.Description = description
	return &r
}
