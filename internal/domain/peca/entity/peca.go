package entity

import (
	"time"

	"github.com/eltonCasacio/controle-estoque/pkg/entity"
)

type Peca struct {
	Id                entity.ID `json:"id"`
	IdFornecedor      string    `json:"id_fornecedor"`
	Codigo            string    `json:"codigo"`
	Descricao         string    `json:"descricao"`
	MateriaPrima      string    `json:"materia_prima"`
	UrlDesenhoTecnico string    `json:"url_desenho_tecnico"`
	UrlFoto           string    `json:"url_foto"`
	DescricaoTecnica  string    `json:"descricao_tecnica"`
	Massa             float64   `json:"massa"`
	Quantidade        int       `json:"quantidade"`
	Created_at        time.Time `json:"created_at"`
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
) (*Peca, error) {
	p := &Peca{
		Id:                entity.NewID(),
		IdFornecedor:      idFornecedor,
		Codigo:            codigo,
		Descricao:         descricao,
		MateriaPrima:      materiaPrima,
		UrlDesenhoTecnico: urlDesenhoTecnico,
		UrlFoto:           urlFoto,
		DescricaoTecnica:  descricaoTecnica,
		Massa:             massa,
		Quantidade:        quantidade,
		Created_at:        time.Now(),
	}

	err := IsValid(p)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func IsValid(p *Peca) error {
	return nil
}
