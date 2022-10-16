package subconjunto_repository_interface

import (
	shared "github.com/eltonCasacio/controle-estoque/internal/domain/shared/repository"
	subconjunto "github.com/eltonCasacio/controle-estoque/internal/domain/subconjunto/entity"
)

type SubconjuntoRepositoryInterface interface {
	shared.RepositoryInterface[subconjunto.SubconjuntoInterface]
}
