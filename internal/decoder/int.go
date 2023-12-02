package decoder

import (
	"encoding/binary"
	"fmt"

	"github.com/penril0326/myMessagePack/internal/definition"
)

func (d *decoder) decodeInt(offset int) (int64, int, error) {
	numberFamily, next, err := d.getTypeFamily(offset)
	if err != nil {
		return 0, -1, err
	}

	if numberFamily == definition.Nil {
		return 0, next, nil
	} else if d.isPositiveFixInt(numberFamily) {
		return int64(numberFamily), next, nil
	} else if d.isNegativeFixNum(numberFamily) {
		return int64(int8(numberFamily)), next, nil
	} else if numberFamily == definition.Uint8 {
		data, next, err := d.readSizeN(next, 1)
		if err != nil {
			return 0, -1, err
		}
		return int64(data[0]), next, nil
	} else if numberFamily == definition.Uint16 {
		data, next, err := d.readSizeN(next, 2)
		if err != nil {
			return 0, -1, err
		}
		return int64(binary.BigEndian.Uint16(data)), next, nil
	} else if numberFamily == definition.Uint32 {
		data, next, err := d.readSizeN(next, 4)
		if err != nil {
			return 0, -1, err
		}
		return int64(binary.BigEndian.Uint32(data)), next, nil
	} else if numberFamily == definition.Uint64 {
		data, next, err := d.readSizeN(next, 8)
		if err != nil {
			return 0, -1, err
		}
		return int64(binary.BigEndian.Uint64(data)), next, nil
	} else if numberFamily == definition.Int8 {
		data, next, err := d.readSizeN(next, 1)
		if err != nil {
			return 0, -1, err
		}
		return int64(int8(data[0])), next, nil
	} else if numberFamily == definition.Int16 {
		data, next, err := d.readSizeN(next, 2)
		if err != nil {
			return 0, -1, err
		}
		return int64(int16(binary.BigEndian.Uint16(data))), next, nil
	} else if numberFamily == definition.Int32 {
		data, next, err := d.readSizeN(next, 4)
		if err != nil {
			return 0, -1, err
		}
		return int64(int32(binary.BigEndian.Uint32(data))), next, nil
	} else if numberFamily == definition.Int64 {
		data, next, err := d.readSizeN(next, 8)
		if err != nil {
			return 0, -1, err
		}
		return int64(int64(binary.BigEndian.Uint64(data))), next, nil
	} else if numberFamily == definition.Float32 {
		f32, next, err := d.decodeFloat32(offset)
		return int64(f32), next, err
	} else if numberFamily == definition.Float64 {
		f64, next, err := d.decodeFloat64(offset)
		return int64(f64), next, err
	} else {
		return 0, -1, fmt.Errorf("Decode uint occured error, code: %v", numberFamily)
	}
}

func (d *decoder) isIntFamily(familyCode byte) bool {
	return d.isNegativeFixNum(familyCode) ||
		(familyCode == definition.Int8) ||
		(familyCode == definition.Int16) ||
		(familyCode == definition.Int32) ||
		(familyCode == definition.Int64)
}
