package decoder

import (
	"encoding/binary"
	"fmt"
	"math"

	"github.com/penril0326/myMessagePack/internal/definition"
)

func (d *decoder) decodeFloat32(offset int) (float32, int, error) {
	floatFamily, next, err := d.getTypeFamily(offset)
	if err != nil {
		return 0.0, -1, err
	}

	if floatFamily == definition.Nil {
		return 0.0, next, nil
	} else if d.isUintFamily(floatFamily) {
		ui, next, err := d.decodeUint(offset)
		return float32(ui), next, err
	} else if d.isIntFamily(floatFamily) {
		i, next, err := d.decodeInt(offset)
		return float32(i), next, err
	} else if floatFamily == definition.Float32 {
		data, next, err := d.readSizeN(next, 4)
		if err != nil {
			return 0.0, -1, err
		}
		return math.Float32frombits(binary.BigEndian.Uint32(data)), next, nil
	} else {
		return 0.0, -1, fmt.Errorf("Decode float32 occured error, code: %v", floatFamily)
	}
}

func (d *decoder) decodeFloat64(offset int) (float64, int, error) {
	floatFamily, next, err := d.getTypeFamily(offset)
	if err != nil {
		return 0.0, -1, err
	}

	if floatFamily == definition.Nil {
		return 0.0, next, nil
	} else if d.isUintFamily(floatFamily) {
		ui, next, err := d.decodeUint(offset)
		return float64(ui), next, err
	} else if d.isIntFamily(floatFamily) {
		i, next, err := d.decodeInt(offset)
		return float64(i), next, err
	} else if floatFamily == definition.Float32 {
		f32, next, err := d.decodeFloat32(offset)
		return float64(f32), next, err
	} else if floatFamily == definition.Float64 {
		data, next, err := d.readSizeN(next, 8)
		if err != nil {
			return 0.0, -1, err
		}
		return math.Float64frombits(binary.BigEndian.Uint64(data)), next, nil
	} else {
		return 0.0, -1, fmt.Errorf("Decode float32 occured error, code: %v", floatFamily)
	}
}

func (d *decoder) isUintFamily(familyCode byte) bool {
	return d.isPositiveFixInt(familyCode) ||
		(familyCode == definition.Uint8) ||
		(familyCode == definition.Uint16) ||
		(familyCode == definition.Uint32) ||
		(familyCode == definition.Uint64)
}

func (d *decoder) isIntFamily(familyCode byte) bool {
	return d.isNegativeFixNum(familyCode) ||
		(familyCode == definition.Int8) ||
		(familyCode == definition.Int16) ||
		(familyCode == definition.Int32) ||
		(familyCode == definition.Int64)
}
