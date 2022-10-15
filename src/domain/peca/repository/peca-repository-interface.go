package peca

import (
	peca "github.com/eltonCasacio/controle-estoque/src/domain/peca/entity"
	shared "github.com/eltonCasacio/controle-estoque/src/domain/shared/repository"
)

type PecaRepositoryInterface interface {
	shared.RepositoryInterface[peca.PecaInterface]
}
