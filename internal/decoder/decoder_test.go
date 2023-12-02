package decoder

import "testing"

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
