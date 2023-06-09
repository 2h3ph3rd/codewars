package kata

import (
	"math"
	"strings"
)

// https://www.codewars.com/kata/56fcc393c5957c666900024d

// reverse return the given string in reverse
// i.e. reverse("Hello!") -> "!olleH"
func reverse(s string) string {
	new := []byte{}
	for i := len(s); i > 0; i-- {
		new = append(new, s[i-1])
	}

	return string(new)
}

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
		words[i] = reverse(words[i])
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
