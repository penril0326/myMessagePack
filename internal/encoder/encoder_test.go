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
			name: "positive fixint-1",
			args: args{
				jsonData: 8,
			},
			want:    []byte{0x08},
			wantErr: false,
		},
		{
			name: "positive fixint-2",
			args: args{
				jsonData: 127,
			},
			want:    []byte{0x7f},
			wantErr: false,
		},
		{
			name: "uint 8-1",
			args: args{
				jsonData: uint8(8),
			},
			want:    []byte{0x08},
			wantErr: false,
		},
		{
			name: "uint 8-2",
			args: args{
				jsonData: uint8(127),
			},
			want:    []byte{0x7f},
			wantErr: false,
		},
		{
			name: "uint 8-3",
			args: args{
				jsonData: uint8(255),
			},
			want:    []byte{0xcc, 0xff},
			wantErr: false,
		},
		{
			name: "uint 16-1",
			args: args{
				jsonData: uint16(256),
			},
			want:    []byte{0xcd, 0x01, 0x00},
			wantErr: false,
		},
		{
			name: "uint 16-2",
			args: args{
				jsonData: uint16(65535),
			},
			want:    []byte{0xcd, 0xff, 0xff},
			wantErr: false,
		},
		{
			name: "uint 32-1",
			args: args{
				jsonData: uint32(65536),
			},
			want:    []byte{0xce, 0x00, 0x01, 0x00, 0x00},
			wantErr: false,
		},
		{
			name: "uint 32-2",
			args: args{
				jsonData: uint32(4294967295),
			},
			want:    []byte{0xce, 0xff, 0xff, 0xff, 0xff},
			wantErr: false,
		},
		{
			name: "uint 64-1",
			args: args{
				jsonData: uint64(4294967296),
			},
			want:    []byte{0xcf, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00},
			wantErr: false,
		},
		{
			name: "uint 64-2",
			args: args{
				jsonData: uint64(18446744073709551615),
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
