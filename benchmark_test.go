package enumerable

import (
	"strconv"
	"testing"
	"time"
)

func BenchmarkParallelMap(b *testing.B) {
	e := New(make([]int, 1000))
	b.Run("ParallelMap", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			e.MapParallel(func(i int) int {
				time.Sleep(time.Millisecond)
				return i + 2
			}).Apply()
		}
	})

	b.Run("Map", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			e.Map(func(i int) int {
				time.Sleep(time.Millisecond)
				return i + 2
			}).Apply()
		}
	})
}

func BenchmarkParallelForEach(b *testing.B) {
	e := New(make([]int, 1000))
	b.Run("ParallelForEach", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			e.ForEachParallel(func(i int) {
				time.Sleep(time.Millisecond)
			})
		}
	})

	b.Run("ForEach", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			e.ForEach(func(i int) {
				time.Sleep(time.Millisecond)
			})
		}
	})
}

func BenchmarkParallelTransform(b *testing.B) {
	e := New(make([]int, 1000))
	b.Run("ParallelTransform", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			TransformParallel(e, func(i int) string {
				time.Sleep(time.Millisecond)
				return strconv.Itoa(i)
			}).Apply()
		}
	})

	b.Run("Transform", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Transform(e, func(i int) string {
				time.Sleep(time.Millisecond)
				return strconv.Itoa(i)
			}).Apply()
		}
	})
}
