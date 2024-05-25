package main

import "fmt"

func HelloWorld() string {
	return `Hello World`
}

func main() {
	result:= HelloWorld()

	fmt.Printf(result)
}