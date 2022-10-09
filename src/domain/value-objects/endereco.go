package valueobjects

type Endereco struct {
	Cidade      string
	Uf          string
	Endereco    string
	Complemento string
	Bairro      string
	Cep         int
	Numero      string
}

func NovoEndereco() *Endereco {
	return &Endereco{}
}
