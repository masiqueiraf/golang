package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	cachorroEmJSON := `{"nome":"Bob","raca":"Labrador","idade":3}`
	var meuCachorro cachorro

	if erro := json.Unmarshal([]byte(cachorroEmJSON), &meuCachorro); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(meuCachorro)

	cachorro2EmJSON := `{"nome":"Toby","raca":"Poodle"}`
	cachorro2 := make(map[string]string)

	if erro := json.Unmarshal([]byte(cachorro2EmJSON), &cachorro2); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorro2)

}
