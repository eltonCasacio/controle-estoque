package entity

type Usuario struct {
	Nome  string
	Senha string
}

func NovoUsuario() *Usuario {
	return &Usuario{}
}
