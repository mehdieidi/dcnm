package frontend

import (
	"github.com/MehdiEidi/dcnm/core"
)

type FrontEnd interface {
	Start(kv *core.KeyValueStore) error
}

type zeroFrontEnd struct{}

func (f zeroFrontEnd) Start(kv *core.KeyValueStore) error {
	return nil
}
