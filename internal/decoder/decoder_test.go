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
