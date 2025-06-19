package is04

import (
	"log"

	"github.com/google/uuid"
)

type ResourceCore struct {
	ID          string   `json:"id"`
	Version     string   `json:"version"`
	Label       string   `json:"label"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}

type IS04Resource interface {
	encode() ([]byte, error)
}

func NewResourceCore(
	label string,
	description string,
	random bool,
) *ResourceCore {
	var resourceId string

	if !random {
		namespace := uuid.MustParse("00000000-0000-0000-0000-000000000000")
		u := uuid.NewMD5(namespace, []byte(label+description))
		resourceId = u.String()
	} else {
		u, err := uuid.NewRandom()
		if err != nil {
			log.Fatalf("Unable to generate random UUID. Error [%s]", err)
		}
		resourceId = u.String()
	}

	return &ResourceCore{
		ID:          resourceId,
		Version:     ResourceVersion(),
		Label:       label,
		Description: description,
		Tags:        []string{},
	}
}
