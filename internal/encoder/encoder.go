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
		e.writeBoolData(rv.Bool())
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
		e.writeUintDate(rv.Uint(), e.calculateUintSize(rv))
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		e.calculateIntSize(rv)
	case reflect.String:
		e.writeStrDate(rv.String(), e.calculateStrSize(rv))
	}

	return e.data, nil
}
