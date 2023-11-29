package encoder

import (
	"github.com/penril0326/myMessagePack/internal/definition"
)

func (e *encoder) writeBoolData(value bool) {
	if value {
		e.data = append(e.data, definition.True)
	} else {
		e.data = append(e.data, definition.False)
	}
}
