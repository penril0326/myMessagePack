package encoder

import (
	"reflect"

	"github.com/penril0326/myMessagePack/internal/definition"
)

type encoder struct {
	data []byte
}

func JsonToMsgPack(jsonData interface{}) ([]byte, error) {
	if jsonData == nil {
		return []byte{definition.Nil}, nil
	}

	e := encoder{
		data: make([]byte, 0),
	}

	return e.encode(jsonData)
}

func (e *encoder) encode(jsonData interface{}) ([]byte, error) {
	rv := reflect.ValueOf(jsonData)
	switch rv.Kind() {
	case reflect.Bool:
		v := rv.Bool()
		e.writeBoolData(v)
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
		v := rv.Uint()
		e.writeUintDate(v, e.calculateUintSize(v))
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		v := rv.Int()
		e.writeIntDate(v, e.calculateIntSize(v))
	case reflect.String:
		v := rv.String()
		e.writeStrDate(v, e.calculateStrSize(v))
	}

	return e.data, nil
}
