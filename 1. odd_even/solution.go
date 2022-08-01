package main

import "fmt"

func main() {
	x, bo := half(2)
	fmt.Println(x, bo)
}

func half(x int) (int, bool) {

	return x / 2, x%2 == 0

	//if x%2 == 0 {
	//	return 1, true
	//}else{
	//	return 0, false
	//}
}

//func expression
//func main(){
//	hal := func(x int) (int, bool){
//		return x / 2, x%2 == 0
//	}
//	fmt.Println(hal(2))
//}
