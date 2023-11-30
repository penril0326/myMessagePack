package encoder

import (
	"fmt"
	"reflect"

	"github.com/penril0326/myMessagePack/internal/definition"
)

type encoder struct {
	data []byte
}

func JsonToMsgPack(jsonData interface{}) ([]byte, error) {
	e := encoder{
		data: make([]byte, 0),
	}

	if err := e.encode(jsonData); err != nil {
		return nil, fmt.Errorf("Unexpected error: %s\n", err.Error())
	}

	return e.data, nil
}

func (e *encoder) encode(jsonData interface{}) error {
	rv := reflect.ValueOf(jsonData)
	switch rv.Kind() {
	case reflect.Bool:
		e.writeBoolData(rv.Bool())
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
		e.writeUintDate(rv.Uint(), e.calculateUintSize(rv.Uint()))
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		e.writeIntDate(rv.Int(), e.calculateIntSize(rv.Int()))
	case reflect.String:
		e.writeStrDate(rv.String(), e.calculateStrSize(rv.String()))
	case reflect.Float32:
		e.writeFloat32Data(rv.Float(), e.calculateFloat32Size())
	case reflect.Float64:
		e.writeFloat64Data(rv.Float(), e.calculateFloat64Size())
	case reflect.Array, reflect.Slice:
		e.writeArrayData(rv, e.calculateArraySize(rv))
	case reflect.Map:
		e.writeMapData(rv, e.calculateMapSize(rv))
	case reflect.Pointer:
	case reflect.Invalid:
		e.data = append(e.data, definition.Nil)
	default:
		return fmt.Errorf("Type not support: %v", rv.Kind())
	}

	return nil
}
