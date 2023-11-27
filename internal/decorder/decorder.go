package decoder

import (
	"errors"
	"reflect"
)

type decoder struct {
	data []byte
}

func Decode(data []byte, v interface{}) error {
	if (len(data) == 0) || (data == nil) {
		return errors.New("data is empty")
	}

	r := reflect.ValueOf(v)
	if r.Kind() != reflect.Pointer {
		return errors.New("v is not a pointer")
	}

	// d := decoder{
	// 	data: data,
	// }

	return nil
}
