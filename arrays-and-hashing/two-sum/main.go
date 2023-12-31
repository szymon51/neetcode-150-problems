package main

import "fmt"

// Given an array of integers nums and an integer target,
// return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution,
// and you may not use the same element twice.
// You can return the answer in any order.
func twoSum(nums []int, target int) []int {

	// hashmap to store every number as a key and its index as the value
	hashmap := make(map[int]int)

	// iterate through every element in the nums array
	for i, number := range nums {

		// if the value of the current number minus the target is in the hashmap
		// return the index of the current number and the index of the number in the hashmap
		if val, found := hashmap[target-number]; found {
			return []int{val, i}
		}

		// else put the number and its index in the hashmap
		hashmap[number] = i
	}
	return nil
}

func main() {

	// Example 1:
	// Input: nums = [2,7,11,15], target = 9
	// Output: [0,1]
	// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
	fmt.Println("Example 1 output: ", twoSum([]int{2, 7, 11, 5}, 9))

	// Example 2:
	// Input: nums = [3,2,4], target = 6
	// Output: [1,2]
	fmt.Println("Example 2 output: ", twoSum([]int{3, 2, 4}, 6))

	// Example 3:
	// Input: nums = [3,3], target = 6
	// Output: [0,1]
	fmt.Println("Example 3 output: ", twoSum([]int{3, 3}, 6))

}
