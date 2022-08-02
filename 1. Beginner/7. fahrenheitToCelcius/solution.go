package main

import "fmt"

func main() {
	f := tempcoverter(32)
	fmt.Println(f)
}

func tempcoverter(c int) int {
	f := (c - 32) * 5 / 9
	return f
}
