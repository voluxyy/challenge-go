package main

import "fmt"

func main() {
	a := -1
	b := 2
	r := 0
	var result bool
	if b < 0 {
		r = a - (-b)
	} else {
		r = a - b
	}
	if r > 0 {
		result = true
	} else {
		result = false
	}
	fmt.Println(result)
}
