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
	} else if arrayFamily == definition.Array32 {
		size, next, err := d.readSizeN(next, 4)
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

func (d *decoder) decodeSlice(offset int, rv reflect.Value) (int, error) {
	arrayFamily, next, err := d.getTypeFamily(offset)
	if err != nil {
		return -1, err
	}

	sliceLen := 0
	if arrayFamily == definition.Nil {
		return 0, nil
	} else if d.isFixArray(arrayFamily) {
		sliceLen = int(arrayFamily - definition.ArrayFixStart)
	} else if arrayFamily == definition.Array16 {
		sliceLenByte, next2, err2 := d.readSizeN(next, 2)
		if err2 != nil {
			return -1, err2
		}

		sliceLen = int(binary.BigEndian.Uint16(sliceLenByte))
		next = next2
	} else if arrayFamily == definition.Array32 {
		sliceLenByte, next2, err2 := d.readSizeN(next, 4)
		if err2 != nil {
			return -1, err2
		}

		sliceLen = int(binary.BigEndian.Uint32(sliceLenByte))
		next = next2
	} else {
		return -1, fmt.Errorf("Decode array occured error, code: %v", arrayFamily)
	}

	slice := reflect.MakeSlice(rv.Type(), sliceLen, sliceLen)
	for i := 0; i < sliceLen; i++ {
		next, err = d.decode(slice.Index(i), next)
		if err != nil {
			return -1, err
		}
	}

	rv.Set(slice)
	return next, nil
}

func (d *decoder) isFixArray(arrayFamily byte) bool {
	return (arrayFamily >= definition.ArrayFixStart) && (arrayFamily <= definition.ArrayFixEnd)
}

func (d *decoder) isArrayFamily(arrayFamily byte) bool {
	return d.isFixArray(arrayFamily) ||
		(arrayFamily == definition.Array16) ||
		(arrayFamily == definition.Array32)
}
