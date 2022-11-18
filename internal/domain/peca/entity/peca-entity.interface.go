package entity

import "github.com/eltonCasacio/controle-estoque/pkg/entity"

type PecaInterface interface {
	GetID() entity.ID
	GetIDFornecedor() string
	GetCodigo() string
	GetDescricao() string
	GetMateriaprima() string
	GetUrlDesenhoTecnico() string
	GetUrlFoto() string
	GetDescricaoTecnica() string
	GetMassa() string
	GetQuantidade() int
	IsValid() bool

	ChangeID() entity.ID
	ChangeIDFornecedor() string
	ChangeCodigo() string
	ChangeDescricao() string
	ChangeMateriaprima() string
	ChangeUrlDesenhoTecnico() string
	ChangeUrlFoto() string
	ChangeDescricaoTecnica() string
	ChangeMassa() string
	ChangeQuantidade() int
}
