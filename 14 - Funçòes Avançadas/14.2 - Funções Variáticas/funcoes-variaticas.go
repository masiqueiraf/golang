package main

import "fmt"

func soma(numeros ...int) int {
	resultado := 0
	for _, numero := range numeros {
		resultado += numero
	}
	return resultado
}

// parâmetro variático precisa ser o último parâmetro na função
func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	totalDaSoma := soma(1, 2, 3, 4, 5, 6, 200, 102, 12, 13)
	fmt.Println(totalDaSoma)

	escrever("Olá Mundo!", 10, 2, 3, 4, 5, 6, 7)
}
