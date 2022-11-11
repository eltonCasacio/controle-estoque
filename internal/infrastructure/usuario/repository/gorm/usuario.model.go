package gorm_repository

import (
	e "github.com/eltonCasacio/controle-estoque/internal/domain/usuario/entity"
	"github.com/eltonCasacio/controle-estoque/pkg/entity"
)

type Usuario struct {
	Id    entity.ID `json:"id"`
	Nome  string    `json:"nome"`
	Senha string    `json:"_"`
	Ativo bool      `json:"ativo"`
}

func NovoUsuarioModel(entidadeUsuario *e.Usuario) *Usuario {
	return &Usuario{
		Id:    entidadeUsuario.Id,
		Nome:  entidadeUsuario.Nome,
		Senha: entidadeUsuario.Senha,
		Ativo: entidadeUsuario.Ativo,
	}
}
