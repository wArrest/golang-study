package iteration

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a",6)
	want := "aaaaaa"

	assert.Equal(t, want, repeated)
}

//基准测试
//基准测试运行时，代码会运行 b.N 次，并测量需要多长时间。
//代码运行的次数不会对你产生影响，测试框架会选择一个它所认为的最佳值，以便让你获得更合理的结果。
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a",6)
	}
}
