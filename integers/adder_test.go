package integers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Adder(2, 2)
	want := 4
	assert.Equal(t, want, sum)
}
