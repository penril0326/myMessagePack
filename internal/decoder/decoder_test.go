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
