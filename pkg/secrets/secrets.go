package secrets

import (
	"fmt"
)

// -- Enum -- //

type SecretEnum struct {
	slug string
}

func (o SecretEnum) String() string {
	return o.slug
}

var (
	invalid = SecretEnum{"invalid"}
	Sops    = SecretEnum{"sops"}
)

func FromString(s string) (SecretEnum, error) {
	switch s {
	case Sops.slug:
		return Sops, nil
	}

	return invalid, fmt.Errorf("no valid secret found in input: %s", s)
}

// -- Factory -- //

type SecretActions interface {
	Export()
	Save(prefix string)
}
