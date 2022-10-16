package shared_repository

import (
	e "github.com/eltonCasacio/controle-estoque/internal/domain/shared/entity"
	"github.com/eltonCasacio/controle-estoque/pkg/entity"
)

type RepositoryInterface[T e.Entity] interface {
	Create(t T) error
	Find(id entity.ID) (T, error)
	FindAll(t T) ([]T, error)
	Update(t T) error
	Delete(id entity.ID) error
}
