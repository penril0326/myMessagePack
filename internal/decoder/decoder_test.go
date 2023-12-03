package decoder

import (
	"reflect"
	"testing"
)

type f func()

func TestBool(t *testing.T) {
	tests := []f{
		func() {
			data := []byte{0xc2}
			var want bool
			err := MsgPackToJson(data, &want)
			t.Run("Bool: 1", func(t *testing.T) {
				if err != nil {
					t.Fatalf("MsgPackToJson got err %s", err.Error())
				}
				if want != false {
					t.Errorf("MsgPackToJson got %v, want %v", want, false)
				}
			})
		},
		func() {
			data := []byte{0xc3}
			var want bool
			err := MsgPackToJson(data, &want)
			t.Run("Bool: 2", func(t *testing.T) {
				if err != nil {
					t.Fatalf("MsgPackToJson got err %s", err.Error())
				}
				if want != true {
					t.Errorf("MsgPackToJson got %v, want %v", want, true)
				}
			})
		},
	}

	for _, fun := range tests {
		fun()
	}
}

func TestUint(t *testing.T) {
	tests := []f{
		func() {
			data := []byte{0x00}
			var want uint8
			err := MsgPackToJson(data, &want)
			t.Run("positivefix: 1", func(t *testing.T) {
				if err != nil {
					t.Fatalf("MsgPackToJson got err %s", err.Error())
				}
				if want != 0 {
					t.Errorf("MsgPackToJson got %v, want %v", want, 0)
				}
			})
		},
		func() {
			data := []byte{0x7f}
			var want uint8
			err := MsgPackToJson(data, &want)
			t.Run("positivefix: 2", func(t *testing.T) {
				if err != nil {
					t.Fatalf("MsgPackToJson got err %s", err.Error())
				}
				if want != 127 {
					t.Errorf("MsgPackToJson got %v, want %v", want, 127)
				}
			})
		},
		func() {
			data := []byte{0xe0}
			var want uint8
			err := MsgPackToJson(data, &want)
			t.Run("negativefix: 1", func(t *testing.T) {
				if err != nil {
					t.Fatalf("MsgPackToJson got err %s", err.Error())
				}
				if want != 224 {
					t.Errorf("MsgPackToJson got %v, want %v", want, 224)
				}
			})
		},
		func() {
			data := []byte{0xff}
			var want uint8
			err := MsgPackToJson(data, &want)
			t.Run("negativefix: 2", func(t *testing.T) {
				if err != nil {
					t.Fatalf("MsgPackToJson got err %s", err.Error())
				}
				if want != 255 {
					t.Errorf("MsgPackToJson got %v, want %v", want, 255)
				}
			})
		},
		func() {
			data := []byte{0xcc, 0x80}
			var want uint8
			err := MsgPackToJson(data, &want)
			t.Run("uint8: 1", func(t *testing.T) {
				if err != nil {
					t.Fatalf("MsgPackToJson got err %s", err.Error())
				}
				if want != 128 {
					t.Errorf("MsgPackToJson got %v, want %v", want, 128)
				}
			})
		},
		func() {
			data := []byte{0xcc, 0xff}
			var want uint8
			err := MsgPackToJson(data, &want)
			t.Run("uint8: 2", func(t *testing.T) {
				if err != nil {
					t.Fatalf("MsgPackToJson got err %s", err.Error())
				}
				if want != 255 {
					t.Errorf("MsgPackToJson got %v, want %v", want, 255)
				}
			})
		},
		func() {
			data := []byte{0xd0, 0xff}
			var want uint8
			err := MsgPackToJson(data, &want)
			t.Run("uint8: 3", func(t *testing.T) {
				if err != nil {
					t.Fatalf("MsgPackToJson got err %s", err.Error())
				}
				if want != 255 {
					t.Errorf("MsgPackToJson got %v, want %v", want, 255)
				}
			})
		},
		func() {
			data := []byte{0xd0, 0x00}
			var want uint8
			err := MsgPackToJson(data, &want)
			t.Run("uint8: 4", func(t *testing.T) {
				if err != nil {
					t.Fatalf("MsgPackToJson got err %s", err.Error())
				}
				if want != 0 {
					t.Errorf("MsgPackToJson got %v, want %v", want, 0)
				}
			})
		},
		func() {
			data := []byte{0xcd, 0x01, 0x00}
			var want uint16
			err := MsgPackToJson(data, &want)
			t.Run("uint16: 1", func(t *testing.T) {
				if err != nil {
					t.Fatalf("MsgPackToJson got err %s", err.Error())
				}
				if want != 256 {
					t.Errorf("MsgPackToJson got %v, want %v", want, 256)
				}
			})
		},
		func() {
			data := []byte{0xcd, 0xff, 0xff}
			var want uint16
			err := MsgPackToJson(data, &want)
			t.Run("uint16: 2", func(t *testing.T) {
				if err != nil {
					t.Fatalf("MsgPackToJson got err %s", err.Error())
				}
				if want != 65535 {
					t.Errorf("MsgPackToJson got %v, want %v", want, 65535)
				}
			})
		},
		func() {
			data := []byte{0xd1, 0xff, 0xff}
			var want uint16
			err := MsgPackToJson(data, &want)
			t.Run("uint16: 3", func(t *testing.T) {
				if err != nil {
					t.Fatalf("MsgPackToJson got err %s", err.Error())
				}
				if want != 65535 {
					t.Errorf("MsgPackToJson got %v, want %v", want, 65535)
				}
			})
		},
		func() {
			data := []byte{0xd1, 0x80, 0x00}
			var want uint16
			err := MsgPackToJson(data, &want)
			t.Run("uint16: 4", func(t *testing.T) {
				if err != nil {
					t.Fatalf("MsgPackToJson got err %s", err.Error())
				}
				if want != 32768 {
					t.Errorf("MsgPackToJson got %v, want %v", want, 32768)
				}
			})
		},
		func() {
			data := []byte{0xce, 0x00, 0x01, 0x00, 0x00}
			var want uint32
			err := MsgPackToJson(data, &want)
			t.Run("uint32: 1", func(t *testing.T) {
				if err != nil {
					t.Fatalf("MsgPackToJson got err %s", err.Error())
				}
				if want != 65536 {
					t.Errorf("MsgPackToJson got %v, want %v", want, 65536)
				}
			})
		},
		func() {
			data := []byte{0xce, 0xff, 0xff, 0xff, 0xff}
			var want uint32
			err := MsgPackToJson(data, &want)
			t.Run("uint32: 1", func(t *testing.T) {
				if err != nil {
					t.Fatalf("MsgPackToJson got err %s", err.Error())
				}
				if want != 4_294_967_295 {
					t.Errorf("MsgPackToJson got %v, want %v", want, 4_294_967_295)
				}
			})
		},
		func() {
			data := []byte{0xcf, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00}
			var want uint64
			err := MsgPackToJson(data, &want)
			t.Run("uint64: 1", func(t *testing.T) {
				if err != nil {
					t.Fatalf("MsgPackToJson got err %s", err.Error())
				}
				if want != 4_294_967_296 {
					t.Errorf("MsgPackToJson got %v, want %v", want, 4_294_967_296)
				}
			})
		},
		func() {
			data := []byte{0xcf, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}
			var want uint64
			err := MsgPackToJson(data, &want)
			t.Run("uint64: 2", func(t *testing.T) {
				if err != nil {
					t.Fatalf("MsgPackToJson got err %s", err.Error())
				}
				if want != 18_446_744_073_709_551_615 {
					t.Errorf("MsgPackToJson got %v, want %v", want, uint64(18_446_744_073_709_551_615))
				}
			})
		},
		func() {
			data := []byte{0xcb, 0x3f, 0xf1, 0xf9, 0xa6, 0xb6, 0xc6, 0xd6, 0xf6}
			var want uint64
			err := MsgPackToJson(data, &want)
			t.Run("uint64: 2", func(t *testing.T) {
				if err != nil {
					t.Fatalf("MsgPackToJson got err %s", err.Error())
				}
				if want != 1 {
					t.Errorf("MsgPackToJson got %v, want %v", want, 1)
				}
			})
		},
	}

	for _, fun := range tests {
		fun()
	}
}

