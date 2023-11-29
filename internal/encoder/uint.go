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

func (e *encoder) writeUintDate(val uint64, size int) {
	byteData := make([]byte, size+1)
	if val <= math.MaxInt8 {
		// positive fixint, no need to append type
		byteData[0] = byte(val)
	} else if val <= math.MaxUint8 {
		byteData[0] = definition.Uint8
		byteData[1] = byte(val)
	} else if val <= math.MaxUint16 {
		byteData[0] = definition.Uint16
		byteData[1] = byte(val >> 8)
		byteData[2] = byte(val)
	} else if val <= math.MaxUint32 {
		byteData[0] = definition.Uint32
		byteData[1] = byte(val >> 24)
		byteData[2] = byte(val >> 16)
		byteData[3] = byte(val >> 8)
		byteData[4] = byte(val)
	} else {
		byteData[0] = definition.Uint64
		byteData[1] = byte(val >> 56)
		byteData[2] = byte(val >> 48)
		byteData[3] = byte(val >> 40)
		byteData[4] = byte(val >> 32)
		byteData[5] = byte(val >> 24)
		byteData[6] = byte(val >> 16)
		byteData[7] = byte(val >> 8)
		byteData[8] = byte(val)
	}

	e.data = append(e.data, byteData...)
}
