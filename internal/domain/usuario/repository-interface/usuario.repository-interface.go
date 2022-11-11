package usuario

import (
	"github.com/eltonCasacio/controle-estoque/internal/domain/shared/repository"
	"github.com/eltonCasacio/controle-estoque/internal/domain/usuario/entity"
)

type UsuarioRepositoryInterface interface {
	repository.RepositoryInterface[entity.Usuario]
	BuscarPorNome(nome string) (*entity.Usuario, error)
}
