package encoder

import (
	"math"

	"github.com/penril0326/myMessagePack/internal/definition"
)

func (e *encoder) calculateStrSize(val string) int {
	length := len(val)
	extraSize := 0
	if length < 32 {
		// do nothing
	} else if length <= math.MaxUint8 {
		extraSize += 1
	} else if length <= math.MaxUint16 {
		extraSize += 2
	} else {
		extraSize += 4
	}

	return extraSize
}

func (e *encoder) writeStrDate(val string, size int) {
	bytedata := make([]byte, size+1)
	length := len(val)
	if length < 32 {
		bytedata[0] = byte(0xa0 + len(val))
	} else if length <= math.MaxUint8 {
		bytedata[0] = definition.Str8
		bytedata[1] = byte(length)
	} else if length <= math.MaxUint16 {
		bytedata[0] = definition.Str16
		bytedata[1] = byte(length >> 8)
		bytedata[2] = byte(length)
	} else {
		bytedata[0] = definition.Str32
		bytedata[1] = byte(length >> 24)
		bytedata[2] = byte(length >> 16)
		bytedata[3] = byte(length >> 8)
		bytedata[4] = byte(length)
	}

	e.data = append(e.data, bytedata...)
	e.data = append(e.data, []byte(val)...)
}
