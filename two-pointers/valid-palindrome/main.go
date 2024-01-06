package main

import (
	"fmt"
	"unicode"
)

func isLetterOrDigit(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}

// A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters,
// it reads the same forward and backward. Alphanumeric characters include letters and numbers.
// Given a string s, return true if it is a palindrome, or false otherwise.
func isPalindrome(s string) bool {
	leftPtr := 0
	rightPtr := len(s) - 1
	arr := []rune(s)

	for leftPtr < rightPtr {
		left := unicode.ToLower(arr[leftPtr])
		right := unicode.ToLower(arr[rightPtr])
		if !isLetterOrDigit(left) {
			leftPtr++
			continue
		}
		if !isLetterOrDigit(right) {
			rightPtr--
			continue
		}
		if left != right {
			return false
		}
		leftPtr++
		rightPtr--
	}
	return true
}

func main() {
	// Example 1:
	// Input: s = "A man, a plan, a canal: Panama"
	// Output: true
	// Explanation: "amanaplanacanalpanama" is a palindrome.
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))

	// Example 2:
	// Input: s = "race a car"
	// Output: false
	// Explanation: "raceacar" is not a palindrome.
	fmt.Println(isPalindrome("raceacar"))

	// Example 3:
	// Input: s = " "
	// Output: true
	// Explanation: s is an empty string "" after removing non-alphanumeric characters.
	// Since an empty string reads the same forward and backward, it is a palindrome.
	fmt.Println(isPalindrome(" "))

}
