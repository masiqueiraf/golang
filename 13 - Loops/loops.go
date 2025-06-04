package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Loops")

	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
		fmt.Println("Incrementando i")
		time.Sleep(time.Second)
	}

	for j := 0; j < 10; j++ {
		fmt.Println("Incrementando j", j)
		time.Sleep(time.Second)
	}

	nomes := [3]string{"João", "Davi", "Lucas"}
	fmt.Println(nomes[0])
	fmt.Println(nomes[1])
	fmt.Println(nomes[2])

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	// se não usar string() vai retornar o código da tabela ascii
	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Leonardo",
		"sobrenome": "Silva",
	}

	fmt.Println(usuario["nome"])

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

}
