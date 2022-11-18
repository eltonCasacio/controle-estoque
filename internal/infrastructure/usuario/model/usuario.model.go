package model

import "github.com/eltonCasacio/controle-estoque/pkg/entity"

type UsuarioModel struct {
	Id    entity.ID
	Nome  string
	Senha string
	Ativo bool
}
