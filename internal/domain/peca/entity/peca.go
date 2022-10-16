package peca_entity

type peca struct {
	IdFornecedor      string  `json:"id_fornecedor"`
	Codigo            string  `json:"codigo"`
	Descricao         string  `json:"descricao"`
	MateriaPrima      string  `json:"materia_prima"`
	UrlDesenhoTecnico string  `json:"url_desenho_tecnico"`
	UrlFoto           string  `json:"url_foto"`
	DescricaoTecnica  string  `json:"descricao_tecnica"`
	Massa             float64 `json:"massa"`
	Quantidade        int     `json:"quantidade"`
}

func NovaPeca(
	idFornecedor,
	codigo,
	descricao,
	materiaPrima,
	urlDesenhoTecnico,
	urlFoto,
	descricaoTecnica string,
	massa float64,
	quantidade int,
) (*peca, error) {
	p := &peca{
		IdFornecedor:      idFornecedor,
		Codigo:            codigo,
		Descricao:         descricao,
		MateriaPrima:      materiaPrima,
		UrlDesenhoTecnico: urlDesenhoTecnico,
		UrlFoto:           urlFoto,
		DescricaoTecnica:  descricaoTecnica,
		Massa:             massa,
		Quantidade:        quantidade,
	}

	err := p.IsValid()
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (p *peca) IsValid() error {
	return nil
}
func (p *peca) GetID() string {
	return p.Codigo
}
