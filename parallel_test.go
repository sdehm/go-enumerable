package enumerable

import (
	"strconv"
	"sync/atomic"
	"testing"
	"time"
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

func TestMapParallel(t *testing.T) {
	e := New([]int{1, 2, 3})
	result := e.MapParallel(func(i int) int {
		if i == 2 {
			// sleep for 10 ms to make sure order is maintained
			time.Sleep(time.Millisecond * 10)
		}
		return i * 2
	}).Apply()
	expected := New([]int{2, 4, 6})

	if len(result.values) != 3 {
		t.Errorf("Expected 3 values, got %d", len(result.values))
	}
	for i, v := range result.values {
		if v != expected.values[i] {
			t.Errorf("Expected %d, got %d", expected.values[i], v)
		}
	}
}

func TestTransformParallel(t *testing.T) {
	e := New([]int{1, 2, 3})
	result := TransformParallel(e, func(i int) string {
		if i == 2 {
			// sleep for 10 ms to make sure order is maintained
			time.Sleep(time.Millisecond * 10)
		}
		return strconv.Itoa(i)
	}).Apply()
	expected := New([]string{"1", "2", "3"})

	if len(result.values) != 3 {
		t.Errorf("Expected 3 values, got %d", len(result.values))
	}
	for i, v := range result.values {
		if v != expected.values[i] {
			t.Errorf("Expected %s, got %s", expected.values[i], v)
		}
	}
}
