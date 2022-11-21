package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPeca_NovaPeca(t *testing.T) {
	p, err := NovaPeca("12cod", "descricao teste", "materia prima", "url desenho tecnico", "url foto", "descricao tecnica", 10.0, 5)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Equal(t, p.codigo, "12cod")
	assert.Equal(t, p.descricao, "descricao teste")
}
