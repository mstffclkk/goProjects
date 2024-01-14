package main

import (
	"fmt"
	"goProject/questions"
)

func main() {
	// question 1
	words := []string{"aaaasd", "a", "aab", "aaabcd", "ef", "cssssssd", "fdz", "kf", "zc", "lklklklklklklklkl", "l"}
	fmt.Println(questions.SortWords(words))

	// question 2
	questions.RecursiveFunction(9)

	// question 3
	data := []string{"apple", "pie", "apple", "red", "red", "red"}
	fmt.Println(questions.MostRepeated(data))

}
