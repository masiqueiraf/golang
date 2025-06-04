package main

import (
	"fmt"
	"time"
)

func main() {
	// CONCORRÊNCIA != PARALELISMO
	canal := make(chan string)
	go escrever("Olá mundo!", canal)

	for {
		mensagem, aberto := <-canal
		if !aberto {
			break
		}
		fmt.Println(mensagem)
	}
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}

	close(canal)
}
