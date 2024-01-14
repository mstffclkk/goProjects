// question 1
package questions

import (
	"fmt"
	"sort"
	"strings"
)

func countA(word string) int {
	return strings.Count(word, "a")
}

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

func RecursiveFunction(n int) {
	if n > 1 {
		RecursiveFunction(n / 2)
		fmt.Println(n)
	}
}

// question 3

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
