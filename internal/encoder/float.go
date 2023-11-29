package encoder

import (
	"math"

	"github.com/penril0326/myMessagePack/internal/definition"
)

func (e *encoder) calculateFloat32Size() int {
	return 4
}

func (e *encoder) calculateFloat64Size() int {
	return 8
}

func (e *encoder) writeFloat32Data(val float64, size int) {
	bytedata := make([]byte, size+1)
	bytedata[0] = definition.Float32

	binary := math.Float32bits(float32(val))

	bytedata[1] = byte(binary >> 24)
	bytedata[2] = byte(binary >> 16)
	bytedata[3] = byte(binary >> 8)
	bytedata[4] = byte(binary)

	e.data = append(e.data, bytedata...)
}

func (e *encoder) writeFloat64Data(val float64, size int) {
	bytedata := make([]byte, size+1)
	bytedata[0] = definition.Float64

	binary := math.Float64bits(val)

	bytedata[1] = byte(binary >> 56)
	bytedata[2] = byte(binary >> 48)
	bytedata[3] = byte(binary >> 40)
	bytedata[4] = byte(binary >> 32)
	bytedata[5] = byte(binary >> 24)
	bytedata[6] = byte(binary >> 16)
	bytedata[7] = byte(binary >> 8)
	bytedata[8] = byte(binary)

	e.data = append(e.data, bytedata...)
}
