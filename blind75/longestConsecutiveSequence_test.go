package blind75

import (
	"fmt"
	"testing"
)

func longestConsecutive(nums []int) int {

	//use map to store nums as key for faster access - O(n) time complexity
	numMap := make(map[int]bool)
	for _, num := range nums {
		numMap[num] = true
	}

	//initialize a variable for the longest value
	longest := 0

	for num := range numMap {

		if !numMap[num] {
			currentStreak := num + 1

			for numMap[currentStreak] == true {
				currentStreak += 1
			}

			longest = max(longest, currentStreak)
		}

	}

	return longest
}

func TestSolution(t *testing.T) {

	numbers := []int{100, 4, 200, 1, 3, 2}

	fmt.Println(longestConsecutive(numbers))
}
