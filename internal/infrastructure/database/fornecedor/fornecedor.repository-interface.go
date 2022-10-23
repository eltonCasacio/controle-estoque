package database

import (
	"github.com/eltonCasacio/controle-estoque/internal/domain/fornecedor/entity"
	"github.com/eltonCasacio/controle-estoque/internal/infrastructure/shared/repository"
)

type FornecedorRepositoryInterface interface {
	repository.RepositoryInterface[entity.Fornecedor]
}
