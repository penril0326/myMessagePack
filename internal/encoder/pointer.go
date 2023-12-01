package encoder

import "reflect"

func (e *encoder) writePointerData(rv reflect.Value) {
	if rv.Kind() == reflect.Pointer {
		e.encode(rv.Elem().Interface())
	} else {
		e.encode(rv.Interface())
	}
}
