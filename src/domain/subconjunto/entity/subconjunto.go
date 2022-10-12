package subconjuntos

import (
	"errors"
	"fmt"
)

type SubconjuntoInterface interface {
	AcrescentarQuantidade() int64
	DebitarQuantidade() int64
	Quantidade() int64
	AtualizarQuantidadePecasEstoque() error
}

type subconjunto struct {
	codigo           string
	descricao        string
	massa            float64
	urlFoto          string
	descricaoTecnica string
	pecas            map[string]int64
	montada          bool
	quantidade       int64
}

func NovoSubconjunto() *subconjunto {
	sc := &subconjunto{
		codigo:           "",
		descricao:        "",
		massa:            0,
		urlFoto:          "",
		descricaoTecnica: "",
		pecas:            nil,
		montada:          false,
		quantidade:       0,
	}
	err := sc.IsValid()
	if err != nil {
		panic(err)
	}

	err = sc.AtualizarQuantidadePecasEstoque()
	if err != nil {
		panic(err)
	}
	return sc
}

func (sc *subconjunto) IsValid() error {
	if sc.pecas == nil {
		return errors.New("pecas nao pode ser vazio")
	}
	return nil
}

func (sc *subconjunto) AtualizarQuantidadePecasEstoque() error {
	for idPeca, qtd := range sc.pecas {
		fmt.Println(idPeca, qtd)
	}
	// a quantidade de pecas é igual a quantidade de subconjunto vezes a quantidade de pecas que sao usadas para montar o subconjunto
	// calcular a quantidade de cada peca, pois o subconjunto pode ter 10 molas por exemplo, ou duas hastes inox
	// entao para debitar do estoque a quantidade de mola, teremos que multiplicar a quantidade de subconjunto por 10 molas
	return nil
}
