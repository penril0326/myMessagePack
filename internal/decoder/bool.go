package decoder

import (
	"fmt"

	"github.com/penril0326/myMessagePack/internal/definition"
)

func (d *decoder) decodeBool(offset int) (bool, error) {
	typeFamily := d.data[offset]

	switch typeFamily {
	case definition.False:
		return false, nil
	case definition.True:
		return true, nil
	}

	return false, fmt.Errorf("Got a non boolean type, code: %x", typeFamily)
}