package entity

import (
	"errors"

	"github.com/eltonCasacio/controle-estoque/pkg/entity"
)

type Subconjunto struct {
	id               entity.ID
	codigo           string
	descricao        string
	massa            float64
	urlFoto          string
	descricaoTecnica string
	idPecas          []string
	status           string
	quantidade       int64
}

func NovoSubconjunto() *Subconjunto {
	sc := &Subconjunto{
		id:               entity.NewID(),
		idPecas:          []string{""},
		codigo:           "",
		descricao:        "",
		massa:            0,
		urlFoto:          "",
		descricaoTecnica: "",
		status:           "",
		quantidade:       0,
	}
	err := sc.IsValid()
	if err != nil {
		panic(err)
	}
	return sc
}

func (sc *Subconjunto) IsValid() error {
	if sc.massa <= 1 {
		return errors.New(MASSA_OBRIGATORIO)
	}
	if sc.idPecas == nil {
		return errors.New(PECAS_OBRIGATORIO)
	}
	return nil
}

func (sc *Subconjunto) GetID() entity.ID {
	return sc.id
}

func (sc *Subconjunto) GetCodigo() string {
	return sc.codigo
}

func (sc *Subconjunto) GetDescricao() string {
	return sc.descricao
}

func (sc *Subconjunto) GetMassa() float64 {
	return sc.massa
}

func (sc *Subconjunto) GetUrlFoto() string {
	return sc.urlFoto
}

func (sc *Subconjunto) GetDescricaoTecnica() string {
	return sc.descricaoTecnica
}

func (sc *Subconjunto) GetIDPecas() []string {
	return sc.idPecas
}

func (sc *Subconjunto) GetStatus() string {
	return sc.status
}

func (sc *Subconjunto) GetQuantidade() int64 {
	return sc.quantidade
}

func (sc *Subconjunto) ChangeID(id entity.ID) {
	sc.id = id
}

func (sc *Subconjunto) ChangeCodigo(codigo string) {
	sc.codigo = codigo
}

func (sc *Subconjunto) ChangeDescricao(descricao string) {
	sc.descricao = descricao
}

func (sc *Subconjunto) ChangeMassa(massa float64) {
	sc.massa = massa
}

func (sc *Subconjunto) ChangeUrlFoto(urlFoto string) {
	sc.urlFoto = urlFoto
}

func (sc *Subconjunto) ChangeDescricaoTecnica(descTecnica string) {
	sc.descricaoTecnica = descTecnica
}

func (sc *Subconjunto) ChangeIDPecas(idPecas []string) {
	sc.idPecas = idPecas
}

func (sc *Subconjunto) ChangeStatus(status string) {
	sc.status = status
}

func (sc *Subconjunto) ChangeQuantidade(qtd int64) {
	sc.quantidade = qtd
}
