package entity

import e "github.com/eltonCasacio/controle-estoque/src/domain/shared/entity"

type UsuarioInterface interface {
	e.Entity
	Nome() string
	Senha() string
}
type usuario struct {
	id    string
	nome  string
	senha string
}

func NovoUsuario() *usuario {
	return &usuario{}
}
func (u *usuario) Id() string {
	return u.id
}
func (u *usuario) IsValid() (bool, error) {
	return false, nil
}
func (u *usuario) Nome() string {
	return u.id
}
func (u *usuario) Senha() string {
	return u.id
}
