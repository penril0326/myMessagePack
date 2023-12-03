package decoder

import (
	"fmt"
	"reflect"

	"github.com/penril0326/myMessagePack/internal/definition"
)

func (d *decoder) decodeInterface(rv reflect.Value, offset int) (interface{}, int, error) {
	family, next, err := d.getTypeFamily(offset)
	if err != nil {
		return nil, -1, err
	}

	if family == definition.Nil {
		return nil, next, nil
	} else if (family == definition.True) || (family == definition.False) {
		return d.decodeBool(offset)
	} else if d.isUintFamily(family) {
		return d.decodeUint(offset)
	} else if d.isIntFamily(family) {
		return d.decodeInt(offset)
	} else if family == definition.Float32 {
		return d.decodeFloat32(offset)
	} else if family == definition.Float64 {
		return d.decodeFloat64(offset)
	} else if d.isStringFamily(family) {
		return d.decodeString(offset)
		// } else if d.isArrayFamily(family) {

	} else {
		return nil, -1, fmt.Errorf("Decode interface occured error, code: %v", family)
	}
}
