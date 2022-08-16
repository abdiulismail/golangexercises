package main

import "fmt"

func main() {
	f := tempcoverter(0)
	fmt.Println(f)
}

func tempcoverter(c int) int {
	f := (c * 9 / 5) + 32
	return f
}
