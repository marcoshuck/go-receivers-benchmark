/*
BenchmarkPointers
	BenchmarkPointers-4   	 5207644	       224 ns/op
*/
package go_receivers_benchmark

import (
	"github.com/marcoshuck/go-receivers-benchmark/logger"
	"github.com/marcoshuck/go-receivers-benchmark/pointers"
	"testing"
)

/*
MEMORY PROFILE
---------------------
Showing nodes accounting for 679.19MB, 99.54% of 682.33MB total
Dropped 14 nodes (cum <= 3.41MB)
      flat  flat%   sum%        cum   cum%
  485.19MB 71.11% 71.11%   485.19MB 71.11%  github.com/marcoshuck/go-receivers-benchmark/logger.(*writer).Write
     194MB 28.43% 99.54%   679.19MB 99.54%  github.com/marcoshuck/go-receivers-benchmark/logger.(*logger).Log
         0     0% 99.54%   679.19MB 99.54%  command-line-arguments.BenchmarkPointers
         0     0% 99.54%   485.19MB 71.11%  fmt.Fprintf
         0     0% 99.54%   679.19MB 99.54%  github.com/marcoshuck/go-receivers-benchmark/pointers.(*service).Do
         0     0% 99.54%   679.19MB 99.54%  testing.(*B).launch
         0     0% 99.54%   679.19MB 99.54%  testing.(*B).runN

CPU PROFILE
---------------------
Showing nodes accounting for 910ms, 62.33% of 1460ms total
Showing top 10 nodes out of 65
      flat  flat%   sum%        cum   cum%
     140ms  9.59%  9.59%      140ms  9.59%  runtime.memmove
     130ms  8.90% 18.49%      440ms 30.14%  fmt.(*pp).doPrintf
     110ms  7.53% 26.03%      340ms 23.29%  runtime.mallocgc
     100ms  6.85% 32.88%      100ms  6.85%  runtime.nextFreeFast (inline)
      90ms  6.16% 39.04%      240ms 16.44%  fmt.(*pp).printArg
      80ms  5.48% 44.52%      120ms  8.22%  fmt.(*buffer).writeString (inline)
      80ms  5.48% 50.00%      100ms  6.85%  sync.(*Pool).Get
      70ms  4.79% 54.79%      920ms 63.01%  fmt.Fprintf
      60ms  4.11% 58.90%       90ms  6.16%  sync.(*Pool).Put
      50ms  3.42% 62.33%      140ms  9.59%  fmt.(*pp).free

*/
func BenchmarkPointers(b *testing.B) {
	l := logger.NewLogger("Pointers")
	s := pointers.NewService(l)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Do()
	}
}
