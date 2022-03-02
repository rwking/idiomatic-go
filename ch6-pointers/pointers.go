package main

import (
	"fmt"
)

func main() {
	a := 5
	fmt.Println("value of a: ", a)
	aPointer := &a
	fmt.Println("stored at: ", aPointer)
	fmt.Println("stored at: ", *aPointer)
	*aPointer = 10
	fmt.Println("stored at: ", aPointer)
	fmt.Println("stored at: ", *aPointer)
	aPointerPointer := &aPointer
	fmt.Println("aPointer stored at: ", aPointerPointer)
	fmt.Println("aPointer stored at: ", *aPointerPointer)
}