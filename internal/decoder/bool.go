package decoder

import (
	"fmt"

	"github.com/penril0326/myMessagePack/internal/definition"
)

func (d *decoder) decodeBool(offset int) (bool, int, error) {
	typeFamily := d.data[offset]
	switch typeFamily {
	case definition.False:
		next := offset + 1
		return false, next, nil
	case definition.True:
		next := offset + 1
		return true, next, nil
	}

	return false, -1, fmt.Errorf("Got a non boolean type, code: %x", typeFamily)
}
