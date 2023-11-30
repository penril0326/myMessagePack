package encoder

import (
	"math"
	"reflect"

	"github.com/penril0326/myMessagePack/internal/definition"
)

const (
	FixArraySize  = 15 // 0x90 ~ 0x9f
	FixArrayStart = 0x90
)

func (e *encoder) calculateArraySize(rv reflect.Value) int {
	length := rv.Len()
	size := 0
	if length <= FixArraySize {
		// do nothing
	} else if length <= math.MaxUint16 {
		size = 2
	} else {
		size = 4
	}

	return size
}

func (e *encoder) writeArrayData(rv reflect.Value, size int) {
	bytedata := make([]byte, size+1)
	length := rv.Len()
	if length <= FixArraySize {
		bytedata[0] = byte(FixArrayStart + length)
	} else if length <= math.MaxUint16 {
		bytedata[0] = definition.Array16
		bytedata[1] = byte(length >> 8)
		bytedata[2] = byte(length)
	} else {
		bytedata[0] = definition.Array32
		bytedata[1] = byte(length >> 24)
		bytedata[2] = byte(length >> 16)
		bytedata[3] = byte(length >> 8)
		bytedata[4] = byte(length)
	}

	e.data = append(e.data, bytedata...)

	for i := 0; i < length; i++ {
		ele := rv.Index(i)
		e.encode(ele.Interface())
	}
}
