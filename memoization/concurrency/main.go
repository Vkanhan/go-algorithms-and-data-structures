package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// Memo struct that holds the cached results and a mutex for synchronization.
type Memo struct {
	f     Func					// The function to be memoized
	cache map[string]result		// Cache to store the results
	mu    sync.Mutex			// Mutex to ensure thread-safe access to the cache

}

// Func is the type of the function to memoize.
type Func func(key string) (any, error)

// result is the result of calling a Func.
type result struct {
	value any
	err   error
}

// New initialize the Memo struct with the function and an empty cache
func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

// Get returns the cached result for the given key.
func (memo *Memo) Get(key string) (any, error) {
	// Lock the mutex to ensure safe access to the cache
	memo.mu.Lock()
	// Check if the result is already cached
	res, found := memo.cache[key]
	if !found {
		// If not found, unlock the mutex before calling the expensive function
		memo.mu.Unlock()
		value, err := memo.f(key)
		// Lock the mutex again before updating the cache
		memo.mu.Lock()
		// Store the result in the cache
		res = result{value, err}
		memo.cache[key] = res
	}
	// Unlock the mutex before returning the result
	memo.mu.Unlock()
	return res.value, res.err
}

// expensiveOperation simulates a function with a delay.
func expensiveOperation(key string) (any, error) {
	// Simulate a time-consuming operation
	time.Sleep(2 * time.Second)
	return fmt.Sprintf("Result for %s", key), nil
}

func main() {
	// Create a memoization of expensiveOperation.
	m := New(expensiveOperation)

	// A list of keys to process.
	keys := []string{"alpha", "beta", "alpha", "gamma", "beta", "gamma", "alpha"}

	// Loop through the keys and process each one
	for _, key := range keys {
		start := time.Now()
		value, err := m.Get(key)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Key: %s, Value: %s, Time: %s\n", key, value, time.Since(start))
	}
}
