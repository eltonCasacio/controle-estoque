package entity

import "github.com/eltonCasacio/controle-estoque/pkg/entity"

type SubconjuntoInterface interface {
	GetID() entity.ID
	GetCodigo() string
	GetDescricao() string
	GetMassa() float64
	GetUrlFoto() string
	GetDescricaoTecnica() string
	GetIDPecas() []string
	GetStatus() string
	GetQuantidade() string

	ChangeID() entity.ID
	ChangeCodigo() string
	ChangeDescricao() string
	ChangeMassa() string
	ChangeUrlFoto() string
	ChangeDescricaoTecnica() string
	ChangeIDPecas() []string
	ChangeStatus() string
	ChangeQuantidade() string

	IsValid()
}
