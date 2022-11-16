package entity

import (
	"errors"

	f "github.com/eltonCasacio/controle-estoque/internal/domain/fornecedor/value-object"
	"github.com/eltonCasacio/controle-estoque/pkg/entity"
)

type Fornecedor struct {
	id           entity.ID
	razaoSocial  string
	nomeFantasia string
	cnpj         string
	ie           string
	endereco     f.Endereco
	contatos     []f.Contato
	idPecas      []string
	ativo        bool
}

func NovoFornecedor(
	nomeFantasia string,
	endereco f.Endereco,
	contatos []f.Contato,
	idPecas []string,
) (*Fornecedor, error) {

	f := &Fornecedor{
		id:           entity.NewID(),
		razaoSocial:  "",
		nomeFantasia: nomeFantasia,
		cnpj:         "",
		ie:           "",
		endereco:     endereco,
		contatos:     contatos,
		idPecas:      idPecas,
		ativo:        true,
	}
	err := f.IsValid()
	if err != nil {
		return nil, err
	}

	return f, nil
}

func (f *Fornecedor) IsValid() error {
	if len(f.contatos) == 0 {
		return errors.New(CONTATO_OBRIGATORIO)
	}

	if f.nomeFantasia == "" {
		return errors.New(NOME_FANTASIA_OBRIGATORIO)
	}
	if len(f.idPecas) < 1 {
		return errors.New(PECA_OBRIGATORIO)
	}
	return nil
}

func (f *Fornecedor) AdicionarContato(contato f.Contato) error {
	if contato.ValidarContato() != nil {
		return errors.New(CONTATO_OBRIGATORIO)
	}
	f.contatos = append(f.contatos, contato)
	return nil
}

func (f *Fornecedor) RemoverContato(nome string) error {
	if len(f.contatos) == 1 {
		return errors.New(CONTATO_NAO_PODE_SER_REMOVIDO)
	}
	if nome == "" {
		return errors.New(NOME_CONTATO_OBRIGATORIO)
	}

	for k, v := range f.contatos {
		if v.Nome == nome {
			f.contatos[k] = f.contatos[0]
			f.contatos = f.contatos[1:]
		}
	}
	return nil
}

func (f *Fornecedor) AtualizarPecas(idPecas []string) error {
	if len(idPecas) == 0 {
		return errors.New(PECA_OBRIGATORIO)
	}
	f.idPecas = idPecas
	return nil
}

func (f *Fornecedor) AtualizarEndereco(endereco f.Endereco) error {
	err := endereco.ValidarEndereco()
	if err != nil {
		return errors.New(ENDERECO_OBRIGATORIO)
	}
	f.endereco = endereco
	return nil
}

func (f *Fornecedor) Ativar() error {
	err := f.IsValid()
	if err != nil {
		return err
	}
	f.ativo = true
	return nil
}

func (f *Fornecedor) Desativar() error {
	f.ativo = false
	return nil
}

func (f *Fornecedor) GetID() entity.ID {
	return f.id
}

func (f *Fornecedor) GetRazaoSocial() string {
	return f.razaoSocial
}

func (f *Fornecedor) GetNomeFantasia() string {
	return f.nomeFantasia
}

func (f *Fornecedor) GetCNPJ() string {
	return f.cnpj
}

func (f *Fornecedor) GetIe() string {
	return f.ie
}

func (f *Fornecedor) GetEndereco() f.Endereco {
	return f.endereco
}

func (f *Fornecedor) GetContatos() []f.Contato {
	return f.contatos
}

func (f *Fornecedor) GetIdPecas() []string {
	return f.idPecas
}

func (f *Fornecedor) IsAtivo() bool {
	return f.ativo
}

func (f *Fornecedor) ChangeID(id entity.ID) {
	f.id = id
}
