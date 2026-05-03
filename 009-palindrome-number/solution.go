package main

import "strconv"

func isPalindrome(x int) bool {
	charArr := []rune(strconv.Itoa(x))

	for i, value := range charArr{
		if value != charArr[len(charArr)-(i+1)] {
			return false
		}
	}
	return true
}