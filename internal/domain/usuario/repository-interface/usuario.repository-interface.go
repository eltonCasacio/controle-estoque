package usuario_repository_interface

import (
	shared_repository "github.com/eltonCasacio/controle-estoque/internal/domain/shared/repository"
	usuario_entity "github.com/eltonCasacio/controle-estoque/internal/domain/usuario/entity"
)

type UsuarioRepositoryInterface interface {
	shared_repository.RepositoryInterface[usuario_entity.UsuarioInterface]
}
