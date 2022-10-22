package entity

import (
	"errors"
	"time"

	"github.com/eltonCasacio/controle-estoque/pkg/entity"
)

type Pecas struct {
	Id         entity.ID `json:"id"`
	Quantidade int64     `json:"quantidade"`
}
type Subconjunto struct {
	Id               entity.ID `json:"id"`
	Codigo           string    `json:"codigo"`
	Descricao        string    `json:"descricao"`
	Massa            float64   `json:"massa"`
	UrlFoto          string    `json:"url-foto"`
	DescricaoTecnica string    `json:"descricao-tecnica"`
	Pecas            []Pecas
	Status           string    `json:"status"`
	Quantidade       int64     `json:"quantidade"`
	Created_at       time.Time `json:"created_at"`
}

func NovoSubconjunto() *Subconjunto {
	sc := &Subconjunto{
		Id:               entity.NewID(),
		Codigo:           "",
		Descricao:        "",
		Massa:            0,
		UrlFoto:          "",
		DescricaoTecnica: "",
		Pecas:            nil,
		Status:           "",
		Quantidade:       0,
		Created_at:       time.Now(),
	}
	err := sc.IsValid()
	if err != nil {
		panic(err)
	}
	return sc
}

func (sc *Subconjunto) IsValid() error {
	if sc.Massa <= 1 {
		return errors.New(MASSA_OBRIGATORIO)
	}
	if sc.Pecas == nil {
		return errors.New(PECAS_OBRIGATORIO)
	}
	return nil
}
