package entity

import (
	. "github.com/eltonCasacio/controle-estoque/src/domain/value-objects"
)

type Fornecedor struct {
	razaoSocial       string
	cnpj              string
	inscricaoEstadual string
	nomeFantasia      string
	contato           Contato
	endereco          Endereco
}

func NovoFornecedor(razaoSocial string, cnpj string, ie string, nomeFantasia string) *Fornecedor {
	return &Fornecedor{
		razaoSocial:       razaoSocial,
		cnpj:              cnpj,
		inscricaoEstadual: ie,
		nomeFantasia:      nomeFantasia,
	}
}
