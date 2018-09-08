package fallocate

import (
	"testing"
)

func TestDarwinFallocate(t *testing.T) {
	sizes := []int64{1, 99, 34, 65536, 88888}
	for _, size := range sizes {
		fallocateWithNewFile(t, size)
	}
}
