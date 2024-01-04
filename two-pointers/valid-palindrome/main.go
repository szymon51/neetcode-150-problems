package main

import "fmt"

// converts a letter to lowercase if it was uppercase and returns it oterwise
func toLowerChar(letter byte) byte {
	if letter >= 'A' && letter <= 'Z' {
		letter += 32
	}
	return letter
}

func isNotLetter(letter byte) bool {
	return letter >= 'A' && letter <= 'Z' || letter >= 'a' && letter <= 'z'
}

func isPalindrome(s string) bool {
	ptrL := 0
	ptrR := len(s) - 1

	for i := 0; i < len(s)-1; i++ {
		for ptrL > ptrR {
			if isNotLetter(s[ptrL]) {
				ptrL++
			}
			if isNotLetter(s[ptrR]) {
				ptrR--
			}
		}
		if toLowerChar(s[ptrL]) != toLowerChar(s[ptrR]) {
			return false
		}
		ptrL++
		ptrR--
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

	// Example 3:
	// Input: s = " "
	// Output: true
	// Explanation: s is an empty string "" after removing non-alphanumeric characters.
	// Since an empty string reads the same forward and backward, it is a palindrome.

	fmt.Println(isPalindrome("H allah"))
	fmt.Println(isPalindrome("heh"))
}
