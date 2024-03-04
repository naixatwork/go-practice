package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	f := func(character rune) rune {
		if !unicode.IsLetter(character) && !unicode.IsNumber(character) {
			return -1
		}

		return unicode.ToLower(character)
	}

	filteredString := strings.Map(f, s)
	leftSidePointer := 0
	rightSidePointer := len(filteredString) - 1

	for leftSidePointer <= rightSidePointer {
		if filteredString[leftSidePointer] != filteredString[rightSidePointer] {
			return false
		}
		leftSidePointer = leftSidePointer + 1
		rightSidePointer = rightSidePointer - 1
	}

	return true
}

func main() {
	fmt.Println(isPalindrome("a."))
	//fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	//fmt.Println(isPalindrome("race a car"))
}
