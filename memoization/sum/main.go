package main

import (
	"fmt"
	"strconv"
)

//Simple sum function
func sum(a, b int) int {
	return a + b
}

//Higher order function which takes a func and outputs a memoized version of it
func memo(fn func(a, b int) int) func(a, b int) int {
	// Cache to store the results
	cache := make(map[string]int)

	return func(a, b int) int {
		// Create a unique key for the cache map by converting int to string
		key := strconv.Itoa(a) + " " + strconv.Itoa(b)

		//check if the result is in cache
		if result, found := cache[key]; found {
			return result
		}

		//Compute the result and store in the cache
		fmt.Println("calculating..")

		// Call the original function and store the result 
		result := fn(a, b)

		//Store the result in cache
		cache[key] = result

		return result
	}
}

func main() {
	// Create a memoized version of the sum function
	memoCache := memo(sum)

	fmt.Println(memoCache(4, 5))
}

