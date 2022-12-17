package enumerable

import (
	"strconv"
	"testing"
)

func TestAppend(t *testing.T) {
	e := New([]int{1, 2, 3})
	result := e.Append(4).Apply()
	expected := New([]int{1, 2, 3, 4})

	if len(result.getValues()) != 4 {
		t.Errorf("Expected 4 values, got %d", len(result.getValues()))
	}
	for i, v := range result.getValues() {
		if v != expected.getValues()[i] {
			t.Errorf("Expected %d, got %d", expected.getValues()[i], v)
		}
	}
}

func TestMapInt(t *testing.T) {
	e := New([]int{1, 2, 3})
	result := e.Map(func(i int) int { return i * 2 }).Apply()
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

func TestMapString(t *testing.T) {
	e := New([]string{"a", "b", "c"})
	result := e.Map(func(s string) string { return s + s }).Apply()
	expected := New([]string{"aa", "bb", "cc"})

	if len(result.getValues()) != 3 {
		t.Errorf("Expected 3 values, got %d", len(result.getValues()))
	}
	for i, v := range result.getValues() {
		if v != expected.getValues()[i] {
			t.Errorf("Expected %s, got %s", expected.getValues()[i], v)
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

	if len(result.getValues()) != 3 {
		t.Errorf("Expected 3 values, got %d", len(result.getValues()))
	}
	for i, v := range result.getValues() {
		if v != expected.getValues()[i] {
			t.Errorf("Expected %s, got %s", expected.getValues()[i], v)
		}
	}
}

func TestReverse(t *testing.T) {
	e := New([]int{1, 2, 3})
	result := e.Reverse().Apply()
	expected := New([]int{3, 2, 1})

	if len(result.getValues()) != 3 {
		t.Errorf("Expected 3 values, got %d", len(result.getValues()))
	}
	for i, v := range result.getValues() {
		if v != expected.getValues()[i] {
			t.Errorf("Expected %d, got %d", expected.getValues()[i], v)
		}
	}
}

func TestFilter(t *testing.T) {
	e := New([]int{1, 2, 3})
	result := e.Filter(func(i int) bool { return i > 1 }).Apply()
	expected := New([]int{2, 3})

	if len(result.getValues()) != 2 {
		t.Errorf("Expected 2 values, got %d", len(result.getValues()))
	}
	for i, v := range result.getValues() {
		if v != expected.getValues()[i] {
			t.Errorf("Expected %d, got %d", expected.getValues()[i], v)
		}
	}
}

func TestToList(t *testing.T) {
	e := New([]int{1, 2, 3})
	result := e.ToList()
	expected := []int{1, 2, 3}

	if len(result) != 3 {
		t.Errorf("Expected 3 values, got %d", len(result))
	}
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], v)
		}
	}
}

func TestMapThenFilter(t *testing.T) {
	e := New([]int{1, 2, 3})
	result := e.Map(func(i int) int { return i * 2 }).Filter(func(i int) bool { return i > 2 }).Apply()
	expected := New([]int{4, 6})

	if len(result.getValues()) != 2 {
		t.Errorf("Expected 2 values, got %d", len(result.getValues()))
	}
	for i, v := range result.getValues() {
		if v != expected.getValues()[i] {
			t.Errorf("Expected %d, got %d", expected.getValues()[i], v)
		}
	}
}

func TestFilterThenMap(t *testing.T) {
	e := New([]int{1, 2, 3})
	result := e.Filter(func(i int) bool { return i > 1 }).Map(func(i int) int { return i * 2 }).Apply()
	expected := New([]int{4, 6})

	if len(result.getValues()) != 2 {
		t.Errorf("Expected 2 values, got %d", len(result.getValues()))
	}
	for i, v := range result.getValues() {
		if v != expected.getValues()[i] {
			t.Errorf("Expected %d, got %d", expected.getValues()[i], v)
		}
	}
}

func TestMapThenReverse(t *testing.T) {
	e := New([]int{1, 2, 3})
	result := e.Map(func(i int) int { return i * 2 }).Reverse().Apply()
	expected := New([]int{6, 4, 2})

	if len(result.getValues()) != 3 {
		t.Errorf("Expected 3 values, got %d", len(result.getValues()))
	}
	for i, v := range result.getValues() {
		if v != expected.getValues()[i] {
			t.Errorf("Expected %d, got %d", expected.getValues()[i], v)
		}
	}
}

