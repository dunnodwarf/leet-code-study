package blind75

import (
	"slices"
	"testing"
)

/*
Integer overflow:
When an arithmetic operation produces a result that is larger than this maximum value, it
"overflows". The behavior can vary by language and type, but often it "wraps around" to a
negative number or a very small positive number, leading to incorrect calculations.

While Go's int is usually large enough, adopting the overflow-safe calculation makes my code
more robust if Go's int behavior were to change in some niche context (unlikely, but good
to be aware of)

Many developers learn this pattern in other languages (like Java or C++) where it's absolutely
critical. Using it consistently across languages and projects reduces cognitive load and potential
for error.

The overflow-safe calculation typically has no measurable performance penalty compared to the
simpler one. The compiler will optimize both effectively.

For those aware of the overflow issue, low + (high - low) / 2 explicitly communicates that you're
being careful about potential overflows, which can be a form of self-documenting code.
*/

// time complexity:
func binarySearch(numb int, arr []int) int {
	slices.Sort(arr) //with this line the time complexity is O(n log n); if the list is already sorted this line is not necessary

	low := 0
	high := len(arr) - 1

	for low <= high {

		mid := low + (high-low)/2 //calculate the middle index, using the overflow-safe method
		//mid := high + low /2 -> could cause integer overflow

		if arr[mid] == numb {
			return mid
		} else if arr[mid] < numb {
			low = mid + 1 //target is in the right half
		} else {
			high = mid - 1 //targe is in the left half
		}

	}

	return -1 //target not found in the list
}

func Test(t *testing.T) {

	tests := []struct {
		name     string
		arr      []int
		target   int
		expected int
	}{
		{"Found in middle", []int{1, 2, 3, 4, 5}, 3, 2},
		{"Found at start", []int{1, 2, 3, 4, 5}, 1, 0},
		{"Found at end", []int{1, 2, 3, 4, 5}, 5, 4},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := binarySearch(testCase.target, testCase.arr)
			if result != testCase.expected {
				t.Errorf("binarySearch(%v, %v) = %v; want %v", testCase.target, testCase.arr, result, testCase.expected)
			}
		})
	}

}
