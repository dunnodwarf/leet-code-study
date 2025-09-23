package array

import (
	"fmt"
	"runtime"
	"strconv"
	"testing"
)

// time complexity: O(n)
// space complexity: O(n)
func fizzBuzz(n int) []string {
	var result []string
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			result = append(result, "FizzBuzz")
		} else if i%3 == 0 {
			result = append(result, "Fizz")
		} else if i%5 == 0 {
			result = append(result, "Buzz")
		} else {
			result = append(result, strconv.Itoa(i))
		}
	}

	return result
}

func TestFizzBuzz(t *testing.T) {
	var m1, m2 runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m1)

	have := fizzBuzz(5)

	runtime.GC()
	runtime.ReadMemStats(&m2)

	fmt.Printf("Memory used: %d bytes\n", m2.TotalAlloc-m1.TotalAlloc)
	fmt.Printf("Heap objects: %d\n", m2.HeapObjects-m1.HeapObjects)
	//want := "arroz"
	fmt.Println("Result:", have)
}
