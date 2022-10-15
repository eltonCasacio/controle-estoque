package subconjunto

import (
	shared "github.com/eltonCasacio/controle-estoque/src/domain/shared/repository"
	subconjunto "github.com/eltonCasacio/controle-estoque/src/domain/subconjunto/entity"
)

type SubconjuntoRepositoryInterface interface {
	shared.RepositoryInterface[subconjunto.SubconjuntoInterface]
}
