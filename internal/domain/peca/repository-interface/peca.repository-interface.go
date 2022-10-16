package peca_repository_interface

import (
	peca "github.com/eltonCasacio/controle-estoque/internal/domain/peca/entity"
	shared "github.com/eltonCasacio/controle-estoque/internal/domain/shared/repository"
)

type PecaRepositoryInterface interface {
	shared.RepositoryInterface[peca.PecaInterface]
}
