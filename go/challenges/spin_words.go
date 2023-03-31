// https://www.codewars.com/kata/5264d2b162488dc400000001

package kata

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func reverse(s string) string {
	new := []byte{}
	for i := len(s); i > 0; i-- {
		new = append(new, s[i-1])
	}

	return string(new)
}

func SpinWords(str string) string {
	arr := strings.Split(str, " ")

	for i, s := range arr {
		if len(s) >= 5 {
			s = reverse(s)
		}

		arr[i] = s
	}

	return strings.Join(arr, " ")
}

func TestSpinWords(t *testing.T) {
	tests := []struct {
		param    string
		expected string
	}{
		{"Hey fellow warriors", "Hey wollef sroirraw"},
		{"This is a test", "This is a test"},
		{"This is another test", "This is rehtona test"},
	}

	for _, test := range tests {
		actual := SpinWords(test.param)
		assert.Equal(t, test.expected, actual)
	}
}
