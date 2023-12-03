package decoder

import (
	"encoding/binary"
	"fmt"
	"reflect"

	"github.com/penril0326/myMessagePack/internal/definition"
)

func (d *decoder) decodeMap(offset int, rv reflect.Value) (int, error) {
	mapFamily, next, err := d.getTypeFamily(offset)
	if err != nil {
		return -1, err
	}

	objectCount := 0
	if mapFamily == definition.Nil {
		return 0, nil
	} else if d.isFixMap(mapFamily) {
		objectCount = int(mapFamily - definition.MapFixStart)
	} else if mapFamily == definition.Map16 {
		objectCountByte, next2, err2 := d.readSizeN(next, 2)
		if err2 != nil {
			return -1, err2
		}

		objectCount = int(binary.BigEndian.Uint16(objectCountByte))
		next = next2
	} else if mapFamily == definition.Map32 {
		objectCountByte, next2, err2 := d.readSizeN(next, 4)
		if err2 != nil {
			return -1, err2
		}

		objectCount = int(binary.BigEndian.Uint32(objectCountByte))
		next = next2
	} else {
		return -1, fmt.Errorf("Decode map occured error, code: %v", mapFamily)
	}

	keyType := rv.Type().Key()
	valueType := rv.Type().Elem()
	if rv.IsNil() {
		newMap := reflect.MakeMapWithSize(rv.Type(), objectCount)
		rv.Set(newMap)
	}

	for i := 0; i < objectCount; i++ {
		newKey := reflect.New(keyType).Elem()
		newVal := reflect.New(valueType).Elem()
		next2, err2 := d.decode(newKey, next)
		if err2 != nil {
			return -1, err
		}

		next2, err2 = d.decode(newVal, next2)
		if err2 != nil {
			return -1, err
		}

		rv.SetMapIndex(newKey, newVal)
		next = next2
	}

	return next, nil
}

func (d *decoder) isFixMap(code byte) bool {
	return (code >= definition.MapFixStart) && (code <= definition.MapFixEnd)
}
