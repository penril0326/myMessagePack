package encoder

import (
	"reflect"
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
