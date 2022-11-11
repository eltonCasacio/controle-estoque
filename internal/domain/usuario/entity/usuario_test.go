package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNovoUsuario(t *testing.T) {
	u, err := NovoUsuario("Elton", "123")
	assert.Nil(t, err)
	assert.NotNil(t, u)
	assert.NotEmpty(t, u.id)
	assert.Equal(t, u.nome, "Elton")
	assert.NotEqual(t, u.senha, "123")
}

func TestUsuario_ValidarSenha(t *testing.T) {
	u, err := NovoUsuario("Elton", "123")
	assert.Nil(t, err)
	assert.True(t, u.ValidarSenha("123"))
	assert.False(t, u.ValidarSenha("1234"))
}
