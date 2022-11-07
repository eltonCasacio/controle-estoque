package dto

type GetJWTInput struct {
	Nome  string `json:"nome"`
	Senha string `json:"senha"`
}

type GetJWTOutput struct {
	AccessToken string `json:"access_token"`
}
