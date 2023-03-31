package kata

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	tests := []struct {
		param1   int
		param2   int
		expected int
	}{
		{10, 10, 2},
		{10, 16, 3},
		{10, 25, 6},
		{10, 100, 24},
		{10, 1000, 249},
	}

	for i, test := range tests {
		actual := Zeroes(test.param1, test.param2)
		// actual := ZerosRec(test.param2)
		msg := fmt.Sprintf("test %d not passed\n", i)
		assert.Equal(t, test.expected, actual, msg)
	}
}
