package entity

import (
	"errors"

	value_object "github.com/eltonCasacio/controle-estoque/internal/domain/fornecedor/value-object"
	"github.com/eltonCasacio/controle-estoque/pkg/entity"
)

type Fornecedor struct {
	id           entity.ID
	razaoSocial  string
	nomeFantasia string
	cnpj         string
	ie           string
	endereco     value_object.Endereco
	contatos     []value_object.Contato
	idPecas      []string
	ativo        bool
}

func NovoFornecedor(
	nomeFantasia string,
	endereco value_object.Endereco,
	contatos []value_object.Contato,
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

func (f *Fornecedor) AdicionarContato(contato value_object.Contato) error {
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

func (f *Fornecedor) AtualizarEndereco(endereco value_object.Endereco) error {
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

func (f *Fornecedor) GetEndereco() value_object.Endereco {
	return f.endereco
}

func (f *Fornecedor) GetContatos() []value_object.Contato {
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

func (f *Fornecedor) ChangeRazaoSocial(value string) error {
	f.razaoSocial = value
	return nil
}

func (f *Fornecedor) ChangeNomeFantasia(value string) error {
	f.nomeFantasia = value
	return nil
}

func (f *Fornecedor) ChangeCNPJ(cnpj string) error {
	f.cnpj = cnpj
	return nil
}

func (f *Fornecedor) ChangeIe(ie string) error {
	f.ie = ie
	return nil
}

func (f *Fornecedor) ChangeEndereco(endereco value_object.Endereco) error {
	f.endereco = endereco
	return nil
}

func (f *Fornecedor) ChangeContatos(contatos []value_object.Contato) error {
	return nil
}

func (f *Fornecedor) ChangeIdPecas(ids []string) error {
	f.idPecas = ids
	return nil
}
