package main

import "fmt"

func main() {
	f := [9]int{2, 3, -1, 5, 7, 9, 10, 15, 95}
	sum := 0

	for _, x := range f {
		//fmt.Println(f)
		sum = sum + x
	}
	fmt.Println(sum)
}
