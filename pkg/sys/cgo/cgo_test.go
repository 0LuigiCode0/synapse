package sys

import (
	"testing"

	"github.com/0LuigiCode0/synapse/pkg/sys"
)

func TestCgo(t *testing.T) {
	sys.Call[sys.N9](addrC9, sys.IsC, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	sys.Call[sys.N12](addrC12, sys.IsC, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12)
	sys.Call[sys.N15](addrC15, sys.IsC, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15)
}
