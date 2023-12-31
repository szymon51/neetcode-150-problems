package main

import "fmt"

// Given an integer array nums and an integer k, return the k most frequent elements.
// You may return the answer in any order.
func topKFrequent(nums []int, k int) []int {
	numsMap := make(map[int]int)

	for _, num := range nums {
		numsMap[num]++
	}

	// array with the frequency of appereance of each number
	frequency := make([][]int, len(nums)+1)
	for key, val := range numsMap {
		frequency[val] = append(frequency[val], key)
	}

	result := make([]int, 0, k)
	for i := len(frequency) - 1; len(result) != k; i-- {
		for _, val := range frequency[i] {
			if len(result) != k {
				result = append(result, val)
			}
		}
	}

	return result

}

func main() {
	// 	Example 1:
	// Input: nums = [1,1,1,2,2,3], k = 2
	// Output: [1,2]
	fmt.Printf("%v\n", topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))

	// Example 2:
	// Input: nums = [1], k = 1
	// Output: [1]
	fmt.Printf("%v\n", topKFrequent([]int{1}, 1))

}
