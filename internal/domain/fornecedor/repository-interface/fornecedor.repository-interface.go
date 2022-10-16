package fornecedor_repository_interface

import (
	e "github.com/eltonCasacio/controle-estoque/internal/domain/fornecedor/entity"
	r "github.com/eltonCasacio/controle-estoque/internal/domain/shared/repository"
)

type FornecedorRepositoryInterface interface {
	r.RepositoryInterface[e.FornecedorInterface]
}
