package union

import "unsafe"

type (
	U4  [4]byte
	U8  [8]byte
	U16 [16]byte
	U32 [32]byte
	U48 [48]byte
	U64 [64]byte
	U80 [80]byte
	U96 [96]byte
)

type unionSize interface {
	U4 | U8 | U16 | U32 | U48 | U64 | U80 | U96
}

type U[size unionSize] struct {
	v size
}

func Get[out comparable, size unionSize](u *U[size]) *out { return (*out)(unsafe.Pointer(&u.v[0])) }
