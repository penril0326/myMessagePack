package decoder

import "github.com/penril0326/myMessagePack/internal/definition"

func (d *decoder) decodeBool(offset int) bool {
	if d.data[offset] == definition.False {
		return false
	}

	return true
}
