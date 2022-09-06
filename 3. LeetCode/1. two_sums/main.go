package main

import "fmt"

func main() {

	var nums = []int{3, 2, 4}
	indices := twoSums(nums, 6)
	for _, x := range indices {
		fmt.Println(x)
	}

}

func twoSums(nums []int, target int) []int {
	m := make(map[int]int)

	for index, value := range nums {
		if v, found := m[target-value]; found {
			return []int{v, index}
		}
		m[value] = index
	}
	return nil
}
