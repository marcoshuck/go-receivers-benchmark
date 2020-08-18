/* RESULTS:
BenchmarkValues
	BenchmarkValues-4     	 4223358	       246 ns/op
*/
package go_receivers_benchmark

import (
	"github.com/marcoshuck/go-receivers-benchmark/logger"
	"github.com/marcoshuck/go-receivers-benchmark/values"
	"testing"
)

/*
MEMORY PROFILE
---------------------
Showing nodes accounting for 641.77MB, 99.69% of 643.79MB total
Dropped 12 nodes (cum <= 3.22MB)
      flat  flat%   sum%        cum   cum%
  469.77MB 72.97% 72.97%   469.77MB 72.97%  github.com/marcoshuck/go-receivers-benchmark/logger.(*writer).Write
     172MB 26.72% 99.69%   641.77MB 99.69%  github.com/marcoshuck/go-receivers-benchmark/logger.(*logger).Log
         0     0% 99.69%   641.77MB 99.69%  command-line-arguments.BenchmarkValues
         0     0% 99.69%   469.77MB 72.97%  fmt.Fprintf
         0     0% 99.69%   641.77MB 99.69%  github.com/marcoshuck/go-receivers-benchmark/values.service.Do
         0     0% 99.69%   641.77MB 99.69%  testing.(*B).launch
         0     0% 99.69%   641.77MB 99.69%  testing.(*B).runN

CPU PROFILE
---------------------
Showing nodes accounting for 910ms, 65.94% of 1380ms total
Showing top 10 nodes out of 85
      flat  flat%   sum%        cum   cum%
     140ms 10.14% 10.14%      430ms 31.16%  runtime.mallocgc
     130ms  9.42% 19.57%      130ms  9.42%  runtime.memmove
     120ms  8.70% 28.26%      370ms 26.81%  fmt.(*pp).doPrintf
     100ms  7.25% 35.51%      120ms  8.70%  fmt.(*buffer).writeString (inline)
     100ms  7.25% 42.75%      100ms  7.25%  runtime.memclrNoHeapPointers
      90ms  6.52% 49.28%       90ms  6.52%  runtime.nextFreeFast (inline)
      70ms  5.07% 54.35%      500ms 36.23%  runtime.convTstring
      60ms  4.35% 58.70%       60ms  4.35%  runtime.procyield
      50ms  3.62% 62.32%      100ms  7.25%  runtime.heapBitsSetType
      50ms  3.62% 65.94%       60ms  4.35%  sync.(*Pool).Get
*/
func BenchmarkValues(b *testing.B) {
	l := logger.NewLogger("Values")
	s := values.NewService(l)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Do()
	}
}
