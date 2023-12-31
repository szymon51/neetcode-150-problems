package main

import "fmt"

// Given two strings s and t, return true if t is an anagram of s, and false otherwise.
// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase,
// typically using all the original letters exactly once.
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	lettersInS := make(map[byte]int)
	lettersInT := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		lettersInS[s[i]] = lettersInS[s[i]] + 1
		lettersInT[t[i]] = lettersInT[t[i]] + 1
	}

	for key, value := range lettersInS {
		if lettersInT[key] != value {
			return false
		}
	}
	return true

}

func main() {
	// 	Example 1:
	// Input: s = "anagram", t = "nagaram"
	// Output: true
	fmt.Println(isAnagram("anagram", "nagaram"))

	// Example 2:
	// Input: s = "rat", t = "car"
	// Output: false
	fmt.Println(isAnagram("rat", "car"))

	fmt.Println(isAnagram("hello", "lleho"))
	fmt.Println(isAnagram("helro", "lleho"))
}
