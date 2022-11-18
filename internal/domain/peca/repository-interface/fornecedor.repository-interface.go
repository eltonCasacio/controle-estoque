package interfaces_peca

import (
	"github.com/eltonCasacio/controle-estoque/internal/domain/peca/entity"
	"github.com/eltonCasacio/controle-estoque/internal/domain/shared/repository"
)

type PecaRepositoryInterface interface {
	repository.RepositoryInterface[entity.Peca]
}
