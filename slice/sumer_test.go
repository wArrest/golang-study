package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		sum := Sum(numbers)
		want := 15

		assert.Equal(t, want, sum)
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		sum := Sum(numbers)
		want := 6
		assert.Equal(t, want, sum)
	})

}

func TestSumAll(t *testing.T) {
	t.Run("not empty slice", func(t *testing.T) {
		sum := SumAll([]int{1, 2, 3}, []int{4, 5})
		want := []int{6, 9}
		assert.Equal(t, want, sum)
	})

	t.Run("empty slice", func(t *testing.T) {
		sum := SumAll([]int{})
		want := []int{}
		assert.Equal(t, want, sum)
	})
}
