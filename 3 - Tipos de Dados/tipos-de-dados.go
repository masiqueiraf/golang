package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int = 1000000000
	fmt.Println(numero)

	var numero2 int32 = 10000
	fmt.Println(numero2)

	// alias
	// int32 = RUNE
	var numero3 rune = 10000
	fmt.Println(numero3)

	// byte = UINT8
	var numero4 byte = 255
	fmt.Println(numero4)

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123
	fmt.Println(numeroReal2)

	numeroReal3 := 12345.67
	fmt.Println(numeroReal3)

	// FIM NUMEROS REAIS

	// STRINGS
	// go não tem char se usar aspas simples
	// retorna o numero na tabela ascii
	char := 'B'
	fmt.Println(char)

	// FIM STRINGS

	// toda variável tem algum valor inicial
	// string = ""
	// numero = 0
	// boolean = false
	var texto string
	fmt.Println(texto)

	// erro em go é um tipo
	var erro error
	fmt.Println(erro)

	var erro2 error = errors.New("Erro interno")
	fmt.Println(erro2)

}
