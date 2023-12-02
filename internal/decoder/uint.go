package decoder

import (
	"encoding/binary"
	"fmt"

	"github.com/penril0326/myMessagePack/internal/definition"
)

func (d *decoder) decodeUint(offset int) (uint64, int, error) {
	numberFamily, next, err := d.getTypeFamily(offset)
	if err != nil {
		return 0, -1, err
	}

	if numberFamily == definition.Nil {
		return 0, next, nil
	} else if d.isPositiveFixInt(numberFamily) {
		return uint64(numberFamily), next, nil
	} else if d.isNegativeFixNum(numberFamily) {
		return uint64(int8(numberFamily)), next, nil
	} else if numberFamily == definition.Uint8 {
		data, next, err := d.readSizeN(next, 1)
		if err != nil {
			return 0, -1, err
		}
		return uint64(data[0]), next, nil
	} else if numberFamily == definition.Uint16 {
		data, next, err := d.readSizeN(next, 2)
		if err != nil {
			return 0, -1, err
		}
		return uint64(binary.BigEndian.Uint16(data)), next, nil
	} else if numberFamily == definition.Uint32 {
		data, next, err := d.readSizeN(next, 4)
		if err != nil {
			return 0, -1, err
		}
		return uint64(binary.BigEndian.Uint32(data)), next, nil
	} else if numberFamily == definition.Uint64 {
		data, next, err := d.readSizeN(next, 8)
		if err != nil {
			return 0, -1, err
		}
		return binary.BigEndian.Uint64(data), next, nil
	} else if numberFamily == definition.Int8 {
		data, next, err := d.readSizeN(next, 1)
		if err != nil {
			return 0, -1, err
		}
		return uint64(int8(data[0])), next, nil
	} else if numberFamily == definition.Int16 {
		data, next, err := d.readSizeN(next, 2)
		if err != nil {
			return 0, -1, err
		}
		return uint64(int16(binary.BigEndian.Uint16(data))), next, nil
	} else if numberFamily == definition.Int32 {
		data, next, err := d.readSizeN(next, 4)
		if err != nil {
			return 0, -1, err
		}
		return uint64(int32(binary.BigEndian.Uint32(data))), next, nil
	} else if numberFamily == definition.Int64 {
		data, next, err := d.readSizeN(next, 8)
		if err != nil {
			return 0, -1, err
		}
		return uint64(int64(binary.BigEndian.Uint64(data))), next, nil
	} else {
		return 0, -1, fmt.Errorf("Decode uint occured error, code: %v", numberFamily)
	}
}

func (d *decoder) parseUint(from, to int) uint {
	return 0
}

func (d *decoder) isPositiveFixInt(numberFamily byte) bool {
	return (numberFamily >= 0x00) && (numberFamily <= 0x7f)
}

func (d *decoder) isNegativeFixNum(numberFamily byte) bool {
	return int8(numberFamily) >= definition.NegativeFixIntMin && int8(numberFamily) <= definition.NegativeFixIntMax
}
