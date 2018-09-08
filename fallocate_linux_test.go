package fallocate

import "testing"

func TestDarwinFallocate(t *testing.T) {
	sizes := []int64{392, 4, 1, 0, 99999}
	for _, size := range sizes {
		fallocateWithNewFile(t, size)
	}
}
