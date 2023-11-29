package definition

const (
	Nil byte = 0xc0

	False byte = 0xc2
	True  byte = 0xc3

	Float32 byte = 0xca
	Float64 byte = 0xcb

	Uint8  byte = 0xcc
	Uint16 byte = 0xcd
	Uint32 byte = 0xce
	Uint64 byte = 0xcf

	Int8  byte = 0xd0
	Int16 byte = 0xd1
	Int32 byte = 0xd2
	Int64 byte = 0xd3

	Str8  byte = 0xd9
	Str16 byte = 0xda
	Str32 byte = 0xdb
)