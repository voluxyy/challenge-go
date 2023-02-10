package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("File name missing")
	}
	if len(os.Args) > 1 {
		if len(os.Args) > 2 {
			fmt.Println("Too many arguments")
		} else {
			arg := os.Args[1]
			ReadFile(arg)
		}
	}
}

func ReadFile(filename string) {
	body, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Can't read file")
	}
	fmt.Printf(string(body))
}
