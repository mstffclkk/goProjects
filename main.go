package main

import (
	"fmt"
	"goProject/questions"
)

func main() {
	// question 1
	fmt.Println("Question 1")
	words := []string{"aaaasd", "a", "aab", "aaabcd", "ef", "cssssssd", "fdz", "kf", "zc", "lklklklklklklklkl", "l"}
	fmt.Println(questions.SortWords(words))
	fmt.Println()

	// question 2
	fmt.Println("Question 2")
	questions.RecursiveFunction(9)
	fmt.Println()

	// question 3
	fmt.Println("Question 3")
	data := []string{"apple", "pie", "apple", "red", "red", "red"}
	fmt.Println(questions.MostRepeated(data))
	fmt.Println()
}
