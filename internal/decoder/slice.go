package decoder

import (
	"encoding/binary"
	"fmt"
	"reflect"

	"github.com/penril0326/myMessagePack/internal/definition"
)

func (d *decoder) decodeArray(offset int, rv reflect.Value) (int, error) {
	arrayFamily, next, err := d.getTypeFamily(offset)
	if err != nil {
		return -1, err
	}

	if arrayFamily == definition.Nil {
		return 0, nil
	} else if d.isFixArray(arrayFamily) {
		for i := 0; i < rv.Len(); i++ {
			next, err = d.decode(rv.Index(i), next)
			if err != nil {
				return -1, err
			}
		}

		return next, nil
	} else if arrayFamily == definition.Array16 {
		size, next, err := d.readSizeN(next, 2)
		if err != nil {
			return -1, err
		}

		arraySize := int(binary.BigEndian.Uint16(size))
		if arraySize != rv.Len() {
			return -1, fmt.Errorf("Data length is not equivalent")
		}

		for i := 0; i < rv.Len(); i++ {
			next, err = d.decode(rv.Index(i), next)
			if err != nil {
				return -1, err
			}
		}

		return next, nil
	} else {
		return -1, fmt.Errorf("Decode array occured error, code: %v", arrayFamily)
	}
}

func (d *decoder) isFixArray(arrayFamily byte) bool {
	return (arrayFamily >= definition.ArrayFixStart) && (arrayFamily <= definition.ArrayFixEnd)
}

func (d *decoder) isArrayFamily(arrayFamily byte) bool {
	return d.isFixArray(arrayFamily) ||
		(arrayFamily == definition.Array16) ||
		(arrayFamily == definition.Array32)
}
