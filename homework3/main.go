package main

import (
	"fmt"
	"homework3/word_count"
)

func main() {
	s := []string{
		"apple", "banana", "apple" ,"durain", "apple", "banana", "banana",
	}

	wordMap := wordcount.WordCount(s)
	fmt.Printf("\"apple\": %d,\n\"banana\": %d,\n\"durain\": %d,", wordMap["apple"], wordMap["banana"], wordMap["durain"])
}

