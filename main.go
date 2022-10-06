package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	var (
		// Text Search Using An Array
		TextSearchArrayResult, _ = TextSearchArray("simpson")
		// Text Search Using A Map
		TextSearchMapResult, _ = TextSearchMap("simpson")
	)
	// Therefore after testing the above functions,
	// it is safe to conclude that using an array
	// for storing data that you are going to
	// be iterating over is the fastest method.

	// This was just a quick example showing the
	// difference between searching through arrays and maps
	fmt.Printf("\n>> TextSearchArray(): %dns\n", TextSearchArrayResult)
	fmt.Printf("\n>> TextSearchMap(): %dns\n\n", TextSearchMapResult)
}

// The TextSearchArray() method stores the
// data inside an array then iterates over said
// array and checks whether it contains the provided
// look_for: string.
func TextSearchArray(look_for string) (int64, []int) {
	// Establish an array filled with strings.
	// This array is what we will be iterating through
	var data []string
	for i := 0; i < 100000; i++ {
		data = append(data, "heymynameistristansimpson")
	}
	var result []int

	// Establish a start time for tracking speed
	var startTime int64 = time.Now().UnixNano()

	// Iterate over the data
	for i := 0; i < len(data); i++ {
		if strings.Contains(data[i], look_for) {
			result = append(result, i)
		}
	}
	// Track the finish time of the program
	return time.Now().UnixNano() - startTime, result
}

// The TextSearchMap() used a map to store the
// data. This method is much slower than the array method.
func TextSearchMap(look_for string) (int64, []int) {
	// Establish an array filled with strings.
	// This array is what we will be iterating through
	var data map[int]string = map[int]string{}
	for i := 0; i < 100000; i++ {
		data[i] = "heymynameistristansimpson"
	}
	var result []int

	// Establish a start time for tracking speed
	var startTime int64 = time.Now().UnixNano()

	// Iterate over the data
	for k, v := range data {
		if strings.Contains(v, look_for) {
			result = append(result, k)
		}
	}
	// Track the finish time of the program
	return time.Now().UnixNano() - startTime, result
}
