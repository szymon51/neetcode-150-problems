package main

import (
	"fmt"
)

// Given an integer array nums, return true if any value appears at least twice in the array,
// and return false if every element is distinct.
func containsDuplicate(nums []int) bool {
	// store the previous values in a hashmap
	previousValues := make(map[int]bool)

	// interate through the array nums
	for _, num := range nums {
		// if the current number is in the previous values the array contains a duplicate
		if previousValues[num] == true {
			return true
		}
		// update the previous values with the current number
		previousValues[num] = true
	}
	// if none of the values were repeating it means that there is no duplicate
	return false
}

func main() {
	// Example 1:
	// Input: nums = [1,2,3,1]
	// Output: true
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))

	// Example 2:
	// Input: nums = [1,2,3,4]
	// Output: false
	fmt.Println(containsDuplicate([]int{1, 2, 3, 4}))

	// Example 3:
	// Input: nums = [1,1,1,3,3,4,3,2,4,2]
	// Output: true
	fmt.Println(containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
}
