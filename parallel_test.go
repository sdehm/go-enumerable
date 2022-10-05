package enumerable

import (
	"sync/atomic"
	"testing"
)

func TestForEachParallel(t *testing.T) {
	e := New([]int{1, 2, 3})
	var result atomic.Int32
	e.ForEachParallel(func(i int) {
		result.Add(int32(i))
	})
	if result.Load() != 6 {
		t.Errorf("Expected 6, got %d", result.Load())
	}
}
