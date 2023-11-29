package encoder

import (
	"math"

	"github.com/penril0326/myMessagePack/internal/definition"
)

const (
	NegativeFixIntMin = -32 // 0xe0
	NegativeFixIntMax = -1  // 0xff
)

func (e *encoder) calculateIntSize(val int64) int {
	size := 0
	if isUnsigned(val) {
		// do nothing
	} else if isNegativeFixInt(val) {
		// do nothing
	} else if val >= math.MinInt8 {
		size = 1
	} else if val >= math.MinInt16 {
		size = 2
	} else if val >= math.MinInt32 {
		size = 4
	} else {
		// int64 or int
		size = 8
	}

	return size
}

func (e *encoder) writeIntDate(val int64, size int) {
	byteData := make([]byte, size+1)
	if isUnsigned(val) {
		e.writeUintDate(uint64(val), e.calculateUintSize(uint64(val)))
		return
	} else if isNegativeFixInt(val) {
		// negative fix int, format directly
		byteData[0] = byte(val)
	} else if val >= math.MinInt8 {
		byteData[0] = definition.Int8
		byteData[1] = byte(val)
	} else if val >= math.MinInt16 {
		byteData[0] = definition.Int16
		byteData[1] = byte(val >> 8)
		byteData[2] = byte(val)
	} else if val >= math.MinInt32 {
		byteData[0] = definition.Int32
		byteData[1] = byte(val >> 24)
		byteData[2] = byte(val >> 16)
		byteData[3] = byte(val >> 8)
		byteData[4] = byte(val)
	} else {
		byteData[0] = definition.Int64
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

func isNegativeFixInt(val int64) bool {
	return (val >= NegativeFixIntMin) && (val <= NegativeFixIntMax)
}

func isUnsigned(val int64) bool {
	return val >= 0
}
