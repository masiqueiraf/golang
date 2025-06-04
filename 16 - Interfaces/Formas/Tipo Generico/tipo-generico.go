package main

import "fmt"

func generica(intef interface{}) {
	fmt.Println(intef)
}

func main() {
	generica("String")
	generica(7)
	generica(true)
}
