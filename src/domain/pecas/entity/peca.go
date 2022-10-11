package entity

import "errors"

type PecaInterface interface {
	IsValid() bool
	GetCodigo() string
	AcrescentarQuantidade() int64
	DebitarQuantidade() int64
	Quantidade() int64
}
type peca struct {
	idFornecedor      string
	codigo            string
	descricao         string
	massa             float64
	materiaPrima      string
	urlDesenhoTecnico string
	urlFoto           string
	descricaoTecnica  string
	quantidade        int64
}

func NovaPeca(
	// idFornecedor string,
	codigo string,
	// descricao string,
	// massa float64,
	// materiaPrima string,
	// urlDesenhoTecnico string,
	// urlFoto string,
	// descricaoTecnica string,
) (*peca, error) {
	return &peca{
		// idFornecedor:      idFornecedor,
		codigo: codigo,
		// descricao:         descricao,
		// massa:             massa,
		// materiaPrima:      materiaPrima,
		// urlDesenhoTecnico: urlDesenhoTecnico,
		// urlFoto:           urlFoto,
		// descricaoTecnica:  descricaoTecnica,
	}, nil
}

func (p *peca) IsValid() bool {
	return true
}

func (p *peca) GetCodigo() string {
	return p.codigo
}

func (p *peca) AcrescentarQuantidade(valor int64) {
	p.quantidade = p.quantidade + valor
}

func (p *peca) DebitarQuantidade(valor int64) error {
	if (p.quantidade - valor) < 0 {
		return errors.New("não há quantidade suficiente")
	}
	p.quantidade = p.quantidade - valor
	return nil
}

func (p *peca) Quantidade() int64 {
	return p.quantidade
}
