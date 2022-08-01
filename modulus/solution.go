package main

import "fmt"

func main(){
	x , bo := half(2)
	fmt.Println(x,bo)
}

func half(x int) (int, bool){
	if x%2 == 0 {
		return 1, true
	}else{
		return 0, false
	}
}
