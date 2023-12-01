package decoder

import (
	"fmt"
	"reflect"
)

type decoder struct {
	data []byte
}

func MsgPackToJson(data []byte, v interface{}) error {
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
		v := d.decodeBool(curIdx)
		rv.SetBool(v)
		curIdx++
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
	case reflect.Invalid:
	default:
		return fmt.Errorf("Got unexpected type: %v", kind)
	}

	return nil
}
