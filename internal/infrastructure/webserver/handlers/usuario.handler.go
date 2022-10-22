package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/eltonCasacio/controle-estoque/internal/domain/usuario/entity"
	database "github.com/eltonCasacio/controle-estoque/internal/infrastructure/database/usuario"
	"github.com/eltonCasacio/controle-estoque/internal/infrastructure/dto"
)

type Error struct {
	ErrorMessage string `json:"error-message"`
}

type UsuarioHandler struct {
	usuarioRepository database.UserRepositoryInterface
}

func NovoUsuarioHandler(repo database.UserRepositoryInterface) *UsuarioHandler {
	return &UsuarioHandler{usuarioRepository: repo}
}

func (h *UsuarioHandler) CriarUsuario(w http.ResponseWriter, r *http.Request) {
	var usuario dto.CriarUsuarioInput
	err := json.NewDecoder(r.Body).Decode(&usuario)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	u, err := entity.NovoUsuario(usuario.Nome, usuario.Senha)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{ErrorMessage: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	h.usuarioRepository.Criar(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{ErrorMessage: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
