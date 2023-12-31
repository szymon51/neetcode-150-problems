package main

import "fmt"

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
	fmt.Println(isAnagram("hello", "lleho"))
	fmt.Println(isAnagram("helro", "lleho"))
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("rat", "car"))
}
