package main

import "fmt"

func func1() {
	fmt.Println("func1")
}
func func2() {
	fmt.Println("func2")
}

func main() {
	defer func1()
	defer func2()
	fmt.Println("main")
}
