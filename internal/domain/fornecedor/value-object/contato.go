package value_object

import "errors"

var (
	NomeEObrigatorio = "Nome é obrigatório"
)

type Contato struct {
	Telefone string `json:"telefone"`
	Email    string `json:"email"`
	Celular  string `json:"celular"`
	Nome     string `json:"nome"`
}

func NovoContato(telefone, email, celular, nome string) (*Contato, error) {
	c := &Contato{
		Telefone: telefone,
		Email:    email,
		Celular:  celular,
		Nome:     nome,
	}
	err := c.ValidarContato()
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c *Contato) ValidarContato() error {
	if c.Nome == "" {
		return errors.New(NomeEObrigatorio)
	}
	return nil
}
