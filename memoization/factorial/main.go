package main 

import(
	"fmt"
)

// FactorialMemo calculates the factorial of a number using memoization.
func FactorialMemo(n int, cache map[int]int) int {
	//Check if the result is already in cache
	if result, found := cache[n]; found {
		return result
	}

	//base case
	if n == 0 {
		return 1
	}

	// Recursive computation with memoization
	result := n * FactorialMemo(n - 1, cache)

	// Store the computed result in the cache
	cache[n] = result

	return result
	
}

func main() {

	cache := make(map[int]int)

	// Test the memoized factorial function
	for i := 0; i <= 10; i++ {
		fmt.Printf("Factorial memo(%d): %d\n", i, FactorialMemo(i, cache))
	}
}