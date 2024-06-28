package main

import (
	"fmt"
)

func main() {
	var a, b int
	_, err := fmt.Scan(&a, &b)
	if err != nil || b == 0 {
		fmt.Print("ошибка")
		return
	}
	fmt.Print(divide(a, b))
}
