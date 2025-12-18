package sys

import (
	"math"
	"os"
	"syscall"
	"testing"
	"unsafe"

	"github.com/0LuigiCode0/synapse/pkg/utf"
	"github.com/stretchr/testify/assert"
)

func TestSys(t *testing.T) {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	addrGetCurrentProcess := kernel32.NewProc("GetCurrentProcess").Addr()
	addrGetEnvironmentVariableW := kernel32.NewProc("GetEnvironmentVariableW").Addr()
	addrGetProcessTimes := kernel32.NewProc("GetProcessTimes").Addr()

	msvcrt := syscall.NewLazyDLL("msvcrt.dll")
	addrPow := msvcrt.NewProc("pow").Addr()

	t.Run("sys0", func(t *testing.T) {
		r, _ := syscall.GetCurrentProcess()
		assert.Equal(t, r, syscall.Handle(Call[N0](addrGetCurrentProcess, IsC)))
	})
	t.Run("sys3", func(t *testing.T) {
		name := utf.StrToNum[[]uint16]("X")
		var l uint32 = 128
		buf1 := make([]uint16, l)
		buf2 := make([]uint16, l)

		os.Setenv("X", "hello")
		n1, _ := syscall.GetEnvironmentVariable(&name[0], &buf1[0], l)
		n2 := uint32(call3(addrGetEnvironmentVariableW, IsC,
			uintptr(unsafe.Pointer(&name[0])),
			uintptr(unsafe.Pointer(&buf2[0])),
			uintptr(l)),
		)
		assert.Equal(t, buf1[:n1], buf2[:n2])
	})
	t.Run("sys3 float", func(t *testing.T) {
		f1, f2 := 2.8, 2.23
		x1 := call15(addrPow, IsC|IsF1|IsF2|FOut, uintptr(math.Float64bits(f1)), uintptr(math.Float64bits(f2)), 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0)
		_, x2, _ := syscall.SyscallN(addrPow, uintptr(math.Float64bits(f1)), uintptr(math.Float64bits(f2)))
		x3 := math.Pow(f1, f2)
		assert.Equal(t, x3, math.Float64frombits(uint64(x1)))
		assert.Equal(t, x3, math.Float64frombits(uint64(x2)))
	})
	t.Run("sys6", func(t *testing.T) {
		handle, _ := syscall.GetCurrentProcess()
		var creationTime1, exitTime1, kernelTime1, userTime1 syscall.Filetime
		var creationTime2, exitTime2, kernelTime2, userTime2 syscall.Filetime
		err := syscall.GetProcessTimes(handle, &creationTime1, &exitTime1, &kernelTime1, &userTime1)
		if err != nil {
			t.Fatal(err)
		}
		call6(addrGetProcessTimes, IsC,
			uintptr(handle),
			uintptr(unsafe.Pointer(&creationTime2)),
			uintptr(unsafe.Pointer(&exitTime2)),
			uintptr(unsafe.Pointer(&kernelTime2)),
			uintptr(unsafe.Pointer(&userTime2)), 0)
		assert.Equal(t, creationTime1, creationTime2)
		assert.Equal(t, exitTime1, exitTime2)
		assert.Equal(t, kernelTime1, kernelTime2)
		assert.Equal(t, userTime1, userTime2)
	})
}

func BenchmarkPow(b *testing.B) {
	f1, f2 := 2.8, 2.23
	math32 := syscall.NewLazyDLL("msvcrt.dll")
	addrPow := math32.NewProc("pow").Addr()

	b.Run("native", func(b *testing.B) {
		for b.Loop() {
			_ = math.Pow(f1, f2)
		}
	})
	b.Run("asm", func(b *testing.B) {
		for b.Loop() {
			_ = call3(addrPow, IsC|IsF1|IsF2|FOut, uintptr(math.Float64bits(f1)), uintptr(math.Float64bits(f2)), 0)
		}
	})
	b.Run("cgo", func(b *testing.B) {
		for b.Loop() {
			_, _, _ = syscall.SyscallN(addrPow, uintptr(math.Float64bits(f1)), uintptr(math.Float64bits(f2)))
		}
	})
}
