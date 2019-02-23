package stringutil

import (
	"fmt"
	"testing"
)

func Test_print_string(t *testing.T) {
	fmt.Printf("Hello Go!")
}
func Test_reverse_string(t *testing.T) {
	fmt.Printf(Reverse("!oG ,olleH"))
}

func Test_ispalindrome_string(t *testing.T) {
	b := IsPalindrome("race a car")
	if b == true {
		t.Fatal("race a car : is not a palindrome")
	}

	b = IsPalindrome("A man, a plan, a canal: Panama")
	if b == false {
		t.Fatal("A man, a plan, a canal: Panama : is a palindrome")
	}
}
