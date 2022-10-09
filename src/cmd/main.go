package main

import (
	entities "github.com/eltonCasacio/controle-estoque/src/domain/entity"
)

func main() {
	fornecedor := entities.NovoFornecedor("", "", "", "")
	println(fornecedor)
}
