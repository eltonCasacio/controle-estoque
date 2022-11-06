package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNovoUsuario(t *testing.T) {
	u, err := NovoUsuario("Elton", "123")
	assert.Nil(t, err)
	assert.NotNil(t, u)
	assert.NotEmpty(t, u.Id)
	assert.Equal(t, u.Nome, "Elton")
	assert.NotEqual(t, u.Senha, "123")
}

func TestUsuario_ValidarSenha(t *testing.T) {
	u, err := NovoUsuario("Elton", "123")
	assert.Nil(t, err)
	assert.True(t, u.ValidarSenha("123"))
	assert.False(t, u.ValidarSenha("1234"))
}
