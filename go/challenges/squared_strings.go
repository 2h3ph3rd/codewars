package kata

import (
	"fmt"
	"math"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/56fcc393c5957c666900024d

func Code(s string) string {
	// len(s) = n*n
	n := int(math.Ceil(math.Sqrt(float64(len(s)))))

	// add padding
	for i := len(s); i < n*n; i++ {
		s += "\v"
	}

	words := []string{}

	// intialize vector
	for i := 0; i < n; i++ {
		words = append(words, "")
	}

	// create blocks
	for i := 0; i < len(s); i++ {
		words[i%n] += string(s[i])
	}

	// update each block in reverse
	for i := range words {
		words[i] = Reverse(words[i])
	}

	// join blocks together
	return strings.Join(words, "\n")
}

func Decode(s string) string {
	// len(s) = n*n
	n := int(math.Sqrt(float64(len(s))))

	//  blocks are divided with newlines
	words := strings.Split(s, "\n")

	decoded := ""
	round := 1
	for i := 0; i < n*n; i++ {
		// increase round after each complete iteration
		if i != 0 && i%n == 0 {
			round += 1
		}

		// recreate original string
		decoded += string(words[i%n][n-round])
	}

	// remove padding
	decoded = strings.TrimRight(decoded, "\v")
	return decoded
}

func TestCode(t *testing.T) {
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
