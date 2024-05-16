package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Carro struct {
	Nome   string `json:"nome"`
	Modelo string `json:"modelo"`
	Ano    int    `json:"ano"`
}

func (c Carro) Andar() {
	fmt.Println(c.Nome, "andando...")
}

func (c Carro) Parar() {
	fmt.Println(c.Nome, "parando...")
}

func main() {
	carro1 := Carro{"Fusca", "VW", 1970}
	carro2 := Carro{"Gol", "VW", 2000}
	carro3 := Carro{"Civic", "Honda", 2010}

	http.HandleFunc("/", home)

	http.ListenAndServe(":8080", nil)

	fmt.Println(carro1.Nome)
	fmt.Println(carro2.Nome)
	fmt.Println(carro3.Nome)

	carro1.Andar()
	carro2.Andar()
	carro3.Andar()

	carro1.Parar()
	carro2.Parar()
	carro3.Parar()
}

func home(w http.ResponseWriter, r *http.Request) {
	carro1 := Carro{"Fusquinha", "VW", 1969}
	json.NewEncoder(w).Encode(carro1)
}
