package entity

import (
	"errors"

	. "github.com/eltonCasacio/controle-estoque/src/domain/value-objects"
)

type FornecedorInterface interface {
	IsValid() (bool, error)
	Ativar() error
	Desativar() error
	RazaoSocial() string
	Cnpj() string
	InscricaoEstadual() string
	NomeFantasia() string
	Contato() Contato
	Endereco() Endereco
	AtualizarContato() error
	AtualizarPecas() error
}

type fornecedor struct {
	razaoSocial       string
	nomeFantasia      string
	cnpj              string
	inscricaoEstadual string
	endereco          Endereco
	ativo             bool
	contatos          []Contato
	pecas             []Peca
}

func NovoFornecedor(
	nomeFantasia string,
	endereco Endereco,
	contatos []Contato,
	pecas []Peca,
) *fornecedor {
	return &fornecedor{
		nomeFantasia: nomeFantasia,
		endereco:     endereco,
		contatos:     contatos,
		pecas:        pecas,
		ativo:        true,
	}
}

func (f *fornecedor) IsValid() (bool, error) {
	if len(f.contatos) == 0 {
		return false, errors.New("Contato invalido")
	}

	if f.nomeFantasia == "" {
		return false, errors.New("Nome invalido")
	}
	return true, nil
}

func (f *fornecedor) AtualizarContato(contatos []Contato) error {
	if len(contatos) == 0 {
		return errors.New("É obrigatorio pelo menos um contato")
	}
	f.contatos = contatos
	return nil
}

func (f *fornecedor) AtualizarPecas(pecas []Peca) error {
	if len(pecas) == 0 {
		return errors.New("É obrigatorio pelo menos uma peça")
	}
	f.pecas = pecas
	return nil
}

func (f *fornecedor) AtualizarEndereco(endereco Endereco) error {
	if len(endereco.Endereco) == 0 {
		return errors.New("Endereço é obrigatorio")
	}
	f.endereco = endereco
	return nil
}

func (f *fornecedor) Ativar() error {
	_, err := f.IsValid()
	if err != nil {
		return err
	}
	f.ativo = true
	return nil
}

func (f *fornecedor) Desativar() error {
	f.ativo = false
	return nil
}

func (f *fornecedor) RazaoSocial() string {
	return f.razaoSocial
}

func (f *fornecedor) Cnpj() string {
	return f.cnpj
}

func (f *fornecedor) InscricaoEstadual() string {
	return f.inscricaoEstadual
}

func (f *fornecedor) NomeFantasia() string {
	return f.nomeFantasia
}

func (f *fornecedor) Endereco() Endereco {
	return f.endereco
}

func (f *fornecedor) Contatos() []Contato {
	return f.contatos
}
