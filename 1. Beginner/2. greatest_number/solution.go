package main

import "fmt"

func main() {
	findgreatest(2, 4, 5, 6, 7)
}

func findgreatest(x ...int) {
	var temp = 0
	for _, s := range x {
		if s > temp {
			temp = s
		} else {
			temp = temp
		}
	}
	fmt.Println("greates number is", temp)
}
