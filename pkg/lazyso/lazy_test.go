package lazyso

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadSO(t *testing.T) {
	dll := NewLazySO("libc.so.6")
	f := dll.NewProc("clock_gettime").Addr()

	assert.NotEqual(t, 0, f)
}
