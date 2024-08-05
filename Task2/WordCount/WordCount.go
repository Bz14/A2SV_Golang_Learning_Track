package main

import (
	"bufio"
	"fmt"
	"os"
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
	if ch != ""{
		wordCount[ch] += 1
	}
	return wordCount
}
func main(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\n\t========== Check Palindrome ==========")
	fmt.Print("\n\tEnter a string: ")
	str, _ := reader.ReadString('\n')
	fmt.Println("\t======================================")
	word := CountFrequency(strings.TrimSpace(str))
	fmt.Println("\t\tWord\t\tCount")
	fmt.Println("\t======================================")
	for k, v := range word{
		fmt.Printf("\t\t%s\t\t%d\n", k, v)
		fmt.Println("\t======================================")
	}
}