package shared

import e "github.com/eltonCasacio/controle-estoque/src/domain/shared/entity"

type RepositoryInterface[T e.Entity] interface {
	Create(t T) error
	Find(id int64) (T, error)
	FindAll(t T) ([]T, error)
	Update(t T) error
	Delete(id int64) error
}
