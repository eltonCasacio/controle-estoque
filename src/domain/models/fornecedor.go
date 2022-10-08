package main

type endereco struct {
	Cidade      string
	Uf          string
	Endereco    string
	Complemento string
	Bairro      string
	Cep         int
	Numero      string
}

type contato struct {
	Telefone    string
	Email       string
	Celular     string
	NomeContato string
}

type Fornecedor struct {
	Id                int
	RazaoSocial       string
	Cnpj              string
	InscricaoEstadual string
	NomeFantasia      string
	Contato           contato
	Endereco          endereco
}
