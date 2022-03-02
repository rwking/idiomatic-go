package main

import (
	"fmt"
)

func printNumAndBit(num int) {
	fmt.Println(num)
	fmt.Printf("%08b", num) // 111101
	fmt.Println("\n")
}

func main() {
	num := 2
	printNumAndBit(num)
	num <<= 1
	printNumAndBit(num)
	num <<= 1
	printNumAndBit(num)
	num <<= 1
	printNumAndBit(num)
	num <<= 1
	printNumAndBit(num)

	
}
