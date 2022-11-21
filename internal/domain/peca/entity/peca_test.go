package entity

import (
	"testing"

	"github.com/eltonCasacio/controle-estoque/pkg/entity"
	"github.com/stretchr/testify/assert"
)

func TestPeca_NovaPeca(t *testing.T) {
	id := entity.NewID()
	p, err := NovaPeca(id, "12cod", "descricao teste", "materia prima", "url desenho tecnico", "url foto", "descricao tecnica", 10.0, 5)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Equal(t, p.codigo, "12cod")
	assert.Equal(t, p.descricao, "descricao teste")
}
