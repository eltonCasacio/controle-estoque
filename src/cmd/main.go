package main

import . "github.com/eltonCasacio/controle-estoque/src/domain/entity"

func main() {
	peca, err := NovaPeca("teste codigo")
	if err != nil {
		panic(err)
	}

	println(peca.GetCodigo())
}
