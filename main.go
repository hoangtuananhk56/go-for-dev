package main

import "strings"

func isCircularSentence(sentence string) bool {
	words := strings.Split(sentence, " ")
	if len(words) == 0 || sentence[0] != sentence[len(sentence)-1] {
		return false
	}
	for i := 0; i < len(words)-1; i++ {
		if words[i][len(words[i])-1] != words[i+1][0] {
			return false
		}
	}
	return true
}

func main() {
	sentence := "Leetcode eisc cool"
	result := isCircularSentence(sentence)
	if result {
		println("true")
	} else {
		println("false")
	}
}
