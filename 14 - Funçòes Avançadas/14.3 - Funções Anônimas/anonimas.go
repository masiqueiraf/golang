package main

import "fmt"

func main() {
	func() {
		fmt.Println("Olá Mundo 1!")
	}()

	func(texto string) {
		fmt.Println(texto)
	}("Olá Mundo 2!")

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Olá Mundo 3!")

	fmt.Println(retorno)
}
