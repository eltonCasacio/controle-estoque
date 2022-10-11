package shared

type Entity interface {
	GetId() string
	IsValid() (bool, error)
}

type RepositoryInterface[T Entity] interface {
	Create(t T) error
	Find(id int64) (T, error)
	FindAll(t T) ([]T, error)
	Atualizar(t T) error
	Delete(id int64) error
}
