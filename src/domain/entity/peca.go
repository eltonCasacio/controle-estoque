package entity

type Peca struct {
	idFornecedor      string
	codigo            string
	descricao         string
	massa             float64
	materiaPrima      string
	urlDesenhoTecnico string
	urlFoto           string
	descricaoTecnica  string
}

func NovaPeca(
	idFornecedor string,
	codigo string,
	descricao string,
	massa float64,
	materiaPrima string,
	urlDesenhoTecnico string,
	urlFoto string,
	descricaoTecnica string) *Peca {
	return &Peca{
		idFornecedor:      idFornecedor,
		codigo:            codigo,
		descricao:         descricao,
		massa:             massa,
		materiaPrima:      materiaPrima,
		urlDesenhoTecnico: urlDesenhoTecnico,
		urlFoto:           urlFoto,
		descricaoTecnica:  descricaoTecnica,
	}
}
