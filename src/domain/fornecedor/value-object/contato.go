package fornecedor

import "errors"

var (
	NomeEObrigatorio = "Nome é obrigatório"
)

type ContatoInterface interface {
	ValidarContato() error
	Nome() string
}

type contato struct {
	Telefone    string `json:"telefone"`
	Email       string `json:"email"`
	Celular     string `json:"celular"`
	NomeContato string `json:"nome"`
}

func NovoContato(telefone, email, celular, nome string) (*contato, error) {
	c := &contato{
		Telefone:    telefone,
		Email:       email,
		Celular:     celular,
		NomeContato: nome,
	}
	err := c.ValidarContato()
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c *contato) ValidarContato() error {
	if c.NomeContato == "" {
		return errors.New(NomeEObrigatorio)
	}
	return nil
}

func (c *contato) Nome() string {
	return c.NomeContato
}
