package entity

import (
	"github.com/eltonCasacio/controle-estoque/pkg/entity"
)

type Peca struct {
	id                entity.ID
	codigo            string
	descricao         string
	materiaPrima      string
	urlDesenhoTecnico string
	urlFoto           string
	descricaoTecnica  string
	massa             float64
	quantidade        int
}

func NovaPeca(
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
		id:                entity.NewID(),
		codigo:            codigo,
		descricao:         descricao,
		materiaPrima:      materiaPrima,
		urlDesenhoTecnico: urlDesenhoTecnico,
		urlFoto:           urlFoto,
		descricaoTecnica:  descricaoTecnica,
		massa:             massa,
		quantidade:        quantidade,
	}

	err := p.IsValid()
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (p *Peca) GetID() entity.ID {
	return p.id
}

func (p *Peca) GetCodigo() string {
	return p.codigo
}

func (p *Peca) GetDescricao() string {
	return p.descricao
}

func (p *Peca) GetMateriaprima() string {
	return p.materiaPrima
}

func (p *Peca) GetUrlDesenhoTecnico() string {
	return p.urlDesenhoTecnico
}

func (p *Peca) GetUrlFoto() string {
	return p.urlFoto
}

func (p *Peca) GetDescricaoTecnica() string {
	return p.descricaoTecnica
}

func (p *Peca) GetMassa() float64 {
	return p.massa
}

func (p *Peca) GetQuantidade() int {
	return p.quantidade
}

func (p *Peca) IsValid() error {
	return nil
}

func (p *Peca) ChangeID(id entity.ID) {
	p.id = id
}

func (p *Peca) ChangeCodigo(codigo string) {
	p.codigo = codigo
}

func (p *Peca) ChangeDescricao(descricao string) {
	p.descricao = descricao
}

func (p *Peca) ChangeMateriaprima(materiaPrima string) {
	p.materiaPrima = materiaPrima
}

func (p *Peca) ChangeUrlDesenhoTecnico(urlDesenhoTecnico string) {
	p.urlDesenhoTecnico = urlDesenhoTecnico
}

func (p *Peca) ChangeUrlFoto(urlFoto string) {
	p.urlFoto = urlFoto
}

func (p *Peca) ChangeDescricaoTecnica(descricaoTecnica string) {
	p.descricaoTecnica = descricaoTecnica
}

func (p *Peca) ChangeMassa(massa float64) {
	p.massa = massa
}

func (p *Peca) ChangeQuantidade(quantidade int) {
	p.quantidade = quantidade
}
