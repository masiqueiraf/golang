package main

import "encoding/json"

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	c := cachorro{"Rex", "Dálmata", 3}
	cachorroEmJSON, erro := json.Marshal(c)
	if erro != nil {
		panic(erro)
	}
	println(string(cachorroEmJSON))

	c2 := map[string]string{
		"nome": "Toby",
		"raca": "Poodle",
	}
	cachorro2EmJSON, erro := json.Marshal(c2)
	if erro != nil {
		panic(erro)
	}
	println(string(cachorro2EmJSON))
}
