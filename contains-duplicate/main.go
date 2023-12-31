package main

import (
	"fmt"
)

func containsDuplicate(nums []int) bool {
	previousValues := make(map[int]bool)
	for _, num := range nums {
		if previousValues[num] == true {
			return true
		}
		previousValues[num] = true
	}
	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))
	fmt.Println(containsDuplicate([]int{1, 2, 3, 4}))
	fmt.Println(containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
}
