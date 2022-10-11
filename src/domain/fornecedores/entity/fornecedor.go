package fornecedores

import (
	"errors"

	. "github.com/eltonCasacio/controle-estoque/src/domain/fornecedores/value-object"
)

type FornecedorInterface interface {
	GetId() string
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
	id                string
	razaoSocial       string
	nomeFantasia      string
	cnpj              string
	inscricaoEstadual string
	endereco          Endereco
	contatos          []Contato
	idPecas           []string
	ativo             bool
}

func NovoFornecedor(
	nomeFantasia string,
	endereco Endereco,
	contatos []Contato,
	idPecas []string,
) *fornecedor {
	return &fornecedor{
		nomeFantasia: nomeFantasia,
		endereco:     endereco,
		contatos:     contatos,
		idPecas:      idPecas,
		ativo:        true,
	}
}

func (f *fornecedor) IsValid() (bool, error) {
	if len(f.contatos) == 0 {
		return false, errors.New("contato invalido")
	}

	if f.nomeFantasia == "" {
		return false, errors.New("nome invalido")
	}
	return true, nil
}

func (f *fornecedor) AtualizarContato(contatos []Contato) error {
	if len(contatos) == 0 {
		return errors.New("é obrigatorio pelo menos um contato")
	}
	f.contatos = contatos
	return nil
}

func (f *fornecedor) AtualizarPecas(idPecas []string) error {
	if len(idPecas) == 0 {
		return errors.New("é obrigatorio pelo menos uma peça")
	}
	f.idPecas = idPecas
	return nil
}

func (f *fornecedor) AtualizarEndereco(endereco Endereco) error {
	if len(endereco.Endereco) == 0 {
		return errors.New("endereço é obrigatorio")
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

func (f *fornecedor) GetId() string {
	return f.id
}
