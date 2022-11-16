package repository

type RepositoryInterface[T interface{}] interface {
	Criar(t *T) error
	BuscarPorID(id string) (*T, error)
	BuscarTodos() ([]T, error)
	Atualizar(t *T) error
	Excluir(id string) error
	BuscarPaginado(page, limit string, sort string) ([]T, error)
}
