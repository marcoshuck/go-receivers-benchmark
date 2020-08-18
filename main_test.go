package go_receivers_benchmark

import (
	"github.com/marcoshuck/go-receivers-benchmark/logger"
	"github.com/marcoshuck/go-receivers-benchmark/pointers"
	"github.com/marcoshuck/go-receivers-benchmark/values"
	"testing"
)

func BenchmarkPointers(b *testing.B) {
	l := logger.NewLogger("Pointers")
	s := pointers.NewService(l)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Do()
	}
}

func BenchmarkValues(b *testing.B) {
	l := logger.NewLogger("Pointers")
	s := values.NewService(l)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Do()
	}
}
