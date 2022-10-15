package fornecedor

import (
	shared "github.com/eltonCasacio/controle-estoque/src/domain/shared/entity"
)

type FornecedorInterface interface {
	shared.Entity
	AtualizarContato() error
	AtualizarEndereco() error
	AtualizarPecas() error
	Ativar() error
	Desativar() error
}
