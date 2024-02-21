package main

import "fmt"

/*
write a function that accepts a pointer to a 2d slice of strings.

this func need to modify the slice in place so each nested slice contains the longest and shorted word in the same order as the original slice.

no return just modify the slice in place


*/

func main() {
	s := [][]string{
		{"hello", "world", "go", "gopher"},
		{"sun", "moon", "star"},
		{"a", "ab", "abc", "abcd"},
		{"consistent", "in", "size"},
	}
	Solution(&s)
	fmt.Println(s)
}

func Solution(wordSlices *[][]string) {
	for wordIdx, words := range *wordSlices {
		newWords := keepLongestAndShortestWords(words)
		(*wordSlices)[wordIdx] = newWords
	}
}

func keepLongestAndShortestWords(words []string) []string {
	longest := words[0]
	shortest := words[0]
	longestIdx := 0
	shortestIdx := 0

	for idx, word := range words {
		if len(word) > len(longest) {
			longest = word
			longestIdx = idx
		}
		if len(word) < len(shortest) {
			shortest = word
			shortestIdx = idx
		}
	}

	if longest == shortest {
		return []string{longest}
	} else if longestIdx > shortestIdx {
		return []string{shortest, longest}
	} else {
		return []string{longest, shortest}
	}
}
