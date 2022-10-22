package entity

import (
	"errors"
	"time"

	f "github.com/eltonCasacio/controle-estoque/internal/domain/fornecedor/value-object"
	"github.com/eltonCasacio/controle-estoque/pkg/entity"
)

type Fornecedor struct {
	Id           entity.ID `json:"id"`
	RazaoSocial  string    `json:"razao_social"`
	NomeFantasia string    `json:"nome_fantasia"`
	CNPJ         string    `json:"cnpj"`
	Ie           string    `json:"ie"`
	Endereco     f.Endereco
	Contatos     []f.Contato
	IdPecas      []string  `json:"id_pecas"`
	Ativo        bool      `json:"ativo"`
	Created_at   time.Time `json:"created_at"`
}

func NovoFornecedor(
	nomeFantasia string,
	endereco f.Endereco,
	contatos []f.Contato,
	idPecas []string,
) (*Fornecedor, error) {

	f := &Fornecedor{
		Id:           entity.NewID(),
		RazaoSocial:  "",
		NomeFantasia: nomeFantasia,
		CNPJ:         "",
		Ie:           "",
		Endereco:     endereco,
		Contatos:     contatos,
		IdPecas:      idPecas,
		Ativo:        true,
		Created_at:   time.Now(),
	}
	err := f.IsValid()
	if err != nil {
		return nil, err
	}

	return f, nil
}

func (f *Fornecedor) IsValid() error {
	if len(f.Contatos) == 0 {
		return errors.New(CONTATO_OBRIGATORIO)
	}

	if f.NomeFantasia == "" {
		return errors.New(NOME_FANTASIA_OBRIGATORIO)
	}
	if len(f.IdPecas) < 1 {
		return errors.New(PECA_OBRIGATORIO)
	}
	return nil
}

func (f *Fornecedor) AdicionarContato(contato f.Contato) error {
	if contato.ValidarContato() != nil {
		return errors.New(CONTATO_OBRIGATORIO)
	}
	f.Contatos = append(f.Contatos, contato)
	return nil
}

func (f *Fornecedor) RemoverContato(nome string) error {
	if len(f.Contatos) == 1 {
		return errors.New(CONTATO_NAO_PODE_SER_REMOVIDO)
	}
	if nome == "" {
		return errors.New(NOME_CONTATO_OBRIGATORIO)
	}

	for k, v := range f.Contatos {
		if v.Nome == nome {
			f.Contatos[k] = f.Contatos[0]
			f.Contatos = f.Contatos[1:]
		}
	}
	return nil
}

func (f *Fornecedor) AtualizarPecas(idPecas []string) error {
	if len(idPecas) == 0 {
		return errors.New(PECA_OBRIGATORIO)
	}
	f.IdPecas = idPecas
	return nil
}

func (f *Fornecedor) AtualizarEndereco(endereco f.Endereco) error {
	err := endereco.ValidarEndereco()
	if err != nil {
		return errors.New(ENDERECO_OBRIGATORIO)
	}
	f.Endereco = endereco
	return nil
}

func (f *Fornecedor) Ativar() error {
	err := f.IsValid()
	if err != nil {
		return err
	}
	f.Ativo = true
	return nil
}

func (f *Fornecedor) Desativar() error {
	f.Ativo = false
	return nil
}
