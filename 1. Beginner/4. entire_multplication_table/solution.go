package main

import "fmt"

//Coding challenge #4
//Print the multiplication
//tables with numbers from
//1 to 10

func main() {

	for x := 1; x < 10; x++ {
		for y := 1; y < 10; y++ {
			fmt.Println(x, "X", y, " = ", x*y)
		}
	}
}
