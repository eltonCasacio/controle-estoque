package fornecedor

import (
	e "github.com/eltonCasacio/controle-estoque/src/domain/fornecedor/entity"
	r "github.com/eltonCasacio/controle-estoque/src/domain/shared/repository"
)

type FornecedorRepositoryInterface interface {
	r.RepositoryInterface[e.FornecedorInterface]
}
