package database

import (
	"github.com/eltonCasacio/controle-estoque/internal/domain/usuario/entity"
	"github.com/eltonCasacio/controle-estoque/internal/infrastructure/shared/repository"
)

type UserRepositoryInterface interface {
	repository.RepositoryInterface[entity.Usuario]
}
