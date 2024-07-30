package main

import (
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