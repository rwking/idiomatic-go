package main

import (
	"fmt"
	"errors"
)


func divideInts(num int, denom int) (result int, remainder int, err error) {
	if denom==0 {
		return 0, 0, errors.New("Cannot divide by zero!")
	}
	result, remainder = num/denom, num%denom
	return result, remainder, err
}


func main() {
	result, remainder, err := divideInts(5,2)
	fmt.Println(result, remainder, err)
}
