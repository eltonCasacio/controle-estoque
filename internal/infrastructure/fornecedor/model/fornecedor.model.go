package model

import "github.com/eltonCasacio/controle-estoque/pkg/entity"

type FornecedorModel struct {
	Id           entity.ID
	RazaoSocial  string
	NomeFantasia string
	CNPJ         string
	Ie           string
	Ativo        bool
	Endereco     Endereco
	Contato      Contato
}
type Endereco struct {
	Cidade       string
	UF           string
	Rua          string
	Complemento  string
	Bairro       string
	CEP          int
	Numero       string
	FornecedorID entity.ID
}
type Contato struct {
	Id           int
	Telefone     string
	Email        string
	Celular      string
	Nome         string
	FornecedorID entity.ID
}
