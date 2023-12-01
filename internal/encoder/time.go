package encoder

import (
	"reflect"
	"time"

	"github.com/penril0326/myMessagePack/internal/definition"
)

var (
	timeType         = reflect.TypeOf(time.Time{})
	timestamp        = -1
	timestampYearLen = 12
)

type timespec struct {
	seconds     int64
	nanoseconds int
}

func (e *encoder) writeTimeData(rv reflect.Value) {
	if rv.Type() == timeType {
		t := rv.Interface().(time.Time)
		ts := timespec{
			seconds:     t.Unix(),
			nanoseconds: t.Nanosecond(),
		}

		seconds := uint64(ts.seconds)
		if seconds>>34 == 0 {
			data64 := uint64(ts.nanoseconds)<<34 | seconds
			if (data64 & 0xffffffff00000000) == 0 {
				// timestamp 32
				data32 := uint32(data64)
				bytedata32 := make([]byte, 6)
				bytedata32[0] = definition.Fixext4
				bytedata32[1] = byte(timestamp)
				bytedata32[2] = byte(data32 >> 24)
				bytedata32[3] = byte(data32 >> 16)
				bytedata32[4] = byte(data32 >> 8)
				bytedata32[5] = byte(data32)
				e.data = append(e.data, bytedata32...)
			} else {
				// timestamp 64
				bytedata64 := make([]byte, 10)
				bytedata64[0] = definition.Fixext8
				bytedata64[1] = byte(timestamp)

				offset := 56
				for i := 2; i < 10; i++ {
					bytedata64[i] = byte(data64 >> offset)
					offset -= 8
				}

				e.data = append(e.data, bytedata64...)
			}
		} else {
			// timestamp 96
			bytedata96 := make([]byte, 15)
			bytedata96[0] = definition.Ext8
			bytedata96[1] = byte(timestampYearLen)
			bytedata96[2] = byte(timestamp)

			offset1 := 24
			offset2 := 56
			for i := 3; i < 15; i++ {
				if i < 7 {
					bytedata96[i] = byte(ts.nanoseconds >> offset1)
					offset1 -= 8
				} else {
					bytedata96[i] = byte(seconds >> offset2)
					offset2 -= 8
				}
			}

			e.data = append(e.data, bytedata96...)
		}
	}
}
