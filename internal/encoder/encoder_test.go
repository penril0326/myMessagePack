package encoder

import (
	"math"
	"reflect"
	"sort"
	"testing"
)

func TestNil(t *testing.T) {
	type args struct {
		jsonData interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "nil case",
			args: args{
				jsonData: nil,
			},
			want:    []byte{0xc0},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := JsonToMsgPack(tt.args.jsonData)
			if (err != nil) != tt.wantErr {
				t.Errorf("JsonToMsgPack() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonToMsgPack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoolean(t *testing.T) {
	type args struct {
		jsonData interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "boolean true",
			args: args{
				jsonData: true,
			},
			want:    []byte{0xc3},
			wantErr: false,
		},
		{
			name: "boolean false",
			args: args{
				jsonData: false,
			},
			want:    []byte{0xc2},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := JsonToMsgPack(tt.args.jsonData)
			if (err != nil) != tt.wantErr {
				t.Errorf("JsonToMsgPack() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonToMsgPack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint(t *testing.T) {
	type args struct {
		jsonData interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "positive fixint: 1",
			args: args{
				jsonData: uint(8),
			},
			want:    []byte{0x08},
			wantErr: false,
		},
		{
			name: "positive fixint: 2",
			args: args{
				jsonData: uint(127),
			},
			want:    []byte{0x7f},
			wantErr: false,
		},
		{
			name: "uint8: 1",
			args: args{
				jsonData: uint8(8),
			},
			want:    []byte{0x08},
			wantErr: false,
		},
		{
			name: "uint8: 2",
			args: args{
				jsonData: uint8(127),
			},
			want:    []byte{0x7f},
			wantErr: false,
		},
		{
			name: "uint8: 3",
			args: args{
				jsonData: uint8(255),
			},
			want:    []byte{0xcc, 0xff},
			wantErr: false,
		},
		{
			name: "uint16: 1",
			args: args{
				jsonData: uint16(256),
			},
			want:    []byte{0xcd, 0x01, 0x00},
			wantErr: false,
		},
		{
			name: "uint16: 2",
			args: args{
				jsonData: uint16(65535),
			},
			want:    []byte{0xcd, 0xff, 0xff},
			wantErr: false,
		},
		{
			name: "uint32: 1",
			args: args{
				jsonData: uint32(65536),
			},
			want:    []byte{0xce, 0x00, 0x01, 0x00, 0x00},
			wantErr: false,
		},
		{
			name: "uint32: 2",
			args: args{
				jsonData: uint32(4294967295),
			},
			want:    []byte{0xce, 0xff, 0xff, 0xff, 0xff},
			wantErr: false,
		},
		{
			name: "uint64: 1",
			args: args{
				jsonData: uint64(4294967296),
			},
			want:    []byte{0xcf, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00},
			wantErr: false,
		},
		{
			name: "uint64: 2",
			args: args{
				jsonData: uint64(18446744073709551615),
			},
			want:    []byte{0xcf, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
			wantErr: false,
		},
		{
			name: "uint: 1",
			args: args{
				jsonData: uint(4294967296),
			},
			want:    []byte{0xcf, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00},
			wantErr: false,
		},
		{
			name: "uint: 2",
			args: args{
				jsonData: uint(18446744073709551615),
			},
			want:    []byte{0xcf, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := JsonToMsgPack(tt.args.jsonData)
			if (err != nil) != tt.wantErr {
				t.Errorf("JsonToMsgPack() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonToMsgPack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt(t *testing.T) {
	type args struct {
		jsonData interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "int8: 1",
			args: args{
				jsonData: int8(8),
			},
			want:    []byte{0x08},
			wantErr: false,
		},
		{
			name: "int8: 2",
			args: args{
				jsonData: int8(-127),
			},
			want:    []byte{0xd0, 0x81},
			wantErr: false,
		},
		{
			name: "int8: 3",
			args: args{
				jsonData: int8(-128),
			},
			want:    []byte{0xd0, 0x80},
			wantErr: false,
		},
		{
			name: "int16: 1",
			args: args{
				jsonData: int16(-129),
			},
			want:    []byte{0xd1, 0xff, 0x7f},
			wantErr: false,
		},
		{
			name: "int16: 2",
			args: args{
				jsonData: int16(-32768),
			},
			want:    []byte{0xd1, 0x80, 0x00},
			wantErr: false,
		},
		{
			name: "int32: 1",
			args: args{
				jsonData: int32(-32769),
			},
			want:    []byte{0xd2, 0xff, 0xff, 0x7f, 0xff},
			wantErr: false,
		},
		{
			name: "int32: 2",
			args: args{
				jsonData: int32(-2147483648),
			},
			want:    []byte{0xd2, 0x80, 0x00, 0x00, 0x00},
			wantErr: false,
		},
		{
			name: "int64: 1",
			args: args{
				jsonData: int64(-2147483649),
			},
			want:    []byte{0xd3, 0xff, 0xff, 0xff, 0xff, 0x7f, 0xff, 0xff, 0xff},
			wantErr: false,
		},
		{
			name: "int64: 2",
			args: args{
				jsonData: int64(-9223372036854775808),
			},
			want:    []byte{0xd3, 0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			wantErr: false,
		},
		{
			name: "constant: 1",
			args: args{
				jsonData: 1,
			},
			want:    []byte{0x01},
			wantErr: false,
		},
		{
			name: "constant: 2",
			args: args{
				jsonData: -128,
			},
			want:    []byte{0xd0, 0x80},
			wantErr: false,
		},
		{
			name: "constant: 3",
			args: args{
				jsonData: -9223372036854775808,
			},
			want:    []byte{0xd3, 0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := JsonToMsgPack(tt.args.jsonData)
			if (err != nil) != tt.wantErr {
				t.Errorf("JsonToMsgPack() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonToMsgPack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32(t *testing.T) {
	type args struct {
		jsonData interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "float32: 1",
			args: args{
				jsonData: float32(1.5),
			},
			want:    []byte{0xca, 0x3f, 0xc0, 0x00, 0x00},
			wantErr: false,
		},
		{
			name: "float32: 2",
			args: args{
				jsonData: float32(1.1),
			},
			want:    []byte{0xca, 0x3f, 0x8c, 0xcc, 0xcd},
			wantErr: false,
		},
		{
			name: "float32: 3",
			args: args{
				jsonData: float32(math.MaxFloat32),
			},
			want:    []byte{0xca, 0x7f, 0x7f, 0xff, 0xff},
			wantErr: false,
		},
		{
			name: "float32: 4",
			args: args{
				jsonData: float32(-0.123454321),
			},
			want:    []byte{0xca, 0xbd, 0xfc, 0xd5, 0x9e},
			wantErr: false,
		},
		{
			name: "float32: 5",
			args: args{
				jsonData: float32(math.SmallestNonzeroFloat32),
			},
			want:    []byte{0xca, 0x00, 0x00, 0x00, 0x01},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := JsonToMsgPack(tt.args.jsonData)
			if (err != nil) != tt.wantErr {
				t.Errorf("JsonToMsgPack() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonToMsgPack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64(t *testing.T) {
	type args struct {
		jsonData interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "float64: 1",
			args: args{
				jsonData: float64(1.5),
			},
			want:    []byte{0xcb, 0x3f, 0xf8, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			wantErr: false,
		},
		{
			name: "float64: 2",
			args: args{
				jsonData: float64(math.MaxFloat64),
			},
			want:    []byte{0xcb, 0x7f, 0xef, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
			wantErr: false,
		},
		{
			name: "float64: 3",
			args: args{
				jsonData: float64(math.MaxFloat32),
			},
			want:    []byte{0xcb, 0x47, 0xef, 0xff, 0xff, 0xe0, 0x00, 0x00, 0x00},
			wantErr: false,
		},
		{
			name: "float64: 4",
			args: args{
				jsonData: float64(math.SmallestNonzeroFloat64),
			},
			want:    []byte{0xcb, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01},
			wantErr: false,
		},
		{
			name: "float64: 5",
			args: args{
				jsonData: float64(-0.01234567654321),
			},
			want:    []byte{0xcb, 0xbf, 0x89, 0x48, 0xb0, 0xa8, 0x00, 0x2a, 0x9d},
			wantErr: false,
		},
		{
			name: "float64: 6",
			args: args{
				jsonData: -0.88888888,
			},
			want:    []byte{0xcb, 0xbf, 0xec, 0x71, 0xc7, 0x17, 0xac, 0x19, 0x23},
			wantErr: false,
		},
		{
			name: "float64: 7",
			args: args{
				jsonData: math.MaxFloat32,
			},
			want:    []byte{0xcb, 0x47, 0xef, 0xff, 0xff, 0xe0, 0x00, 0x00, 0x00},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := JsonToMsgPack(tt.args.jsonData)
			if (err != nil) != tt.wantErr {
				t.Errorf("JsonToMsgPack() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonToMsgPack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStr(t *testing.T) {
	type args struct {
		jsonData interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "fixstr 1",
			args: args{
				jsonData: "a",
			},
			want:    []byte{0xa1, 0x61},
			wantErr: false,
		},
		{
			name: "fixstr 2",
			args: args{
				jsonData: "",
			},
			want:    []byte{0xa0},
			wantErr: false,
		},
		{
			name: "fixstr 3",
			args: args{
				jsonData: "abcjalkjsdijqpiweqpwuepcndjlnkq",
			},
			want: []byte{0xbf, 0x61, 0x62, 0x63, 0x6a, 0x61, 0x6c, 0x6b, 0x6a, 0x73,
				0x64, 0x69, 0x6a, 0x71, 0x70, 0x69, 0x77, 0x65, 0x71, 0x70, 0x77, 0x75,
				0x65, 0x70, 0x63, 0x6e, 0x64, 0x6a, 0x6c, 0x6e, 0x6b, 0x71},
			wantErr: false,
		},
		{
			name: "str 8-1",
			args: args{
				jsonData: "abcjalkjsdijqpiweqpwuepcndjlnkqa",
			},
			want: []byte{0xd9, 0x20, 0x61, 0x62, 0x63, 0x6a, 0x61, 0x6c, 0x6b, 0x6a, 0x73,
				0x64, 0x69, 0x6a, 0x71, 0x70, 0x69, 0x77, 0x65, 0x71, 0x70, 0x77, 0x75,
				0x65, 0x70, 0x63, 0x6e, 0x64, 0x6a, 0x6c, 0x6e, 0x6b, 0x71, 0x61},
			wantErr: false,
		},
		{
			name: "str 8-2",
			args: args{
				jsonData: "abcjalkjsdijqpiweqpwuepcndjlnkqaabcjalkjsdijqpiweqpwuepcndjlnkqa",
			},
			want: []byte{0xd9, 0x40, 0x61, 0x62, 0x63, 0x6a, 0x61, 0x6c, 0x6b, 0x6a, 0x73,
				0x64, 0x69, 0x6a, 0x71, 0x70, 0x69, 0x77, 0x65, 0x71, 0x70, 0x77, 0x75,
				0x65, 0x70, 0x63, 0x6e, 0x64, 0x6a, 0x6c, 0x6e, 0x6b, 0x71, 0x61,
				0x61, 0x62, 0x63, 0x6a, 0x61, 0x6c, 0x6b, 0x6a, 0x73,
				0x64, 0x69, 0x6a, 0x71, 0x70, 0x69, 0x77, 0x65, 0x71, 0x70, 0x77, 0x75,
				0x65, 0x70, 0x63, 0x6e, 0x64, 0x6a, 0x6c, 0x6e, 0x6b, 0x71, 0x61},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := JsonToMsgPack(tt.args.jsonData)
			if (err != nil) != tt.wantErr {
				t.Errorf("JsonToMsgPack() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonToMsgPack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArray(t *testing.T) {
	type args struct {
		jsonData interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "fixarray: 1",
			args: args{
				jsonData: [3]int{1, 2, 3},
			},
			want:    []byte{0x93, 0x01, 0x02, 0x03},
			wantErr: false,
		},
		{
			name: "fixarray: 2",
			args: args{
				jsonData: [3]string{"a", "b", "c"},
			},
			want:    []byte{0x93, 0xa1, 0x61, 0xa1, 0x62, 0xa1, 0x63},
			wantErr: false,
		},
		{
			name: "fixarray: 3",
			args: args{
				jsonData: [15]float32{1.1, 1.2, 1.3, 1.4, 1.5, 1.6, 1.7, 1.8, 0.0, -0.1, -0.2, -0.3, -0.4, -0.5, -0.6},
			},
			want: []byte{0x9f, 0xca, 0x3f, 0x8c, 0xcc, 0xcd, 0xca, 0x3f, 0x99, 0x99, 0x9a, 0xca, 0x3f, 0xa6, 0x66, 0x66, 0xca, 0x3f, 0xb3, 0x33, 0x33, 0xca, 0x3f, 0xc0, 0x0, 0x0, 0xca, 0x3f, 0xcc, 0xcc, 0xcd, 0xca, 0x3f, 0xd9, 0x99, 0x9a,
				0xca, 0x3f, 0xe6, 0x66, 0x66, 0xca, 0x0, 0x0, 0x0, 0x0, 0xca, 0xbd, 0xcc, 0xcc, 0xcd, 0xca, 0xbe, 0x4c, 0xcc, 0xcd, 0xca, 0xbe, 0x99, 0x99, 0x9a, 0xca, 0xbe, 0xcc, 0xcc, 0xcd, 0xca, 0xbf, 0x0, 0x0, 0x0, 0xca, 0xbf, 0x19, 0x99, 0x9a},
			wantErr: false,
		},
		{
			name: "fixarray: 4",
			args: args{
				jsonData: [3]interface{}{"a", 1, true},
			},
			want:    []byte{0x93, 0xa1, 0x61, 0x01, 0xc3},
			wantErr: false,
		},
		{
			name: "array16: 1",
			args: args{
				jsonData: [16]interface{}{false, 1, "2", 3, 4.0, -5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			},
			want: []byte{0xdc, 0x00, 0x10, 0xc2, 0x01, 0xa1, 0x32, 0x03, 0xcb, 0x40, 0x10, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0xfb, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f},
			wantErr: false,
		},
		{
			name: "fixslice: 1",
			args: args{
				jsonData: []interface{}{false, 1, "2", 3, 4.0, -5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			},
			want: []byte{0xdc, 0x00, 0x10, 0xc2, 0x01, 0xa1, 0x32, 0x03, 0xcb, 0x40, 0x10, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0xfb, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f},
			wantErr: false,
		},
		{
			name: "fixslice: 2",
			args: args{
				jsonData: []interface{}{false, 1, "2", 3, 4.0, -5, 6, 7, 8, 9, 10, 11, 12, 13, 14},
			},
			want: []byte{0x9f, 0xc2, 0x01, 0xa1, 0x32, 0x03, 0xcb, 0x40, 0x10, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0xfb, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := JsonToMsgPack(tt.args.jsonData)
			if (err != nil) != tt.wantErr {
				t.Errorf("JsonToMsgPack() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonToMsgPack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMap(t *testing.T) {
	type args struct {
		jsonData interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "fixmap: 1",
			args: args{
				jsonData: map[int]bool{
					0: true,
				},
			},
			want:    []byte{0x81, 0x00, 0xc3},
			wantErr: false,
		},
		{
			name: "fixmap: 2",
			args: args{
				jsonData: map[int]interface{}{
					0: nil,
					1: false,
				},
			},
			want:    []byte{0x82, 0x00, 0xc0, 0x01, 0xc2},
			wantErr: false,
		},
		{
			name: "fixmap: 3",
			args: args{
				jsonData: map[interface{}]interface{}{
					0:          nil,
					1:          true,
					"a":        "ccc",
					true:       1,
					false:      0,
					[1]uint{1}: "uint array",
				},
			},
			want: []byte{0x86, 0x00, 0xc0, 0x01, 0xc3, 0xa1, 0x61, 0xa3, 0x63, 0x63, 0x63, 0xc3, 0x01, 0xc2,
				0x00, 0x91, 0x01, 0xaa, 0x75, 0x69, 0x6e, 0x74, 0x20, 0x61, 0x72, 0x72, 0x61, 0x79},
			wantErr: false,
		},
		{
			name: "fixmap: 4",
			args: args{
				jsonData: map[interface{}]interface{}{
					0:                                  nil,
					1:                                  true,
					"a":                                "ccc",
					true:                               1,
					false:                              0,
					[1]uint{1}:                         "uint array",
					-20:                                "negative int",
					"-20":                              "negative int",
					"中文":                               "",
					0.5:                                0.5,
					2.0:                                "2.0",
					3:                                  "negative int",
					4:                                  "negative int",
					65536:                              "negative int",
					"abcdefghijklmnopqrsyuvwxyz123456": "string16 key",
				},
			},
			want: []byte{0x8f, 0x03, 0xac, 0x6e, 0x65, 0x67, 0x61, 0x74, 0x69, 0x76,
				0x65, 0x20, 0x69, 0x6e, 0x74, 0x00, 0xc0, 0xc3, 0x01, 0xc2,
				0x00, 0xa3, 0x2d, 0x32, 0x30, 0xac, 0x6e, 0x65, 0x67, 0x61,
				0x74, 0x69, 0x76, 0x65, 0x20, 0x69, 0x6e, 0x74, 0xcb, 0x3f,
				0xe0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xcb, 0x3f, 0xe0,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0xc3, 0xa1, 0x61,
				0xa3, 0x63, 0x63, 0x63, 0x91, 0x01, 0xaa, 0x75, 0x69, 0x6e,
				0x74, 0x20, 0x61, 0x72, 0x72, 0x61, 0x79, 0x04, 0xac, 0x6e,
				0x65, 0x67, 0x61, 0x74, 0x69, 0x76, 0x65, 0x20, 0x69, 0x6e,
				0x74, 0xec, 0xac, 0x6e, 0x65, 0x67, 0x61, 0x74, 0x69, 0x76,
				0x65, 0x20, 0x69, 0x6e, 0x74, 0xa6, 0xe4, 0xb8, 0xad, 0xe6,
				0x96, 0x87, 0xa3, 0xef, 0xa3, 0xbf, 0xcb, 0x40, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0xa3, 0x32, 0x2e, 0x30, 0xce,
				0x00, 0x01, 0x00, 0x00, 0xac, 0x6e, 0x65, 0x67, 0x61, 0x74,
				0x69, 0x76, 0x65, 0x20, 0x69, 0x6e, 0x74, 0xd9, 0x20, 0x61,
				0x62, 0x63, 0x64, 0x65, 0x66, 0x67, 0x68, 0x69, 0x6a, 0x6b,
				0x6c, 0x6d, 0x6e, 0x6f, 0x70, 0x71, 0x72, 0x73, 0x79, 0x75,
				0x76, 0x77, 0x78, 0x79, 0x7a, 0x31, 0x32, 0x33, 0x34, 0x35,
				0x36, 0xac, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x31, 0x36,
				0x20, 0x6b, 0x65, 0x79},
			wantErr: false,
		},
		{
			name: "fixmap: 5",
			args: args{
				jsonData: map[string]interface{}{
					"Null": nil,
					"nested_map": map[string]string{
						"a": "aaa",
						"z": "zzz",
					},
				},
			},
			want: []byte{0x82, 0xa4, 0x4e, 0x75, 0x6c, 0x6c, 0xc0, 0xaa, 0x6e, 0x65,
				0x73, 0x74, 0x65, 0x64, 0x5f, 0x6d, 0x61, 0x70, 0x82, 0xa1,
				0x61, 0xa3, 0x61, 0x61, 0x61, 0xa1, 0x7a, 0xa3, 0x7a, 0x7a,
				0x7a},
			wantErr: false,
		},
		{
			name: "fixmap: 6",
			args: args{
				jsonData: map[string]interface{}{
					"Null": nil,
					"nested_map": map[string]string{
						"a": "aaa",
						"z": "zzz",
					},
					"multiple_nested_map": map[int]interface{}{
						1: map[string]string{
							"double": "oops!",
						},
						2: map[int]interface{}{
							3: map[int]interface{}{
								4: map[int]int{
									100: 100,
								},
							},
						},
					},
				},
			},
			want: []byte{0x83, 0xa4, 0x4e, 0x75, 0x6c, 0x6c, 0xc0, 0xaa, 0x6e, 0x65,
				0x73, 0x74, 0x65, 0x64, 0x5f, 0x6d, 0x61, 0x70, 0x82, 0xa1,
				0x7a, 0xa3, 0x7a, 0x7a, 0x7a, 0xa1, 0x61, 0xa3, 0x61, 0x61,
				0x61, 0xb3, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65,
				0x5f, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x6d, 0x61,
				0x70, 0x82, 0x01, 0x81, 0xa6, 0x64, 0x6f, 0x75, 0x62, 0x6c,
				0x65, 0xa5, 0x6f, 0x6f, 0x70, 0x73, 0x21, 0x02, 0x81, 0x03,
				0x81, 0x04, 0x81, 0x64, 0x64},
			wantErr: false,
		},
		{
			name: "map16: 1",
			args: args{
				jsonData: map[interface{}]interface{}{
					0:                                  nil,
					1:                                  true,
					"a":                                "ccc",
					true:                               1,
					false:                              0,
					[1]uint{1}:                         "uint array",
					-20:                                "negative int",
					"-20":                              "negative int",
					"中文":                               "",
					0.5:                                0.5,
					2.0:                                "2.0",
					3:                                  "negative int",
					4:                                  "negative int",
					65536:                              "negative int",
					"abcdefghijklmnopqrsyuvwxyz123456": "string16 key",
					"wow":                              "map16",
				},
			},
			want: []byte{0xde, 0x00, 0x10, 0xa6, 0xe4, 0xb8, 0xad, 0xe6, 0x96, 0x87,
				0xa3, 0xef, 0xa3, 0xbf, 0xce, 0x00, 0x01, 0x00, 0x00, 0xac,
				0x6e, 0x65, 0x67, 0x61, 0x74, 0x69, 0x76, 0x65, 0x20, 0x69,
				0x6e, 0x74, 0xec, 0xac, 0x6e, 0x65, 0x67, 0x61, 0x74, 0x69,
				0x76, 0x65, 0x20, 0x69, 0x6e, 0x74, 0x91, 0x01, 0xaa, 0x75,
				0x69, 0x6e, 0x74, 0x20, 0x61, 0x72, 0x72, 0x61, 0x79, 0x04,
				0xac, 0x6e, 0x65, 0x67, 0x61, 0x74, 0x69, 0x76, 0x65, 0x20,
				0x69, 0x6e, 0x74, 0x01, 0xc3, 0xc3, 0x01, 0xc2, 0x00, 0xa3,
				0x2d, 0x32, 0x30, 0xac, 0x6e, 0x65, 0x67, 0x61, 0x74, 0x69,
				0x76, 0x65, 0x20, 0x69, 0x6e, 0x74, 0xcb, 0x3f, 0xe0, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0xcb, 0x3f, 0xe0, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0xcb, 0x40, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0xa3, 0x32, 0x2e, 0x30, 0x00, 0xc0, 0x03,
				0xac, 0x6e, 0x65, 0x67, 0x61, 0x74, 0x69, 0x76, 0x65, 0x20,
				0x69, 0x6e, 0x74, 0xd9, 0x20, 0x61, 0x62, 0x63, 0x64, 0x65,
				0x66, 0x67, 0x68, 0x69, 0x6a, 0x6b, 0x6c, 0x6d, 0x6e, 0x6f,
				0x70, 0x71, 0x72, 0x73, 0x79, 0x75, 0x76, 0x77, 0x78, 0x79,
				0x7a, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0xac, 0x73, 0x74,
				0x72, 0x69, 0x6e, 0x67, 0x31, 0x36, 0x20, 0x6b, 0x65, 0x79,
				0xa3, 0x77, 0x6f, 0x77, 0xa5, 0x6d, 0x61, 0x70, 0x31, 0x36,
				0xa1, 0x61, 0xa3, 0x63, 0x63, 0x63},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := JsonToMsgPack(tt.args.jsonData)
			sortbyte(got)
			sortbyte(tt.want)
			if (err != nil) != tt.wantErr {
				t.Errorf("JsonToMsgPack() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonToMsgPack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func sortbyte(b []byte) {
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
}
