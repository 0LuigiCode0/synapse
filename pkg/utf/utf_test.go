package utf

import (
	"fmt"
	"testing"
	"unicode/utf16"
	"unsafe"
)

func TestConv(t *testing.T) {
	s := "hello world привет мир 우리 반 친구들은 공부도 열심히 하고우리 반 친구들은 공부도 열심히 하고우리 반 친구들은 공부도 열심히 하고우리 반 친구들은 공부도 열심히 하고우리 반 친구들은 공부도 열심히 하고우리 반 친구들은 공부도 열심히 하고우리 반 친구들은 공부도 열심히 하고우리 반 친구들은 공부도 열심히 하고 𔓙𔓙𔓙𔓙𔓙𔓙𔓙𔓙𔓙𔓙𔓙𔓙𔓙𔓙𔓙𔓙𔓙𔓙𔓙𔓙\000"
	sr := []rune(s)
	sl := utf16.Encode(sr)

	t.Run("str num utf16", func(t *testing.T) {
		fmt.Println(StrToNum[[]uint16](s))
		fmt.Println(utf16.Encode(sr))
		fmt.Println(UtfToUft[[]uint16](sr))
	})
	t.Run("str num utf32", func(t *testing.T) {
		fmt.Println(StrToNum[[]uint32](s))
		// fmt.Println(utf32.([]rune(s)))
	})
	t.Run("num str utf16", func(t *testing.T) {
		fmt.Println(NumToStr([]uint16{'g', 'g', 'п', '리', 0xD811, 0xDCD9}))
	})
	t.Run("num str utf32", func(t *testing.T) {
		fmt.Println(NumToStr([]rune(s)))
	})
	t.Run("str ptr", func(t *testing.T) {
		fmt.Println(*StrToPtr[uint32](s))
	})
	t.Run("ptr str", func(t *testing.T) {
		fmt.Println(PtrToStr[byte](unsafe.Pointer(unsafe.StringData(s))))
		fmt.Println(PtrToStr[uint16](unsafe.Pointer(&sl[0])))
		fmt.Println(PtrToStr[rune](unsafe.Pointer(&sr[0])))
	})
}

func BenchmarkConv(b *testing.B) {
	s := "hello world привет мир 우리 반 친구들은 공부도 열심히 하고우리 반 친구들은 공부도 열심히 하고우리 반 친구들은 공부도 열심히 하고우리 반 친구들은 공부도 열심히 하고우리 반 친구들은 공부도 열심히 하고우리 반 친구들은 공부도 열심히 하고우리 반 친구들은 공부도 열심히 하고우리 반 친구들은 공부도 열심히 하고 𔓙𔓙𔓙𔓙𔓙𔓙𔓙𔓙𔓙𔓙𔓙𔓙𔓙𔓙𔓙𔓙𔓙𔓙𔓙𔓙\000"
	sp := unsafe.Pointer(unsafe.StringData(s))
	sr := []rune(s)
	sl := utf16.Encode(sr)
	// sb := []byte(s)

	b.Run("rune unf16", func(b *testing.B) {
		for b.Loop() {
			utf16.Encode(sr)
		}
	})
	b.Run("str unf16", func(b *testing.B) {
		for b.Loop() {
			utf16.Encode([]rune(s))
		}
	})
	b.Run("utf utf", func(b *testing.B) {
		for b.Loop() {
			UtfToUft[[]uint16](sr)
		}
	})
	b.Run("str num", func(b *testing.B) {
		for b.Loop() {
			StrToNum[[]uint16](s)
		}
	})
	b.Run("unf16 str", func(b *testing.B) {
		for b.Loop() {
			_ = string(utf16.Decode(sl))
		}
	})
	b.Run("num str utf32", func(b *testing.B) {
		for b.Loop() {
			NumToStr(sr)
		}
	})
	b.Run("num str utf16", func(b *testing.B) {
		for b.Loop() {
			NumToStr(sl)
		}
	})
	b.Run("str ptr", func(b *testing.B) {
		for b.Loop() {
			StrToPtr[uint16](s)
		}
	})
	b.Run("ptr str", func(b *testing.B) {
		for b.Loop() {
			PtrToStr[byte](sp)
		}
	})
	b.Run("ptr str utf16", func(b *testing.B) {
		for b.Loop() {
			PtrToStr[uint16](unsafe.Pointer(&sl[0]))
		}
	})
}

// BenchmarkConv/rune_unf16-24 				 3748804	       316.4 ns/op	     480 B/op	       1 allocs/op
// BenchmarkConv/str_unf16
// BenchmarkConv/str_unf16-24          	  958645	      1185 ns/op	    1376 B/op	       2 allocs/op
// BenchmarkConv/utf_utf
// BenchmarkConv/utf_utf-24            	 3986185	       298.5 ns/op	     896 B/op	       1 allocs/op
// BenchmarkConv/str_num
// BenchmarkConv/str_num-24            	 2766536	       430.4 ns/op	    1152 B/op	       1 allocs/op
// BenchmarkConv/unf16_str
// BenchmarkConv/unf16_str-24          	  940605	      1197 ns/op	    2368 B/op	       4 allocs/op
// BenchmarkConv/num_str_utf32
// BenchmarkConv/num_str_utf32-24      	 3274206	       364.4 ns/op	     896 B/op	       1 allocs/op
// BenchmarkConv/num_str_utf16
// BenchmarkConv/num_str_utf16-24      	 3219848	       369.4 ns/op	     704 B/op	       1 allocs/op
// BenchmarkConv/str_ptr
// BenchmarkConv/str_ptr-24            	 2753913	       434.0 ns/op	    1152 B/op	       1 allocs/op
// BenchmarkConv/ptr_str
// BenchmarkConv/ptr_str-24            	 2924223	       410.2 ns/op	     576 B/op	       1 allocs/op
// BenchmarkConv/ptr_str_utf16
// BenchmarkConv/ptr_str_utf16-24      	 2874633	       416.9 ns/op	     704 B/op	       1 allocs/op
