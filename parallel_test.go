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

	if len(result.getValues()) != 3 {
		t.Errorf("Expected 3 values, got %d", len(result.getValues()))
	}
	for i, v := range result.getValues() {
		if v != expected.getValues()[i] {
			t.Errorf("Expected %d, got %d", expected.getValues()[i], v)
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

	if len(result.getValues()) != 3 {
		t.Errorf("Expected 3 values, got %d", len(result.getValues()))
	}
	for i, v := range result.getValues() {
		if v != expected.getValues()[i] {
			t.Errorf("Expected %s, got %s", expected.getValues()[i], v)
		}
	}
}

func TestSetNumWorkers(t *testing.T) {
	workers := setNumWorkers(1)
	if workers != 1 {
		t.Errorf("Expected 1, got %d", workers)
	}
}
