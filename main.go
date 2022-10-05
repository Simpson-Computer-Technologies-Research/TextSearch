package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	var (
		// Full Text Search Using A Byte Array
		FullTextSearchBytesResult, _ = FullTextSearchBytes([]byte("simpson"))
		// Full Text Search Using An Array
		FullTextSearchArrayResult, _ = FullTextSearchArray("simpson")
		// Full Text Search Using A Map
		FullTextSearchMapResult, _ = FullTextSearchMap("simpson")
	)
	// Therefore after testing the above functions,
	// it is safe to conclude that using an array
	// for storing data that you are going to
	// be iterating over is the fastest method.

	// This was just a quick example showing the
	// difference between searching through
	// bytes, arrays and maps
	fmt.Printf("\n>> FullTextSearchBytes(): %dns\n", FullTextSearchBytesResult)
	fmt.Printf("\n>> FullTextSearchArray(): %dns\n", FullTextSearchArrayResult)
	fmt.Printf("\n>> FullTextSearchMap(): %dns\n\n", FullTextSearchMapResult)
}

// The FullTextSearchBytes() method stores data
// in a byte array.
//
// To find the provided 'look_for: string', a for loop is
// utilized, iterating over the characters in the data string.
//
// If the current index is equal to the provided 'look_for: string's
// first letter, then append the current index to the result array
func FullTextSearchBytes(look_for []byte) (int64, []int) {
	// Declare Variables used for testing
	var result []int
	var data []byte
	for i := 0; i < 100000; i++ {
		data = append(data, "heymynameistristansimpson"...)
	}
	// Establish a start time for tracking speed
	var startTime int64 = time.Now().UnixNano()

	// Iterate over the data
	for i := 0; i < len(data); i++ {
		if data[i] == look_for[0] && data[i+len(look_for)-1] == look_for[len(look_for)-1] {
			result = append(result, i)
			i = i + len(look_for)
		}
	}
	// Track the finish time of the program
	return time.Now().UnixNano() - startTime, result
}

// The FullTextSearchArray() method stores the
// data inside an array then iterates over said
// array and checks whether it contains the provided
// look_for: string.
func FullTextSearchArray(look_for string) (int64, []int) {
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

// The FullTextSearchMap() used a map to store the
// data. This method is much slower than the array method.
func FullTextSearchMap(look_for string) (int64, []int) {
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
