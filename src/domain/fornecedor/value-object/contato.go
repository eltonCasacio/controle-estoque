package fornecedores

type Contato struct {
	Telefone    string
	Email       string
	Celular     string
	NomeContato string
}

func NovoContato() *Contato {
	return &Contato{}
}
