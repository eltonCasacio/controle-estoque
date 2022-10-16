package fornecedor_entity

import (
	"errors"

	f "github.com/eltonCasacio/controle-estoque/internal/domain/fornecedor/value-object"
)

type fornecedor struct {
	Id           string `json:"id"`
	RazaoSocial  string `json:"razao-social"`
	NomeFantasia string `json:"nome-fantasia"`
	CNPJ         string `json:"cnpj"`
	Ie           string `json:"ie"`
	Endereco     f.EnderecoInterface
	Contatos     []f.ContatoInterface
	IdPecas      []string `json:"id-pecas"`
	Ativo        bool     `json:"ativo"`
}

func NovoFornecedor(
	nomeFantasia string,
	endereco f.EnderecoInterface,
	contatos []f.ContatoInterface,
	idPecas []string,
) (*fornecedor, error) {

	f := &fornecedor{
		RazaoSocial:  "",
		NomeFantasia: nomeFantasia,
		CNPJ:         "",
		Ie:           "",
		Endereco:     endereco,
		Contatos:     contatos,
		IdPecas:      idPecas,
		Ativo:        true,
	}
	err := f.IsValid()
	if err != nil {
		return nil, err
	}

	return f, nil
}

func (f *fornecedor) IsValid() error {
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

func (f *fornecedor) GetID() string {
	return f.Id
}

func (f *fornecedor) AdicionarContato(contato f.ContatoInterface) error {
	if contato == nil {
		return errors.New(CONTATO_OBRIGATORIO)
	}
	f.Contatos = append(f.Contatos, contato)
	return nil
}

func (f *fornecedor) RemoverContato(nome string) error {
	if len(f.Contatos) == 1 {
		return errors.New(CONTATO_NAO_PODE_SER_REMOVIDO)
	}
	if nome == "" {
		return errors.New(NOME_CONTATO_OBRIGATORIO)
	}

	for k, v := range f.Contatos {
		if v.Nome() == nome {
			f.Contatos[k] = f.Contatos[0]
			f.Contatos = f.Contatos[1:]
		}
	}
	return nil
}

func (f *fornecedor) AtualizarPecas(idPecas []string) error {
	if len(idPecas) == 0 {
		return errors.New(PECA_OBRIGATORIO)
	}
	f.IdPecas = idPecas
	return nil
}

func (f *fornecedor) AtualizarEndereco(endereco f.EnderecoInterface) error {
	err := endereco.ValidarEndereco()
	if err != nil {
		return errors.New(ENDERECO_OBRIGATORIO)
	}
	f.Endereco = endereco
	return nil
}

func (f *fornecedor) Ativar() error {
	err := f.IsValid()
	if err != nil {
		return err
	}
	f.Ativo = true
	return nil
}

func (f *fornecedor) Desativar() error {
	f.Ativo = false
	return nil
}
