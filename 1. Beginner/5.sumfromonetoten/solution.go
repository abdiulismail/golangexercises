package main

import "fmt"

//Coding challenge #5
//Calculate the sum of
//numbers from 1 to 10

func main() {
	sum := 0
	for x := 1; x < 1000; x++ {
		sum = sum + x
	}
	fmt.Println(sum)
}
