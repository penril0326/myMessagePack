package decoder

import (
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
