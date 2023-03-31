// https://www.codewars.com/kata/521c2db8ddc89b9b7a0000c1

package kata

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Snail(snaipMap [][]int) []int {
	snail := []int{}

	// empty matrix
	if len(snaipMap) == 1 && len(snaipMap[0]) == 0 {
		return snail
	}

	// 1x1 matrix
	if len(snaipMap) == 1 && len(snaipMap[0]) == 1 {
		return []int{snaipMap[0][0]}
	}

	// n is the limit
	n := len(snaipMap)

	// indexes for row and column
	row, col := 0, 0

	// current direction
	dir := 0

	// loop over all elements
	// n*n is the number of total elements
	for i := 0; i < n*n; i++ {
		snail = append(snail, snaipMap[row][col])

		// set -1 after storing the cell
		// it will be used as limit
		snaipMap[row][col] = -1

		switch dir {
		case 0: // right
			col += 1
			if col == n-1 || snaipMap[row][col+1] == -1 {
				dir = 1
			}
		case 1: // down
			row += 1
			if row == n-1 || snaipMap[row+1][col] == -1 {
				dir = 2
			}
		case 2: // left
			col -= 1
			if col == 0 || snaipMap[row][col-1] == -1 {
				dir = 3
			}
		case 3: // top
			row -= 1
			if row == 0 || snaipMap[row-1][col] == -1 {
				dir = 0
			}
		}
	}

	return snail
}

func TestSnail(t *testing.T) {
	tests := []struct {
		param    [][]int
		expected []int
	}{
		{[][]int{{}}, []int{}},
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
		{[][]int{{1}}, []int{1}},
	}

	for i, test := range tests {
		actual := Snail(test.param)
		assert.Equal(t, test.expected, actual)
		fmt.Printf("test %d not passed\n", i)
	}
}
