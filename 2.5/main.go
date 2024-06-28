package main

import (
	"fmt"
	"unicode"
)

func main() {
	var s string
	_, _ = fmt.Scan(&s)
	checkPassword(s)
}

func checkPassword(s string) {
	stringRunes := []rune(s)
	if len(stringRunes) < 5 {
		fmt.Print("Wrong password")
		return
	} else {
		for i := 0; i < len(stringRunes); i++ {
			if unicode.IsDigit(stringRunes[i]) || unicode.Is(unicode.Latin, stringRunes[i]) {
				continue
			} else {
				fmt.Print("Wrong password")
				return
			}
		}
	}
	fmt.Print("Ok")
}
