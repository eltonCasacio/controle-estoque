package usuario_entity

import (
	shared "github.com/eltonCasacio/controle-estoque/internal/domain/shared/entity"
	"github.com/eltonCasacio/controle-estoque/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

type UsuarioInterface interface {
	shared.Entity
}
type Usuario struct {
	Id    entity.ID `json:"id"`
	Nome  string    `json:"nome"`
	Senha string    `json:"senha"`
	Ativo bool      `json:"ativo"`
}

func NovoUsuario(nome, senha string) (*Usuario, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	u := &Usuario{
		Id:    entity.NewID(),
		Nome:  nome,
		Senha: string(hash),
		Ativo: true,
	}
	err = IsValid(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func IsValid(u *Usuario) error {
	return nil
}

func (u *Usuario) ValidarSenha(senha string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Senha), []byte(senha))
	return err == nil
}
