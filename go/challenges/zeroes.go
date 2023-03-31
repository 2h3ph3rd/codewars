package kata

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
https://www.codewars.com/kata/52f787eb172a8b4ae1000a34

To calculate the number of final zeros of a factorial, it is sufficient to calculate the number of 5s.

In fact, the only way to get a final zero is to multiply by 10 = 5 * 2.

It is not enough to divide the number by 5 because 25 (= 5 * 5) and its multiples has one additional five.

Zeroes(13) == 2 -> 13! = ... * 5 * 2 * 10 * ...
Zeroes(17) == 3 -> 17! = ... * 5 * 2 * 10 * 15 (5*3) * ...
Zeroes(25) == 6 -> 17! = ... * 5 * 2 * 10 * 15 (5*3) * 20 * 25 (5*5) * ...
*/

func Zeros(base, number int) int {
	num, err := strconv.ParseInt(strconv.FormatInt(int64(number), 10), base, 64)
	if err != nil {
		return 0
	}

	zeros := 0
	for num != 0 {
		num /= 5
		zeros += int(num)
	}

	return zeros
}

func ZerosRec(n int) int {
	if n == 0 {
		return 0
	}
	return n/5 + ZerosRec(n/5)
}

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
		actual := Zeros(test.param1, test.param2)
		// actual := ZerosRec(test.param2)
		msg := fmt.Sprintf("test %d not passed\n", i)
		assert.Equal(t, test.expected, actual, msg)
	}
}