func TestReverseThenMap(t *testing.T) {
	e := New([]int{1, 2, 3})
	result := e.Reverse().Map(func(i int) int { return i * 2 }).Apply()
	expected := New([]int{6, 4, 2})

	if len(result.getValues()) != 3 {
		t.Errorf("Expected 3 values, got %d", len(result.getValues()))
	}
	for i, v := range result.getValues() {
		if v != expected.getValues()[i] {
			t.Errorf("Expected %d, got %d", expected.getValues()[i], v)
		}
	}
}

func TestContains(t *testing.T) {
	e := New([]int{1, 2, 3})

	if !e.Contains(1) {
		t.Errorf("Expected true, got false")
	}
	if e.Contains(4) {
		t.Errorf("Expected false, got true")
	}
}

func TestAny(t *testing.T) {
	e := New([]int{1, 2, 3})

	if !e.Any(func(i int) bool { return i == 1 }) {
		t.Errorf("Expected true, got false")
	}
	if e.Any(func(i int) bool { return i == 4 }) {
		t.Errorf("Expected false, got true")
	}
}

func TestFilterThenReduce(t *testing.T) {
	e := New([]int{1, 2, 3})
	result := e.Filter(func(i int) bool { return i > 1 }).Reduce(func(a int, b int) int { return a + b })
	expected := 5

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestFilterThenForEach(t *testing.T) {
	e := New([]int{1, 2, 3})
	result := 0
	e.Filter(func(i int) bool { return i > 1 }).ForEach(func(i int) { result += i })
	expected := 5

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestFilterThenContains(t *testing.T) {
	e := New([]int{1, 2, 3})

	if !e.Filter(func(i int) bool { return i > 1 }).Contains(2) {
		t.Errorf("Expected true, got false")
	}
	if e.Filter(func(i int) bool { return i > 1 }).Contains(1) {
		t.Errorf("Expected false, got true")
	}
}

func TestFilterThenAny(t *testing.T) {
	e := New([]int{1, 2, 3})

	if !e.Filter(func(i int) bool { return i > 1 }).Any(func(i int) bool { return i == 2 }) {
		t.Errorf("Expected true, got false")
	}
	if e.Filter(func(i int) bool { return i > 1 }).Any(func(i int) bool { return i == 1 }) {
		t.Errorf("Expected false, got true")
	}
}

func TestFilterThenTransform(t *testing.T) {
	e := New([]int{1, 2, 3})
	result := Transform(e.Filter(func(i int) bool { return i > 1 }), func(i int) string { return strconv.Itoa(i) })
	expected := New([]string{"2", "3"})

	if len(result.getValues()) != 2 {
		t.Errorf("Expected 2 values, got %d", len(result.getValues()))
	}
	for i, v := range result.getValues() {
		if v != expected.getValues()[i] {
			t.Errorf("Expected %s, got %s", expected.getValues()[i], v)
		}
	}
}

func TestAll(t *testing.T) {
	e := New([]int{1, 2, 3})

	if !e.All(func(i int) bool { return i > 0 }) {
		t.Errorf("Expected true, got false")
	}
	if e.All(func(i int) bool { return i > 2 }) {
		t.Errorf("Expected false, got true")
	}
}

func TestTake(t *testing.T) {
	e := New([]int{1, 2, 3})
	result := e.Take(2).Apply()
	expected := New([]int{1, 2})

	if len(result.getValues()) != 2 {
		t.Errorf("Expected 2 values, got %d", len(result.getValues()))
	}
	for i, v := range result.getValues() {
		if v != expected.getValues()[i] {
			t.Errorf("Expected %d, got %d", expected.getValues()[i], v)
		}
	}
}

func TestTakeOverflow(t *testing.T) {
	e := New([]int{1, 2, 3})
	result := e.Take(5).Apply()
	expected := New([]int{1, 2, 3})

	if len(result.getValues()) != 3 {
		t.Errorf("Expected 3 values, got %d", len(result.getValues()))
	}
	for i, v := range result.getValues() {
		if v != expected.getValues()[i] {
			t.Errorf("Expected %d, got %d", expected.getValues()[i], v)
		}
	}
}

func TestTakeNegative(t *testing.T) {
	e := New([]int{1, 2, 3})
	result := e.Take(-2).Apply()
	expected := New([]int{2, 3})

	if len(result.getValues()) != 2 {
		t.Errorf("Expected 0 values, got %d", len(result.getValues()))
	}
	for i, v := range result.getValues() {
		if v != expected.getValues()[i] {
			t.Errorf("Expected %d, got %d", expected.getValues()[i], v)
		}
	}
}

func TestSkip(t *testing.T) {
	e := New([]int{1, 2, 3})
	result := e.Skip(2).Apply()
	expected := New([]int{3})

	if len(result.getValues()) != 1 {
		t.Errorf("Expected 1 values, got %d", len(result.getValues()))
	}
	for i, v := range result.getValues() {
		if v != expected.getValues()[i] {
			t.Errorf("Expected %d, got %d", expected.getValues()[i], v)
		}
	}
}

func TestSkipOverflow(t *testing.T) {
	e := New([]int{1, 2, 3})
	result := e.Skip(5).Apply()
	expected := New([]int{})

	if len(result.getValues()) != 0 {
		t.Errorf("Expected 0 values, got %d", len(result.getValues()))
	}
	for i, v := range result.getValues() {
		if v != expected.getValues()[i] {
			t.Errorf("Expected %d, got %d", expected.getValues()[i], v)
		}
	}
}

func TestSkipNegative(t *testing.T) {
	e := New([]int{1, 2, 3})
	result := e.Skip(-2).Apply()
	expected := New([]int{1})

	if len(result.getValues()) != 1 {
		t.Errorf("Expected 1 values, got %d", len(result.getValues()))
	}
	for i, v := range result.getValues() {
		if v != expected.getValues()[i] {
			t.Errorf("Expected %d, got %d", expected.getValues()[i], v)
		}
	}
}

func TestTakeWhile(t *testing.T) {
	e := New([]int{1, 2, 3})
	result := e.TakeWhile(func(i int) bool { return i < 3 }).Apply()
	expected := New([]int{1, 2})

	if len(result.getValues()) != 2 {
		t.Errorf("Expected 2 values, got %d", len(result.getValues()))
	}
	for i, v := range result.getValues() {
		if v != expected.getValues()[i] {
			t.Errorf("Expected %d, got %d", expected.getValues()[i], v)
		}
	}
}

func TestTakeWhileFalse(t *testing.T) {
	e := New([]int{1, 2, 3})
	result := e.TakeWhile(func(i int) bool { return i < 0 }).Apply()
	expected := New([]int{})

	if len(result.getValues()) != 0 {
		t.Errorf("Expected 0 values, got %d", len(result.getValues()))
	}
	for i, v := range result.getValues() {
		if v != expected.getValues()[i] {
			t.Errorf("Expected %d, got %d", expected.getValues()[i], v)
		}
	}
}

func TestSkipWhile(t *testing.T) {
	e := New([]int{1, 2, 3})
	result := e.SkipWhile(func(i int) bool { return i < 3 }).Apply()
	expected := New([]int{3})

	if len(result.getValues()) != 1 {
		t.Errorf("Expected 1 values, got %d", len(result.getValues()))
	}
	for i, v := range result.getValues() {
		if v != expected.getValues()[i] {
			t.Errorf("Expected %d, got %d", expected.getValues()[i], v)
		}
	}
}

func TestSkipWhileFalse(t *testing.T) {
	e := New([]int{1, 2, 3})
	result := e.SkipWhile(func(i int) bool { return i < 0 }).Apply()
	expected := New([]int{1, 2, 3})

	if len(result.getValues()) != 3 {
		t.Errorf("Expected 3 values, got %d", len(result.getValues()))
	}
	for i, v := range result.getValues() {
		if v != expected.getValues()[i] {
			t.Errorf("Expected %d, got %d", expected.getValues()[i], v)
		}
	}
}

func TestNestedAppend(t *testing.T) {
	e1 := New([]int{1, 2, 3})
	e2 := New([]int{4, 5, 6})
	e3 := New([]int{7, 8, 9})
	e := New([]IEnumerable[int]{e1, e2})
	result := e.Append(e3).Apply()
	expected := New([]IEnumerable[int]{e1, e2, e3})
	if len(result.getValues()) != 3 {
		t.Errorf("Expected 3 values, got %d", len(result.getValues()))
	}
	for i, v := range result.getValues() {
		for j, w := range v.getValues() {
			if w != expected.getValues()[i].getValues()[j] {
				t.Errorf("Expected %d, got %d", expected.getValues()[i].getValues()[j], w)
			}
		}
	}
}
