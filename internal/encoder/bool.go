package encoder

import (
	"reflect"

	"github.com/penril0326/myMessagePack/internal/definition"
)

func (e *encoder) setBool(rv reflect.Value) {
	e.writeBoolType(rv.Bool())
}

func (e *encoder) writeBoolType(value bool) {
	if value {
		e.data = append(e.data, definition.True)
	} else {
		e.data = append(e.data, definition.False)
	}
}
