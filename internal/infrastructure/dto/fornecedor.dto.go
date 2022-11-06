package dto

import value_object "github.com/eltonCasacio/controle-estoque/internal/domain/fornecedor/value-object"

type CriarFornecedorInput struct {
	NomeFantasia string                 `json:"nome_fantasia"`
	Endereco     value_object.Endereco  `json:"endereco"`
	Contatos     []value_object.Contato `json:"contatos"`
	IdPecas      []string               `json:"id_pecas"`
}
