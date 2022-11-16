package entity

import "github.com/eltonCasacio/controle-estoque/pkg/entity"

type UsuarioInterface interface {
	GetID() entity.ID
	GetNome() string
	GetSenha() string
	IsAtivo() bool
	Desativar()
	Ativar()
	ChangeID(id entity.ID)
	ChangeNome(nome string)
	ChangeSenha(senha string)
}
