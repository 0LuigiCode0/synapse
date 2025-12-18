package union

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

func TestUnion(t *testing.T) {
	type (
		a struct {
			x int64
		}
		b struct {
			arr [4]byte
		}
		c struct {
			f float64
		}
		test struct {
			v U[U8]
		}
	)
	t.Run("int", func(t *testing.T) {
		_a := a{x: 12345678}
		_test := test{}
		copy(unsafe.Slice((*byte)(unsafe.Pointer(&_test.v.v[0])), 8), unsafe.Slice((*byte)(unsafe.Pointer(&_a)), unsafe.Sizeof(_a)))
		assert.Equal(t, _a, *Get[a](&_test.v))
	})
	t.Run("arr", func(t *testing.T) {
		_b := b{arr: [4]byte{1, 2, 3, 4}}
		_test := test{}
		copy(unsafe.Slice((*byte)(unsafe.Pointer(&_test.v.v[0])), 8), unsafe.Slice((*byte)(unsafe.Pointer(&_b)), unsafe.Sizeof(_b)))
		assert.Equal(t, _b, *Get[b](&_test.v))
	})
	t.Run("float", func(t *testing.T) {
		_c := c{f: 12345678.322}
		_test := test{}
		copy(unsafe.Slice((*byte)(unsafe.Pointer(&_test.v.v[0])), 8), unsafe.Slice((*byte)(unsafe.Pointer(&_c)), unsafe.Sizeof(_c)))
		assert.Equal(t, _c, *Get[c](&_test.v))
	})
}