func TestInt(t *testing.T) {
	tests := []f{
		func() {
			data := []byte{0x00}
			var want int8
			err := MsgPackToJson(data, &want)
			t.Run("positivefix: 1", func(t *testing.T) {
				if err != nil {
					t.Fatalf("MsgPackToJson got err %s", err.Error())
				}
				if want != 0 {
					t.Errorf("MsgPackToJson got %v, want %v", want, 0)
				}
			})
		},
		func() {
			data := []byte{0x7f}
			var want int8
			err := MsgPackToJson(data, &want)
			t.Run("positivefix: 2", func(t *testing.T) {
				if err != nil {
					t.Fatalf("MsgPackToJson got err %s", err.Error())
				}
				if want != 127 {
					t.Errorf("MsgPackToJson got %v, want %v", want, 127)
				}
			})
		},
		func() {
			data := []byte{0xe0}
			var want int8
			err := MsgPackToJson(data, &want)
			t.Run("negativefix: 1", func(t *testing.T) {
				if err != nil {
					t.Fatalf("MsgPackToJson got err %s", err.Error())
				}
				if want != -32 {
					t.Errorf("MsgPackToJson got %v, want %v", want, -32)
				}
			})
		},
		func() {
			data := []byte{0xff}
			var want int8
			err := MsgPackToJson(data, &want)
			t.Run("negativefix: 2", func(t *testing.T) {
				if err != nil {
					t.Fatalf("MsgPackToJson got err %s", err.Error())
				}
				if want != -1 {
					t.Errorf("MsgPackToJson got %v, want %v", want, -1)
				}
			})
		},
		func() {
			data := []byte{0xd0, 0x80}
			var want int8
			err := MsgPackToJson(data, &want)
			t.Run("int8: 1", func(t *testing.T) {
				if err != nil {
					t.Fatalf("MsgPackToJson got err %s", err.Error())
				}
				if want != -128 {
					t.Errorf("MsgPackToJson got %v, want %v", want, -128)
				}
			})
		},
		func() {
			data := []byte{0xd0, 0x7f}
			var want int8
			err := MsgPackToJson(data, &want)
			t.Run("int8: 2", func(t *testing.T) {
				if err != nil {
					t.Fatalf("MsgPackToJson got err %s", err.Error())
				}
				if want != 127 {
					t.Errorf("MsgPackToJson got %v, want %v", want, 127)
				}
			})
		},
		func() {
			data := []byte{0xd1, 0x80, 0x00}
			var want int16
			err := MsgPackToJson(data, &want)
			t.Run("int16: 1", func(t *testing.T) {
				if err != nil {
					t.Fatalf("MsgPackToJson got err %s", err.Error())
				}
				if want != -32768 {
					t.Errorf("MsgPackToJson got %v, want %v", want, -32768)
				}
			})
		},
		func() {
			data := []byte{0xd1, 0x7f, 0xff}
			var want int16
			err := MsgPackToJson(data, &want)
			t.Run("int16: 2", func(t *testing.T) {
				if err != nil {
					t.Fatalf("MsgPackToJson got err %s", err.Error())
				}
				if want != 32767 {
					t.Errorf("MsgPackToJson got %v, want %v", want, 32767)
				}
			})
		},
		func() {
			data := []byte{0xd2, 0x7f, 0xff, 0xff, 0xff}
			var want int32
			err := MsgPackToJson(data, &want)
			t.Run("int32: 1", func(t *testing.T) {
				if err != nil {
					t.Fatalf("MsgPackToJson got err %s", err.Error())
				}
				if want != 2_147_483_647 {
					t.Errorf("MsgPackToJson got %v, want %v", want, 2_147_483_647)
				}
			})
		},
		func() {
			data := []byte{0xd2, 0x80, 0x00, 0x00, 0x00}
			var want int32
			err := MsgPackToJson(data, &want)
			t.Run("int32: 2", func(t *testing.T) {
				if err != nil {
					t.Fatalf("MsgPackToJson got err %s", err.Error())
				}
				if want != -2_147_483_648 {
					t.Errorf("MsgPackToJson got %v, want %v", want, -2_147_483_648)
				}
			})
		},
		func() {
			data := []byte{0xd3, 0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
			var want int64
			err := MsgPackToJson(data, &want)
			t.Run("int64: 1", func(t *testing.T) {
				if err != nil {
					t.Fatalf("MsgPackToJson got err %s", err.Error())
				}
				if want != -9_223_372_036_854_775_808 {
					t.Errorf("MsgPackToJson got %v, want %v", want, -9_223_372_036_854_775_808)
				}
			})
		},
		func() {
			data := []byte{0xd3, 0x7f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}
			var want int64
			err := MsgPackToJson(data, &want)
			t.Run("int64: 2", func(t *testing.T) {
				if err != nil {
					t.Fatalf("MsgPackToJson got err %s", err.Error())
				}
				if want != 9_223_372_036_854_775_807 {
					t.Errorf("MsgPackToJson got %v, want %v", want, 9_223_372_036_854_775_807)
				}
			})
		},
		func() {
			data := []byte{0xca, 0x3f, 0xf1, 0xf9, 0xa6}
			var want int64
			err := MsgPackToJson(data, &want)
			t.Run("int64: 3", func(t *testing.T) {
				if err != nil {
					t.Fatalf("MsgPackToJson got err %s", err.Error())
				}
				if want != 1 {
					t.Errorf("MsgPackToJson got %v, want %v", want, 1)
				}
			})
		},
		func() {
			data := []byte{0xce, 0x3f, 0xf1, 0xf9, 0xa6}
			var want int
			err := MsgPackToJson(data, &want)
			t.Run("int: 1", func(t *testing.T) {
				if err != nil {
					t.Fatalf("MsgPackToJson got err %s", err.Error())
				}
				if want != 1_072_822_694 {
					t.Errorf("MsgPackToJson got %v, want %v", want, 1_072_822_694)
				}
			})
		},
		func() {
			data := []byte{0xcf, 0x3f, 0xf1, 0xf9, 0x01}
			var want int
			err := MsgPackToJson(data, &want)
			t.Run("int: 2", func(t *testing.T) {
				if err == nil {
					t.Fatalf("MsgPackToJson error got nil, want an error")
				}
			})
		},
		func() {
			data := []byte{0xCB, 0x40, 0xC8, 0x1C, 0xC5, 0x87, 0xE7, 0xC0, 0x6E}
			var want int
			err := MsgPackToJson(data, &want)
			t.Run("int: 3", func(t *testing.T) {
				if err != nil {
					t.Fatalf("MsgPackToJson got err %s", err.Error())
				}
				if want != 12345 {
					t.Errorf("MsgPackToJson got %v, want %v", want, 12345)
				}
			})
		},
	}

	for _, fun := range tests {
		fun()
	}
}

func TestFloat32(t *testing.T) {
	tests := []f{
		func() {
			data := []byte{0x7f}
			var want float32
			err := MsgPackToJson(data, &want)
			t.Run("float32: 1", func(t *testing.T) {
				if err != nil {
					t.Fatalf("MsgPackToJson got err %s", err.Error())
				}
				if want != 127.0 {
					t.Errorf("MsgPackToJson got %v, want %v", want, 127.0)
				}
			})
		},
		func() {
			data := []byte{0xec}
			var want float32
			err := MsgPackToJson(data, &want)
			t.Run("flaot32: 2", func(t *testing.T) {
				if err != nil {
					t.Fatalf("MsgPackToJson got err %s", err.Error())
				}
				if want != -20.0 {
					t.Errorf("MsgPackToJson got %v, want %v", want, -20.0)
				}
			})
		},
		func() {
			data := []byte{0xCf, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}
			var want float32
			err := MsgPackToJson(data, &want)
			t.Run("flaot32: 3", func(t *testing.T) {
				if err != nil {
					t.Fatalf("MsgPackToJson got err %s", err.Error())
				}
				if want != 1.8446744e+19 {
					t.Errorf("MsgPackToJson got %v, want %v", want, 1.8446744e+19)
				}
			})
		},
		func() {
			data := []byte{0xca, 0x3f, 0xf1, 0xf9, 0xa6}
			var want float32
			err := MsgPackToJson(data, &want)
			t.Run("flaot32: 4", func(t *testing.T) {
				if err != nil {
					t.Fatalf("MsgPackToJson got err %s", err.Error())
				}
				if want != 1.8904312 {
					t.Errorf("MsgPackToJson got %v, want %v", want, 1.8904312)
				}
			})
		},
	}

	for _, fun := range tests {
		fun()
	}
}

func TestFloat64(t *testing.T) {
	tests := []f{
		func() {
			data := []byte{0xff}
			var want float64
			err := MsgPackToJson(data, &want)
			t.Run("float64: 1", func(t *testing.T) {
				if err != nil {
					t.Fatalf("MsgPackToJson got err %s", err.Error())
				}
				if want != -1.0 {
					t.Errorf("MsgPackToJson got %v, want %v", want, -1.0)
				}
			})
		},
		func() {
			data := []byte{0xca, 0x3f, 0xf1, 0xf9, 0xa6}
			var want float64
			err := MsgPackToJson(data, &want)
			t.Run("float64: 2", func(t *testing.T) {
				if err != nil {
					t.Fatalf("MsgPackToJson got err %s", err.Error())
				}
				if want != 1.8904311656951904 {
					t.Errorf("MsgPackToJson got %v, want %v", want, 1.8904311656951904)
				}
			})
		},
		func() {
			data := []byte{0xcb, 0x3f, 0xf1, 0xf9, 0xa6, 0xb6, 0xc6, 0xd6, 0xf6}
			var want float64
			err := MsgPackToJson(data, &want)
			t.Run("float64: 3", func(t *testing.T) {
				if err != nil {
					t.Fatalf("MsgPackToJson got err %s", err.Error())
				}
				if want != 1.123450006457856 {
					t.Errorf("MsgPackToJson got %v, want %v", want, 1.123450006457856)
				}
			})
		},
	}

	for _, fun := range tests {
		fun()
	}
}

func TestString(t *testing.T) {
	tests := []f{
		func() {
			data := []byte{0xa0}
			var want string
			err := MsgPackToJson(data, &want)
			t.Run("fixstr: 1", func(t *testing.T) {
				if err != nil {
					t.Fatalf("MsgPackToJson got err %s", err.Error())
				}
				if want != "" {
					t.Errorf("MsgPackToJson got %v, want %v", want, "")
				}
			})
		},
		func() {
			data := []byte{0xa1, 0x61}
			var want string
			err := MsgPackToJson(data, &want)
			t.Run("fixstr: 2", func(t *testing.T) {
				if err != nil {
					t.Fatalf("MsgPackToJson got err %s", err.Error())
				}
				if want != "a" {
					t.Errorf("MsgPackToJson got %v, want %v", want, "a")
				}
			})
		},
		func() {
			data := []byte{0xbf, 0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x30, 0x31, 0x32,
				0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39,
				0x30}
			var want string
			err := MsgPackToJson(data, &want)
			t.Run("fixstr: 3", func(t *testing.T) {
				if err != nil {
					t.Fatalf("MsgPackToJson got err %s", err.Error())
				}
				if want != "0123456789012345678901234567890" {
					t.Errorf("MsgPackToJson got %v, want %v", want, "0123456789012345678901234567890")
				}
			})
		},
		func() {
			data := []byte{0xd9, 0x24, 0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x30, 0x31, 0x32,
				0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39,
				0x30, 0x31, 0x32, 0x33, 0x34, 0x35}
			var want string
			err := MsgPackToJson(data, &want)
			t.Run("str8: 1", func(t *testing.T) {
				if err != nil {
					t.Fatalf("MsgPackToJson got err %s", err.Error())
				}
				if want != "012345678901234567890123456789012345" {
					t.Errorf("MsgPackToJson got %v, want %v", want, "012345678901234567890123456789012345")
				}
			})
		},
	}

	for _, fun := range tests {
		fun()
	}
}

func TestArray(t *testing.T) {
	tests := []f{
		func() {
			data := []byte{0x92, 0x01, 0x02}
			var got [2]int
			want := [2]int{1, 2}
			err := MsgPackToJson(data, &got)
			t.Run("fixarray: 1", func(t *testing.T) {
				if err != nil {
					t.Fatalf("MsgPackToJson got err %s", err.Error())
				}
				if !reflect.DeepEqual(got, want) {
					t.Errorf("MsgPackToJson got %v, want %v", got, want)
				}
			})
		},
		func() {
			data := []byte{0x92, 0xa1, 0x61, 0xa1, 0x62}
			var got [2]string
			want := [2]string{"a", "b"}
			err := MsgPackToJson(data, &got)
			t.Run("fixarray: 2", func(t *testing.T) {
				if err != nil {
					t.Fatalf("MsgPackToJson got err %s", err.Error())
				}
				if !reflect.DeepEqual(got, want) {
					t.Errorf("MsgPackToJson got %v, want %v", got, want)
				}
			})
		},
		func() {
			data := []byte{0x92, 0xCa, 0x3F, 0xF8, 0x00, 0x00, 0xCa, 0xBF, 0xF8, 0x00, 0x00}
			var got [2]float32
			want := [2]float32{1.9375, -1.9375}
			err := MsgPackToJson(data, &got)
			t.Run("fixarray: 3", func(t *testing.T) {
				if err != nil {
					t.Fatalf("MsgPackToJson got err %s", err.Error())
				}
				if !reflect.DeepEqual(got, want) {
					t.Errorf("MsgPackToJson got %v, want %v", got, want)
				}
			})
		},
		// func() {
		// 	data := []byte{0x92, 0xc3, 0x01}
		// 	var got [2]interface{}
		// 	want := [2]interface{}{true, uint64(1)}
		// 	err := MsgPackToJson(data, &got)
		// 	t.Run("fixarray: 2", func(t *testing.T) {
		// 		if err != nil {
		// 			t.Fatalf("MsgPackToJson got err %s", err.Error())
		// 		}
		// 		if !reflect.DeepEqual(got, want) {
		// 			t.Errorf("MsgPackToJson got %v, want %v", got, want)
		// 		}
		// 	})
		// },
		// func() {
		// 	data := []byte{0x92, 0xc2, 0xff}
		// 	var got [2]interface{}
		// 	want := [2]interface{}{false, int64(-1)}
		// 	err := MsgPackToJson(data, &got)
		// 	t.Run("fixarray: 3", func(t *testing.T) {
		// 		if err != nil {
		// 			t.Fatalf("MsgPackToJson got err %s", err.Error())
		// 		}
		// 		if !reflect.DeepEqual(got, want) {
		// 			t.Errorf("MsgPackToJson got %v, want %v", got, want)
		// 		}
		// 	})
		// },
		// func() {
		// 	data := []byte{0x92, 0xa1, 0x61, 0xca, 0x40, 0xc8, 0x1c, 0xc5}
		// 	var got [2]interface{}
		// 	want := [2]interface{}{"a", float32(6.253511905670166)}
		// 	err := MsgPackToJson(data, &got)
		// 	t.Run("fixarray: 4", func(t *testing.T) {
		// 		if err != nil {
		// 			t.Fatalf("MsgPackToJson got err %s", err.Error())
		// 		}
		// 		if !reflect.DeepEqual(got, want) {
		// 			t.Errorf("MsgPackToJson got %v, want %v", got, want)
		// 		}
		// 	})
		// },
		// func() {
		// 	data := []byte{0x92, 0x92, 0xc2, 0xe0, 0x92, 0xA1, 0x7A, 0xCB, 0x40, 0xC8, 0x1C, 0xC5, 0xC8, 0xA7, 0x12, 0x47}
		// 	var got [2][2]interface{}
		// 	want := [2][2]interface{}{{false, int64(-32)}, {"z", float64(12345.545185932087)}}
		// 	err := MsgPackToJson(data, &got)
		// 	t.Run("fixarray: 5", func(t *testing.T) {
		// 		if err != nil {
		// 			t.Fatalf("MsgPackToJson got err %s", err.Error())
		// 		}
		// 		if !reflect.DeepEqual(got, want) {
		// 			t.Errorf("MsgPackToJson got %v, want %v", got, want)
		// 		}
		// 	})
		// },
		// // func() {
		// // 	data := []byte{0x92, 0x92, 0xC3, 0xA1, 0x61, 0xCB, 0x40, 0x88, 0x4C, 0x5C, 0x8A, 0x71, 0x24, 0x6C}
		// // 	var got [2]interface{}
		// // 	want := [2]interface{}{[2]interface{}{true, "a"}, float64(777.545185932087)}
		// // 	err := MsgPackToJson(data, &got)
		// // 	t.Run("fixarray: 6", func(t *testing.T) {
		// // 		if err != nil {
		// // 			t.Fatalf("MsgPackToJson got err %s", err.Error())
		// // 		}
		// // 		if !reflect.DeepEqual(got, want) {
		// // 			t.Errorf("MsgPackToJson got %v, want %v", got, want)
		// // 		}
		// // 	})
		// // },
		// func() {
		// 	data := []byte{0xDC, 0x00, 0x10, 0xC2, 0xE0, 0xA1, 0x7A, 0xCB, 0x40, 0xC8, 0x1C, 0xC5, 0xC8, 0xA7, 0x12, 0x47, 0xA1, 0x7A, 0xA1,
		// 		0x79, 0xA1, 0x77, 0xA1, 0x78, 0xA1, 0x66, 0xA1, 0x74, 0xC3, 0x7B, 0xA3, 0x61, 0x62, 0x63, 0xA1, 0x61, 0xA1, 0x63, 0xA1, 0x76}
		// 	var got [16]interface{}
		// 	want := [16]interface{}{false, int64(-32), "z", float64(12345.545185932087), "z", "y", "w", "x", "f",
		// 		"t", true, uint64(123), "abc", "a", "c", "v"}
		// 	err := MsgPackToJson(data, &got)
		// 	t.Run("array16: 1", func(t *testing.T) {
		// 		if err != nil {
		// 			t.Fatalf("MsgPackToJson got err %s", err.Error())
		// 		}
		// 		if !reflect.DeepEqual(got, want) {
		// 			t.Errorf("MsgPackToJson got %v, want %v", got, want)
		// 		}
		// 	})
		// },
	}

	for _, fun := range tests {
		fun()
	}
}

func TestSlice(t *testing.T) {
	tests := []f{
		func() {
			data := []byte{0x92, 0x01, 0x02}
			var got []int
			want := []int{1, 2}
			err := MsgPackToJson(data, &got)
			t.Run("slice: 1", func(t *testing.T) {
				if err != nil {
					t.Fatalf("MsgPackToJson got err %s", err.Error())
				}
				if !reflect.DeepEqual(got, want) {
					t.Errorf("MsgPackToJson got %v, want %v", got, want)
				}
			})
		},
		func() {
			data := []byte{0x92, 0xc3, 0xc2}
			var got []bool
			want := []bool{true, false}
			err := MsgPackToJson(data, &got)
			t.Run("slice: 2", func(t *testing.T) {
				if err != nil {
					t.Fatalf("MsgPackToJson got err %s", err.Error())
				}
				if !reflect.DeepEqual(got, want) {
					t.Errorf("MsgPackToJson got %v, want %v", got, want)
				}
			})
		},
	}

	for _, fun := range tests {
		fun()
	}
}

func TestMap(t *testing.T) {
	tests := []f{
		func() {
			data := []byte{0x81, 0xa1, 0x61, 0xa4, 0x30, 0x78, 0x36, 0x31}
			var got map[string]string
			want := map[string]string{
				"a": "0x61",
			}
			err := MsgPackToJson(data, &got)
			t.Run("map: 1", func(t *testing.T) {
				if err != nil {
					t.Fatalf("MsgPackToJson got err %s", err.Error())
				}
				if !reflect.DeepEqual(got, want) {
					t.Errorf("MsgPackToJson got %v, want %v", got, want)
				}
			})
		},
		func() {
			data := []byte{0x82, 0xa1, 0x54, 0xc3, 0xa1, 0x46, 0xc2}
			var got map[string]bool
			want := map[string]bool{
				"T": true,
				"F": false,
			}
			err := MsgPackToJson(data, &got)
			t.Run("slice: 2", func(t *testing.T) {
				if err != nil {
					t.Fatalf("MsgPackToJson got err %s", err.Error())
				}
				if !reflect.DeepEqual(got, want) {
					t.Errorf("MsgPackToJson got %v, want %v", got, want)
				}
			})
		},
		// func() {
		// 	data := []byte{0x82, 0xa1, 0x54, 0xc3, 0xa1, 0x46, 0xc2}
		// 	var got map[string]interface{}
		// 	want := map[string]interface{}{
		// 		"T": true,
		// 		"F": false,
		// 	}
		// 	err := MsgPackToJson(data, &got)
		// 	t.Run("slice: 2", func(t *testing.T) {
		// 		if err != nil {
		// 			t.Fatalf("MsgPackToJson got err %s", err.Error())
		// 		}
		// 		if !reflect.DeepEqual(got, want) {
		// 			t.Errorf("MsgPackToJson got %v, want %v", got, want)
		// 		}
		// 	})
		// },
	}

	for _, fun := range tests {
		fun()
	}
}
