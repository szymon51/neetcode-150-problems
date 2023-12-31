package main

import "fmt"

// Given an integer array nums, return an array answer such that answer[i] is equal to the product
// of all the elements of nums except nums[i].
// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
// You must write an algorithm that runs in O(n) time and without using the division operation.
func productExceptSelf(nums []int) []int {
	prefix := 1
	postfix := 1
	output := make([]int, len(nums))

	for i, num := range nums {
		// store the prefix sums in the output array
		output[i] = prefix
		prefix *= num
	}

	for i := len(nums) - 1; i >= 0; i-- {
		// store the postfix sums multiplied by prefix values in the output array
		output[i] = output[i] * postfix
		postfix = postfix * nums[i]
	}

	return output
}

func main() {
	// 	Example 1:
	// Input: nums = [1,2,3,4]
	// Output: [24,12,8,6]
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))

	// Example 2:
	// Input: nums = [-1,1,0,-3,3]
	// // Output: [0,0,9,0,0]
	fmt.Println(productExceptSelf([]int{-1, 1, 0, -3, 3}))

}
