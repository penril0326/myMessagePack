package encoder

import (
	"math"
	"reflect"

	"github.com/penril0326/myMessagePack/internal/definition"
)

const (
	FixMapSize  = 0x0f // 0x80 ~ 0x8f
	FixMapStart = 0x80
)

func (e *encoder) calculateMapSize(rv reflect.Value) int {
	length := rv.Len()
	size := 0
	if length <= FixMapSize {
		// do nothing
	} else if length <= math.MaxUint16 {
		size = 2
	} else {
		size = 4
	}

	return size
}

func (e *encoder) writeMapData(rv reflect.Value, size int) {
	bytedata := make([]byte, size+1)
	length := rv.Len()
	if length <= FixMapSize {
		bytedata[0] = byte(FixMapStart + length)
	} else if length <= math.MaxUint16 {
		bytedata[0] = definition.Map16
		bytedata[1] = byte(length >> 8)
		bytedata[2] = byte(length)
	} else {
		bytedata[0] = definition.Map32
		bytedata[1] = byte(length >> 24)
		bytedata[2] = byte(length >> 16)
		bytedata[3] = byte(length >> 8)
		bytedata[4] = byte(length)
	}

	e.data = append(e.data, bytedata...)

	iter := rv.MapRange()
	for iter.Next() {
		key := iter.Key().Interface()
		value := iter.Value().Interface()
		e.encode(key)
		e.encode(value)
	}
}
