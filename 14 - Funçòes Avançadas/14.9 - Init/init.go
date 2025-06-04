package main

import "fmt"

// exacutada antes da main. Pode haver 1 por arquivo
func init() {
	fmt.Println("init")
}

func main() {
	fmt.Println("main")
}
