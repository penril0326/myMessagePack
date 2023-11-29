package encoder

import (
	"math"

	"github.com/penril0326/myMessagePack/internal/definition"
)

func (e *encoder) calculateUintSize(val uint64) int {
	size := 0
	if val <= math.MaxInt8 {
		// do nothing
	} else if val <= math.MaxUint8 {
		size = 1
	} else if val <= math.MaxUint16 {
		size = 2
	} else if val <= math.MaxUint32 {
		size = 4
	} else {
		// uint64 or uint
		size = 8
	}

	return size
}

func (e *encoder) writeUintDate(value uint64, size int) {
	byteData := make([]byte, size+1)
	if value <= math.MaxInt8 {
		// positive fixint, no need to append type
		byteData[0] = byte(value)
	} else if value <= math.MaxUint8 {
		byteData[0] = definition.Uint8
		byteData[1] = byte(value)
	} else if value <= math.MaxUint16 {
		byteData[0] = definition.Uint16
		byteData[1] = byte(value >> 8)
		byteData[2] = byte(value)
	} else if value <= math.MaxUint32 {
		byteData[0] = definition.Uint32
		byteData[1] = byte(value >> 24)
		byteData[2] = byte(value >> 16)
		byteData[3] = byte(value >> 8)
		byteData[4] = byte(value)
	} else {
		byteData[0] = definition.Uint64
		byteData[1] = byte(value >> 56)
		byteData[2] = byte(value >> 48)
		byteData[3] = byte(value >> 40)
		byteData[4] = byte(value >> 32)
		byteData[5] = byte(value >> 24)
		byteData[6] = byte(value >> 16)
		byteData[7] = byte(value >> 8)
		byteData[8] = byte(value)
	}

	e.data = append(e.data, byteData...)
}
