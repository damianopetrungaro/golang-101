// I took it online. Do not hate me.
package algo

import (
	"regexp"
	"testing"
	"time"
)

func BenchmarkTimeRegex(b *testing.B) {
	t := "1999-06-05"
	for i := 0; i < b.N; i++ {
		regexp.MatchString("^[0-9]{4}-[0-9]{2}-[0-9]{2}$", t)
	}
}

func BenchmarkTimeParse(b *testing.B) {
	t := "1999-06-05"
	format := "2006-01-02"
	for i := 0; i < b.N; i++ {
		time.Parse(format, t)
	}
}

/**
benchmark                   iter       time/iter   bytes alloc         allocs
---------                   ----       ---------   -----------         ------
BenchmarkTimeRegex-12     300000   4747.00 ns/op     5410 B/op   69 allocs/op
BenchmarkTimeParse-12   20000000     93.20 ns/op        0 B/op    0 allocs/op

ok      githun.com/damianopetrungaro/nice-testing/9_benchmark   3.438s
*/
