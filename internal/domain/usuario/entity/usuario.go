package entity

import (
	"github.com/eltonCasacio/controle-estoque/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

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
type Usuario struct {
	id    entity.ID
	nome  string
	senha string
	ativo bool
}

func NovoUsuario(nome, senha string) (*Usuario, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &Usuario{
		id:    entity.NewID(),
		nome:  nome,
		senha: string(hash),
		ativo: true,
	}, nil
}

func (u *Usuario) ValidarSenha(senha string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.senha), []byte(senha))
	return err == nil
}

func (u *Usuario) GetID() entity.ID {
	return u.id
}
func (u *Usuario) GetNome() string {
	return u.nome
}
func (u *Usuario) GetSenha() string {
	return u.senha
}
func (u *Usuario) IsAtivo() bool {
	return u.ativo
}
func (u *Usuario) Desativar() {
	u.ativo = false
}
func (u *Usuario) Ativar() {
	u.ativo = true
}
func (u *Usuario) ChangeID(id entity.ID) {
	u.id = id
}
func (u *Usuario) ChangeNome(nome string) {
	u.nome = nome
}
func (u *Usuario) ChangeSenha(senha string) {
	u.senha = senha
}
