package main

import (
	"fmt"
)


func main() {
	
	chess := make(map[int]uint64, 64)
	counter := uint64(1)
	for i:=1; i<64; i++ {
		chess[i] = counter
		counter *= 2
		fmt.Println(i, counter)
	}
	

}