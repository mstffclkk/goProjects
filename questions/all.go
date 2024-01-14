package questions

import (
	"fmt"
	"sort"
	"strings"
)

// question 1
func countA(word string) int {
	return strings.Count(word, "a")
}

// SortWords sorts a slice of strings in descending order based on the number of occurrences of the letter 'A' in each word.
// If two words have the same number of 'A' occurrences, the longer word comes first.
// The input slice is modified in-place and also returned.
func SortWords(words []string) []string {
	sort.Slice(words, func(i, j int) bool {
		countA1 := countA(words[i])
		countA2 := countA(words[j])
		if countA1 == countA2 {
			return len(words[j]) < len(words[i])
		}
		return countA1 > countA2
	})

	return words
}

//question 2

// RecursiveFunction is a recursive function that takes an integer n as input and prints the value of n.
// It recursively calls itself with n divided by 2 until n is greater than 1.
func RecursiveFunction(n int) {
	if n > 1 {
		RecursiveFunction(n / 2)
		fmt.Println(n)
	}
}

// question 3

// MostRepeated returns the most repeated string in the given slice of strings.
// It counts the occurrences of each string and returns the one with the highest count.

func MostRepeated(data []string) string {
	repeated := make(map[string]int)
	for _, v := range data {
		repeated[v]++
	}
	var max int
	var result string
	for k, v := range repeated {
		if v > max {
			max = v
			result = k
		}
	}
	return result
}
