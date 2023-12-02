package decoder

import (
	"encoding/binary"
	"fmt"

	"github.com/penril0326/myMessagePack/internal/definition"
)

func (d *decoder) decodeString(offset int) (string, int, error) {
	strFamily, next, err := d.getTypeFamily(offset)
	if err != nil {
		return "", -1, err
	}

	if d.isFixString(strFamily) {
		d.getFixStrLen(strFamily)
		data, next, err := d.readSizeN(next, d.getFixStrLen(strFamily))
		if err != nil {
			return "", -1, err
		}

		return string(data), next, nil
	} else if strFamily == definition.Str8 {
		return d.getString(next, 1)
	} else if strFamily == definition.Str16 {
		return d.getString(next, 2)
	} else if strFamily == definition.Str32 {
		return d.getString(next, 4)
	} else {
		return "", -1, fmt.Errorf("Decode string occured error, code: %v", strFamily)
	}
}

func (d *decoder) isFixString(strFamily byte) bool {
	return (strFamily >= definition.StrFixStart) && (strFamily <= definition.StrFixEnd)
}

func (d *decoder) getFixStrLen(strFamily byte) int {
	return int(strFamily - definition.StrFixStart)
}

func (d *decoder) getString(offset int, familySize int) (string, int, error) {
	dataLen, next, err := d.readSizeN(offset, familySize)
	if err != nil {
		return "", -1, err
	}

	readLength := 0
	if familySize == 1 {
		readLength = int(dataLen[0])
	} else if familySize == 2 {
		readLength = int(binary.BigEndian.Uint16(dataLen))
	} else {
		readLength = int(binary.BigEndian.Uint32(dataLen))
	}

	data, next, err := d.readSizeN(next, readLength)
	if err != nil {
		return "", -1, err
	}

	return string(data), next, nil
}

func (d *decoder) isStringFamily(familyCode byte) bool {
	return d.isFixString(familyCode) ||
		(familyCode == definition.Str8) ||
		(familyCode == definition.Str16) ||
		(familyCode == definition.Str32)
}
