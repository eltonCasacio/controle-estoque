package fornecedor_entity

import (
	shared "github.com/eltonCasacio/controle-estoque/internal/domain/shared/entity"
)

type FornecedorInterface interface {
	shared.Entity
	AtualizarContato() error
	AtualizarEndereco() error
	AtualizarPecas() error
	Ativar() error
	Desativar() error
}
