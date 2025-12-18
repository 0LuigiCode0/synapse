package utf

import "unsafe"

type utf interface {
	~int | ~int8 | ~int16 | ~int32 |
		~uint | ~uint8 | ~uint16 | ~uint32
}

const (
	_UTF8  = 1
	_UTF16 = 2
	_UTF32 = 4
)

func StrToNum[to ~[]t1, t1 utf](in string) (out to) {
	if len(in) == 0 {
		return nil
	}

	out = make(to, len(in))

	n := utfTouft[byte, t1](unsafe.Pointer(unsafe.StringData(in)), unsafe.Pointer(&out[0]), len(in))

	return out[:n:n]
}

func NumToStr[from ~[]t1, t1 utf](in from) string {
	if len(in) == 0 {
		return ""
	}

	_sizeIn := int(unsafe.Sizeof(*(*t1)(nil)))
	k := 1
	switch _sizeIn {
	case _UTF16:
		k = 3
	case _UTF32:
		k = 4
	}

	out := make([]byte, len(in)*k)

	n := utfTouft[t1, byte](unsafe.Pointer(&in[0]), unsafe.Pointer(&out[0]), len(in))

	return unsafe.String(&out[0], n)
}

func PtrToStr[t1 utf](ptrIn unsafe.Pointer) string {
	if ptrIn == nil {
		return ""
	}
	startPtr := ptrIn
	var lIn int
	_sizeIn := int(unsafe.Sizeof(*(*t1)(nil)))

	for *(*t1)(ptrIn) != 0 {
		lIn++
		ptrIn = unsafe.Add(ptrIn, _sizeIn)
	}

	k := 1
	switch _sizeIn {
	case _UTF16:
		k = 3
	case _UTF32:
		k = 4
	}

	out := make([]byte, lIn*k)

	n := utfTouft[t1, byte](startPtr, unsafe.Pointer(&out[0]), lIn)

	return unsafe.String(&out[0], n)
}

func StrToPtr[tChar utf](in string) *tChar { return &StrToNum[[]tChar](in)[0] }

func UtfToUft[to ~[]t2, from ~[]t1, t1, t2 utf](in from) (out to) {
	if len(in) == 0 {
		return nil
	}

	_sizeIn := int(unsafe.Sizeof(*(*t1)(nil)))
	_sizeOut := int(unsafe.Sizeof(*(*t2)(nil)))
	k := 1
	if _sizeIn != _sizeOut {
		switch _sizeIn {
		case _UTF16:
			switch _sizeOut {
			case _UTF8:
				k = 3
			}
		case _UTF32:
			switch _sizeOut {
			case _UTF8:
				k = 4
			case _UTF16:
				k = 2
			}
		}
	}

	out = make(to, len(in)*k)

	n := utfTouft[t1, t2](unsafe.Pointer(&in[0]), unsafe.Pointer(&out[0]), len(in))

	return out[:n:n]
}

func a() {
	_ = StrToNum[[]uint16]("")
}

//go:nosplit
func utfTouft[t1, t2 utf](in, out unsafe.Pointer, lIn int) (n int) {
	_sizeIn := int(unsafe.Sizeof(*(*t1)(nil)))
	_sizeOut := int(unsafe.Sizeof(*(*t2)(nil)))

	if _sizeIn == _sizeOut {
		for range lIn {
			*(*t2)(unsafe.Add(out, n*_sizeOut)) = *(*t2)(unsafe.Add(in, n*_sizeIn))
			n++
		}
		return
	}

	var r uint32
	var i int
	for i < lIn {
		c := *(*t1)(unsafe.Add(in, i*_sizeIn))
		if byte(c) == 0 {
			break
		}
		shiftIn := 1
		shiftOut := 1

		switch _sizeIn {
		case _UTF8:
			switch {
			case byte(c)&0x80 == 0:
				r = uint32(c)
			case (byte(c)>>5)^0x06 == 0:
				r = (uint32(c&0x1f) << 6) | uint32(*(*t1)(unsafe.Add(in, i+1))&0x3f)
				shiftIn = 2
			case (byte(c)>>4)^0x0e == 0:
				r = (uint32(c&0x0f) << 12) | (uint32(*(*byte)(unsafe.Add(in, i+1))&0x3f) << 6) | uint32(*(*byte)(unsafe.Add(in, 2))&0x3f)
				shiftIn = 3
			default:
				r = (uint32(c&0x07) << 18) | (uint32(*(*byte)(unsafe.Add(in, i+1))&0x3f) << 12) | (uint32(*(*byte)(unsafe.Add(in, i+2))&0x3f) << 6) | uint32(*(*byte)(unsafe.Add(in, i+3))&0x3f)
				shiftIn = 4
			}
		case _UTF16:
			if byte(uint16(c)>>11)^0x1b == 0 {
				r = uint32(uint16(c)&0x07ff)<<10 | uint32(*(*uint16)(unsafe.Add(in, (i+1)*_sizeIn))&0x03ff) | 0x10000
				shiftIn = 2
			} else {
				r = uint32(c)
			}
		case _UTF32:
			r = uint32(c)
		}

		switch _sizeOut {
		case _UTF8:
			switch {
			case byte(uint16(r)>>8)&0xff == 0:
				*(*byte)(unsafe.Add(out, n)) = byte(r) & 0x7f
			case byte(uint16(r)>>8)&0xf8 == 0:
				*(*byte)(unsafe.Add(out, n)) = 0xc0 | (byte(uint16(r)>>6) & 0x1f)
				*(*byte)(unsafe.Add(out, n+1)) = 0x80 | (byte(r) & 0x3f)
				shiftOut = 2
			case byte(uint32(r)>>16)&0xff == 0:
				*(*byte)(unsafe.Add(out, n)) = 0xe0 | (byte(uint16(r)>>12) & 0x0f)
				*(*byte)(unsafe.Add(out, n+1)) = 0x80 | (byte(uint16(r)>>6) & 0x3f)
				*(*byte)(unsafe.Add(out, n+2)) = 0x80 | (byte(r) & 0x3f)
				shiftOut = 3
			default:
				*(*byte)(unsafe.Add(out, n)) = 0xf0 | (byte(r>>18) & 0x07)
				*(*byte)(unsafe.Add(out, n+1)) = 0x80 | (byte(r>>12) & 0x3f)
				*(*byte)(unsafe.Add(out, n+2)) = 0x80 | (byte(uint16(r)>>6) & 0x3f)
				*(*byte)(unsafe.Add(out, n+3)) = 0x80 | (byte(r) & 0x3f)
				shiftOut = 4
			}
		case _UTF16:
			if byte(r>>16)&0x01 != 0 {
				r &= 0xfeffff
				*(*t2)(unsafe.Add(out, n*_sizeOut)) = t2(uint16(0xd800) | uint16(r>>10))
				*(*t2)(unsafe.Add(out, (n+1)*_sizeOut)) = t2(uint16(0xdc00) | uint16(r&0x03ff))
				shiftOut = 2
			} else {
				*(*t2)(unsafe.Add(out, n*_sizeOut)) = t2(r)
			}
		case _UTF32:
			*(*t2)(unsafe.Add(out, n*_sizeOut)) = t2(r)
		}

		i += shiftIn
		n += shiftOut
	}
	return
}
