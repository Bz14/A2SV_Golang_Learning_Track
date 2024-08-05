package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CheckPalindrome(st string)bool{
	var newSt [] string
	for _, ch := range st{
		isLetter := (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
		isDigit := (ch >= '0' && ch <= '9')
		if isDigit || isLetter{
			newSt = append(newSt, strings.ToLower(string(ch)))
		}
	}
	l := 0
	r := len(newSt) - 1
	for l <= r{
		if newSt[l] != newSt[r]{
			return false
		}
		l += 1
		r -= 1
	}
	return true
}

func main(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\n\t========== Check Palindrome ==========")
	fmt.Print("\n\tEnter a string: ")
	str, _ := reader.ReadString('\n')
	if CheckPalindrome(strings.TrimSpace(str)){
		fmt.Printf("\t%s is a palindrome.\n",strings.TrimSpace(str))
	}else{
		fmt.Printf("\t%s is not a palindrome.\n",strings.TrimSpace(str))
	}
	fmt.Println("\n\t======================================")
}