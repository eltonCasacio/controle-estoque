package entity

import (
	"github.com/eltonCasacio/controle-estoque/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

type Usuario struct {
	Id    entity.ID `json:"id"`
	Nome  string    `json:"nome"`
	Senha string    `json:"_"`
	Ativo bool      `json:"ativo"`
}

func NovoUsuario(nome, senha string) (*Usuario, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &Usuario{
		Id:    entity.NewID(),
		Nome:  nome,
		Senha: string(hash),
		Ativo: true,
	}, nil
}

func (u *Usuario) ValidarSenha(senha string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Senha), []byte(senha))
	return err == nil
}
