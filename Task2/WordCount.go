package main

import (
	"strings"
)

func CountFrequency(st string)map[string]int{
	wordCount := map[string]int{}
	lowerCaseString := strings.ToLower(st)
	ch := ""
	for _, char := range lowerCaseString{
		if (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9'){
			ch += string(char)
		}else if char == ' ' && len(ch) > 0{
			wordCount[ch] += 1
			ch = ""
		}
	}
	wordCount[ch] += 1
	return wordCount
}