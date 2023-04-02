package kata

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	tests := []struct {
		param    string
		expected string
	}{
		{"I.was.going.fishing.that.morning.at.ten.o'clock", "c.nhsoI\nltiahi.\noentinw\ncng.nga\nk..mg.s\n\voao.f.\n\v'trtig"},
		{"Process terminated with status 0 (0 minute(s), 6 second(s))", "s t setP\n)se(tder\n)e(0a ro\n\vcs twmc\n\vo)muiie\n\vn,istns\n\vd n has\n\v(6u0 t "},
	}

	for i, test := range tests {
		actual := Code(test.param)
		msg := fmt.Sprintf("test Code %d not passed\n", i)
		assert.Equal(t, test.expected, actual, msg)

		actual = Decode(test.expected)
		msg = fmt.Sprintf("test Decode %d not passed\n", i)
		assert.Equal(t, test.param, actual, msg)
	}
}
