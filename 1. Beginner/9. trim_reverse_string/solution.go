package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	r := Reverse("de75s1rev2er")
	fmt.Println(r)
}

func Reverse(str string) string {

	regex := regexp.MustCompile("[0-9]+")
	str1 := regex.ReplaceAllString(str, "")

	trimmed := strings.TrimSpace(str1)

	var result string
	for _, v := range trimmed {

		result = string(v) + result
	}
	return result
}
