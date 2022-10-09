package entity

type Subconjunto struct {
	Codigo           string
	Descricao        string
	Massa            float64
	UrlFoto          string
	DescricaoTecnica string
}

func NovoSubconjunto() *Subconjunto {
	return &Subconjunto{}
}
