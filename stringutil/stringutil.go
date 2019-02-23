// Package stringutil contains utility functions for working with strings.
package stringutil

import (
	"strings"
)

// Reverse returns its argument string reversed rune-wise left to right.
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func IsPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	for i < j {
		if string(s[i]) == " " || string(s[i]) == "," || string(s[i]) == ":" {
			i++
			continue
		}
		if string(s[j]) == " " || string(s[j]) == "," || string(s[j]) == ":" {
			j--
			continue
		}
		if strings.ToLower(string(s[i])) != strings.ToLower(string(s[j])) {
			return false
		}
		i++
		j--
	}
	return true
}
