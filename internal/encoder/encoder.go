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
		e.setBool(rv)
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
		e.setUint(rv)
	}

	return e.data, nil
}
