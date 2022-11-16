package value_object

import (
	"errors"
)

var (
	CIDADE_OBRIGATORIO = "cidade é obrigatório"
	UF_OBRIGATORIO     = "uf é obrigatório"
	RUA_OBRIGATORIO    = "rua é obrigatório"
	BAIRRO_OBRIGATORIO = "bairro é obrigatório"
)

type Endereco struct {
	Cidade      string
	UF          string
	Rua         string
	Complemento string
	Bairro      string
	CEP         int
	Numero      string
}

func NovoEndereco(cidade, uf, rua, complemento, bairro, numero string, cep int) (*Endereco, error) {
	e := &Endereco{
		Cidade:      cidade,
		UF:          uf,
		Rua:         rua,
		Complemento: complemento,
		Bairro:      bairro,
		CEP:         cep,
		Numero:      numero,
	}
	err := e.ValidarEndereco()
	if err != nil {
		return nil, err
	}

	return e, nil
}
func (e *Endereco) ValidarEndereco() error {
	if e.Cidade == "" {
		return errors.New(CIDADE_OBRIGATORIO)
	}
	if e.UF == "" {
		return errors.New(UF_OBRIGATORIO)
	}
	if e.Rua == "" {
		return errors.New(RUA_OBRIGATORIO)
	}
	if e.Bairro == "" {
		return errors.New(BAIRRO_OBRIGATORIO)
	}
	return nil
}
