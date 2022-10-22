package repository

import (
	"github.com/eltonCasacio/controle-estoque/pkg/entity"
)

type RepositoryInterface[T interface{}] interface {
	Criar(t *T) error
	BuscarPorID(id entity.ID) (*T, error)
	BuscarTodos() ([]T, error)
	Atualizar(t *T) error
	Excluir(id entity.ID) error
}
