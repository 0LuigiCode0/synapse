package sys

const (
	BadCall = ^uintptr(0)
)

type mask uintptr

//go:noescape
//go:norace
func call3(f uintptr, mask mask, x1, x2, x3 uintptr) uintptr

//go:noescape
//go:norace
func call6(f uintptr, mask mask, x1, x2, x3, x4, x5, x6 uintptr) uintptr

//go:noescape
//go:norace
func call9(f uintptr, mask mask, x1, x2, x3, x4, x5, x6, x7, x8, x9 uintptr) uintptr

//go:noescape
//go:norace
func call12(f uintptr, mask mask, x1, x2, x3, x4, x5, x6, x7, x8, x9, x10, x11, x12 uintptr) uintptr

//go:noescape
//go:norace
func call15(f uintptr, mask mask, x1, x2, x3, x4, x5, x6, x7, x8, x9, x10, x11, x12, x13, x14, x15 uintptr) uintptr

type (
	N0  [0]struct{}
	N1  [1]struct{}
	N2  [2]struct{}
	N3  [3]struct{}
	N4  [4]struct{}
	N5  [5]struct{}
	N6  [6]struct{}
	N7  [7]struct{}
	N8  [8]struct{}
	N9  [9]struct{}
	N10 [10]struct{}
	N11 [11]struct{}
	N12 [12]struct{}
	N13 [13]struct{}
	N14 [14]struct{}
	N15 [15]struct{}
)

type size interface {
	N0 | N1 | N2 | N3 | N4 | N5 | N6 | N7 | N8 | N9 | N10 | N11 | N12 | N13 | N14 | N15
}

//go:norace
func Call[n size](f uintptr, mask mask, args ...uintptr) uintptr {
	var _n n
	switch len(_n) {
	case 0:
		return call3(f, mask, 0, 0, 0)
	case 1:
		return call3(f, mask, args[0], 0, 0)
	case 2:
		return call3(f, mask, args[0], args[1], 0)
	case 3:
		return call3(f, mask, args[0], args[1], args[2])
	case 4:
		return call6(f, mask, args[0], args[1], args[2], args[3], 0, 0)
	case 5:
		return call6(f, mask, args[0], args[1], args[2], args[3], args[4], 0)
	case 6:
		return call6(f, mask, args[0], args[1], args[2], args[3], args[4], args[5])
	case 7:
		return call9(f, mask, args[0], args[1], args[2], args[3], args[4], args[5], args[6], 0, 0)
	case 8:
		return call9(f, mask, args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], 0)
	case 9:
		return call9(f, mask, args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8])
	case 10:
		return call12(f, mask, args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8], args[9], 0, 0)
	case 11:
		return call12(f, mask, args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8], args[9], args[10], 0)
	case 12:
		return call12(f, mask, args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8], args[9], args[10], args[11])
	case 13:
		return call15(f, mask, args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8], args[9], args[10], args[11], args[12], 0, 0)
	case 14:
		return call15(f, mask, args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8], args[9], args[10], args[11], args[12], args[13], 0)
	case 15:
		return call15(f, mask, args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8], args[9], args[10], args[11], args[12], args[13], args[14])
	}
	return BadCall
}
