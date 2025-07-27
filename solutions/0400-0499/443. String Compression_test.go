package main

import (
	"strconv"
	"testing"
)

/*
Given an array of characters chars, compress it using the following algorithm:

Begin with an empty string s. For each group of consecutive repeating characters in chars:

If the group's length is 1, append the character to s.
Otherwise, append the character followed by the group's length.
The compressed string s should not be returned separately, but instead, be stored in the input character array chars. Note that group lengths that are 10 or longer will be split into multiple characters in chars.

After you are done modifying the input array, return the new length of the array.

You must write an algorithm that uses only constant extra space.


Example 1:
Input: chars = ["a","a","b","b","c","c","c"]
Output: Return 6, and the first 6 characters of the input array should be: ["a","2","b","2","c","3"]
Explanation: The groups are "aa", "bb", and "ccc". This compresses to "a2b2c3".

Example 2:
Input: chars = ["a"]
Output: Return 1, and the first character of the input array should be: ["a"]
Explanation: The only group is "a", which remains uncompressed since it's a single character.

Example 3:
Input: chars = ["a","b","b","b","b","b","b","b","b","b","b","b","b"]
Output: Return 4, and the first 4 characters of the input array should be: ["a","b","1","2"].
Explanation: The groups are "a" and "bbbbbbbbbbbb". This compresses to "ab12".


Constraints:
1 <= chars.length <= 2000
chars[i] is a lowercase English letter, uppercase English letter, digit, or symbol.
*/

// compress performs in-place compression of a character array.
// Consecutive repeating characters are replaced by the character followed by the count
// (if count > 1). The result is stored in the input array and the new length is returned.
//
// Example:
//
//	Input:  ['a','a','b','b','c','c','c']
//	Output: ['a','2','b','2','c','3'] with return value 6
//
// Parameters:
//
//	chars []byte - Input array of characters to compress in-place
//
// Returns:
//
//	int - New length of the compressed array
func compress(chars []byte) int {
	// Handle edge cases: empty or single character arrays
	if len(chars) == 0 || len(chars) == 1 {
		return len(chars)
	}

	// Initialize variables for compression
	var compress []byte       // Temporary slice to build compressed result
	lastCharacter := chars[0] // Track the current character being counted
	counter := 1              // Count of consecutive occurrences

	// Iterate through the array starting from index 1
	for i := 1; i < len(chars); i++ {
		currentCharacter := chars[i]

		// If current character matches the last one, increment counter
		if currentCharacter == lastCharacter {
			counter++
			continue
		}

		// Character changed: add previous character and count (if > 1) to result
		compress = append(compress, lastCharacter)
		if counter > 1 {
			// Convert counter to individual digits and append each digit
			countStr := strconv.Itoa(counter)
			for j := 0; j < len(countStr); j++ {
				compress = append(compress, countStr[j])
			}
		}

		// Update tracking variables for new character
		lastCharacter = currentCharacter
		counter = 1
	}

	// Handle the last group of characters
	compress = append(compress, lastCharacter)
	if counter > 1 {
		countStr := strconv.Itoa(counter)
		for j := 0; j < len(countStr); j++ {
			compress = append(compress, countStr[j])
		}
	}

	// Copy the compressed result back to the original array
	for i := 0; i < len(compress) && i < len(chars); i++ {
		chars[i] = compress[i]
	}

	// Return the length of the compressed result
	return len(compress)
}

func TestCompress(t *testing.T) {
	testCases := []struct {
		in   []byte
		want int
	}{
		{in: []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}, want: 6},
		{in: []byte{'a'}, want: 1},
		{in: []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}, want: 4},
	}
	for _, tc := range testCases {
		t.Run("443. String Compression", func(t *testing.T) {
			out := compress(tc.in)

			if out != tc.want {
				t.Errorf("compress(%v) = %d, want: %d", tc.in, out, tc.want)
			}
		})
	}
}
