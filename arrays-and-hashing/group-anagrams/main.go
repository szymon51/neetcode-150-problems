package main

import "fmt"

// Given an array of strings strs, group the anagrams together. You can return the answer in any order.
// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase,
// typically using all the original letters exactly once.
func groupAnagrams(strs []string) [][]string {
	// initialize a map with list of
	anagramMap := make(map[[26]int][]string)

	// iterate through words from the strs array
	for _, word := range strs {
		// array with 26 elements for representing each letter of the alphabet
		charCount := [26]int{}

		// iterate through characters of every word
		for _, char := range word {
			// add one to the index of the current character in the character array
			charCount[char-'a']++
		}
		// for words that are each others anagrams add them to the corresponding anagram group
		anagramMap[charCount] = append(anagramMap[charCount], word)
	}

	result := make([][]string, len(anagramMap))
	index := 0

	// iterate through values of the anagram hashmap
	for _, val := range anagramMap {
		// add the anagram group to the result array
		result[index] = val
		index++
	}
	return result
}

func main() {
	// 	Example 1:
	// Input: strs = ["eat","tea","tan","ate","nat","bat"]
	// Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))

	// Example 2:
	// Input: strs = [""]
	// Output: [[""]]
	fmt.Println(groupAnagrams([]string{""}))

	// Example 3:
	// Input: strs = ["a"]
	// Output: [["a"]]
	fmt.Println(groupAnagrams([]string{"a"}))

}
