package sys

import (
	"math"
	"os"
	"syscall"
	"testing"
	"unsafe"

	"github.com/0LuigiCode0/synapse/pkg/lazyso"
	"github.com/0LuigiCode0/synapse/pkg/utf"
	"github.com/stretchr/testify/assert"
)

func TestSys(t *testing.T) {
	libc := lazyso.NewLazySO("libc.so.6")
	addrGetpid := libc.NewProc("getpid").Addr()
	addrGetEnvironmentVariable := libc.NewProc("getenv").Addr()
	addrWrite := libc.NewProc("write").Addr()
	addrDPrintF := libc.NewProc("dprintf").Addr()

	libm := lazyso.NewLazySO("libm.so.6")
	addrPow := libm.NewProc("jn").Addr()

	t.Run("sys0", func(t *testing.T) {
		pid := syscall.Getpid()
		assert.Equal(t, pid, int(Call[N0](addrGetpid, IsC)))
		t.Log(pid)
	})
	t.Run("sys3", func(t *testing.T) {
		name := "X"

		os.Setenv("X", "hello")
		env1, _ := syscall.Getenv("X")
		char := Call[N1](addrGetEnvironmentVariable, IsC,
			uintptr(unsafe.Pointer(utf.StrToPtr[byte](name+"\000"))))
		assert.Equal(t, env1, utf.PtrToStr[byte](unsafe.Pointer(char)))
		t.Log(env1)
	})
	t.Run("sys3 float", func(t *testing.T) {
		f1, f2 := 2, 2.23
		x1 := call9(addrPow, IsC|IsF2|FOut, uintptr(f1), uintptr(math.Float64bits(f2)), 0, 0, 0, 0, 0, 0, 0)
		x3 := math.Jn(f1, f2)
		assert.Equal(t, float32(x3), float32(math.Float64frombits(uint64(x1))))
		t.Log(float32(x3))
	})
	t.Run("sys6 write", func(t *testing.T) {
		msg := "hello world\n"
		call6(addrWrite, IsC,
			uintptr(syscall.Stdout),
			uintptr(unsafe.Pointer(unsafe.StringData(msg))), uintptr(len(msg)), 0, 0, 0)
	})
	t.Run("sys6 print", func(t *testing.T) {
		msg := "hello wrld %d %d %f %d %d %d %d\012\000"
		call9(addrDPrintF, IsC|IsF5,
			uintptr(syscall.Stdout),
			uintptr(unsafe.Pointer(unsafe.StringData(msg))), 20, 40, uintptr(math.Float64bits(3.4)), 0, 52134, 434, 2)
	})
}

func BenchmarkPow(b *testing.B) {
	f1, f2 := 2.8, 2.23
	math32 := lazyso.NewLazySO("libm.so.6")
	addrPow := math32.NewProc("pow").Addr()

	b.Run("native", func(b *testing.B) {
		for b.Loop() {
			_ = math.Pow(f1, f2)
		}
	})
	b.Run("asm", func(b *testing.B) {
		for b.Loop() {
			_ = call15(addrPow, IsC|IsF1|IsF2|FOut, uintptr(math.Float64bits(f1)), uintptr(math.Float64bits(f2)), 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0)
		}
	})
	b.Run("cgo", func(b *testing.B) {
		for b.Loop() {
			_, _, _ = syscall.RawSyscall(addrPow, uintptr(math.Float64bits(f1)), uintptr(math.Float64bits(f2)), 0)
		}
	})
}
