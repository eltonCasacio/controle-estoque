package subconjunto

import (
	"errors"
)

type pecas struct {
	Id         string `json:"id"`
	Quantidade int64  `json:"quantidade"`
}
type subconjunto struct {
	Codigo           string  `json:"codigo"`
	Descricao        string  `json:"descricao"`
	Massa            float64 `json:"massa"`
	UrlFoto          string  `json:"url-foto"`
	DescricaoTecnica string  `json:"descricao-tecnica"`
	Pecas            []pecas
	Status           string `json:"status"`
	Quantidade       int64  `json:"quantidade"`
}

func NovoSubconjunto() *subconjunto {
	sc := &subconjunto{
		Codigo:           "",
		Descricao:        "",
		Massa:            0,
		UrlFoto:          "",
		DescricaoTecnica: "",
		Pecas:            nil,
		Status:           "",
		Quantidade:       0,
	}
	err := sc.IsValid()
	if err != nil {
		panic(err)
	}
	return sc
}

func (sc *subconjunto) IsValid() error {
	if sc.Massa <= 1 {
		return errors.New(MASSA_OBRIGATORIO)
	}
	if sc.Pecas == nil {
		return errors.New(PECAS_OBRIGATORIO)
	}
	return nil
}
func (s *subconjunto) GetID() string {
	return s.Codigo
}
