package decoder

import (
	"encoding/binary"
	"fmt"
	"reflect"
)

type decoder struct {
	data []byte
}

func MsgPackToJson(data []byte, v interface{}) error {
	if len(data) == 0 {
		return fmt.Errorf("Empty data")
	}

	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Pointer {
		return fmt.Errorf("v should be pointer: %v\n", rv.Kind())
	}

	d := decoder{
		data: data,
	}

	return d.decode(rv.Elem(), 0)
}

func (d *decoder) decode(rv reflect.Value, curIdx int) error {
	kind := rv.Kind()
	switch kind {
	case reflect.Bool:
		v, next, err := d.decodeBool(curIdx)
		if err != nil {
			return err
		}

		rv.SetBool(v)
		curIdx = next
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		ui, next, err := d.decodeUint(curIdx)
		if err != nil {
			return err
		}

		rv.SetUint(ui)
		curIdx = next
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i, next, err := d.decodeInt(curIdx)
		if err != nil {
			return err
		}

		rv.SetInt(i)
		curIdx = next
	case reflect.Float32:
		f32, next, err := d.decodeFloat32(curIdx)
		if err != nil {
			return err
		}

		rv.SetFloat(float64(f32))
		curIdx = next
	case reflect.Float64:
		f64, next, err := d.decodeFloat64(curIdx)
		if err != nil {
			return err
		}

		rv.SetFloat(f64)
		curIdx = next
	case reflect.String:
		s, next, err := d.decodeString(curIdx)
		if err != nil {
			return err
		}

		rv.SetString(s)
		curIdx = next
	case reflect.Invalid:
	default:
		return fmt.Errorf("Got unexpected type: %v", kind)
	}

	return nil
}

func (d *decoder) getTypeFamily(offset int) (byte, int, error) {
	if len(d.data) < offset {
		return 0, -1, fmt.Errorf("Can not get type family code, data byte too short")
	}

	nextIdx := offset + 1
	return d.data[offset], nextIdx, nil
}

func (d *decoder) readSizeN(start, n int) ([]byte, int, error) {
	if len(d.data) < start+n {
		return nil, -1, fmt.Errorf("Data too short to get")
	}

	return d.data[start : start+n], start + n, nil
}

func (d *decoder) getFamilySize(dataLen []byte) uint {
	length := len(dataLen)
	if length == 0 {
		return 0
	} else if length == 1 {
		return uint(dataLen[0])
	} else if length == 2 {
		return uint(binary.BigEndian.Uint16(dataLen))
	} else if length == 4 {
		return uint(binary.BigEndian.Uint32(dataLen))
	} else {
		return uint(binary.BigEndian.Uint64(dataLen))
	}
}
