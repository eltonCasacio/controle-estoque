package entity

import (
	value_object "github.com/eltonCasacio/controle-estoque/internal/domain/fornecedor/value-object"
	"github.com/eltonCasacio/controle-estoque/pkg/entity"
)

type FornecedorInterface interface {
	GetID() entity.ID
	GetRazaoSocial() string
	GetNomeFantasia() string
	GetCNPJ() string
	GetIe() string
	GetEndereco() value_object.Endereco
	GetContatos() []value_object.Contato
	IsValid() bool
	IsAtivo() bool
	Desativar()
	Ativar()
	ChangeID(id entity.ID)
	ChangeRazaoSocial() string
	ChangeNomeFantasia() string
	ChangeCNPJ() string
	ChangeIe() string
	ChangeEndereco() value_object.Endereco
	ChangeContatos() []value_object.Contato
}
