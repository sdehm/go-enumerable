package enumerable

import (
	"strconv"
	"testing"
)

func TestMapInt(t *testing.T) {
	e := New([]int{1, 2, 3})
	result := e.Map(func(i int) int { return i * 2 })
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

func TestMapString(t *testing.T) {
	e := New([]string{"a", "b", "c"})
	result := e.Map(func(s string) string { return s + s })
	expected := New([]string{"aa", "bb", "cc"})

	if len(result.values) != 3 {
		t.Errorf("Expected 3 values, got %d", len(result.values))
	}
	for i, v := range result.values {
		if v != expected.values[i] {
			t.Errorf("Expected %s, got %s", expected.values[i], v)
		}
	}
}

func TestReduceInt(t *testing.T) {
	e := New([]int{1, 2, 3})
	result := e.Reduce(func(a, b int) int { return a + b })
	expected := 6

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestReduceString(t *testing.T) {
	e := New([]string{"a", "b", "c"})
	result := e.Reduce(func(a, b string) string { return a + b })
	expected := "abc"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestForEach(t *testing.T) {
	e := New([]int{1, 2, 3})
	var result int
	e.ForEach(func(i int) { result += i })
	expected := 6

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestTransform(t *testing.T) {
	e := New([]int{1, 2, 3})
	result := Transform(e, func(i int) string { return strconv.Itoa(i) })
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
