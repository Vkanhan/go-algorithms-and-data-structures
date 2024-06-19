package main

import (
	"fmt"
)

// Simple Memoization in fibonacci
func fib(n int) int {
	cache := make(map[int]int)
	if n <= 1 {
		return n
	}
	// Check if result is in the cache
	if result, found := cache[n]; found {
		return result
	}
	//recursivly compute the fibonacci
	result := fib(n-1) + fib(n-2)
	cache[n] = result
	return result
}

// Memoized struct to hold the function and its cache
type Memoized struct {
	f     func(int) int
	cache map[int]int
}

// memoize function to create a Memoized instance for the given function
func memoize(f func(int) int) *Memoized {
	return &Memoized{f: f, cache: make(map[int]int)}
}

// call method to get the result with memoization
func (m *Memoized) call(x int) int {
	// Check if the result is already in the cache
	if result, found := m.cache[x]; found {
		return result
	}

	// Compute the result using the original function
	result := m.f(x)
	// Store the computed result in the cache
	m.cache[x] = result

	return result
}

// Fibonacci function
func fibonacci(n int) int {
	// Base case: return n for 0 and 1
	if n <= 1 {
		return n
	}

	// Recursive computation of Fibonacci
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {

	//Fibonacci memoization
	fmt.Println(fib(6))

	// Create a memoized version of the Fibonacci function
	memoizedFib := memoize(fibonacci)

	// Test the memoized Fibonacci function for the first 10 numbers
	for i := 0; i <= 10; i++ {
		fmt.Printf("Fibonacci(%d): %d\n", i, memoizedFib.call(i))
	}
}
