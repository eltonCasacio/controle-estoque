package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/eltonCasacio/controle-estoque/internal/domain/usuario/entity"
	database "github.com/eltonCasacio/controle-estoque/internal/infrastructure/database/usuario"
	"github.com/eltonCasacio/controle-estoque/internal/infrastructure/dto"
	pkg "github.com/eltonCasacio/controle-estoque/pkg/entity"
	"github.com/go-chi/chi"
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

func (h *UsuarioHandler) BuscarUsuarioPorID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	usuario, err := h.usuarioRepository.BuscarPorID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(usuario)
}

func (h *UsuarioHandler) BuscarTodos(w http.ResponseWriter, r *http.Request) {
	usuarios, err := h.usuarioRepository.BuscarTodos()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(usuarios)
}

func (h *UsuarioHandler) Atualizar(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	usuario := entity.Usuario{}
	err := json.NewDecoder(r.Body).Decode(&usuario)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	usuario.Id, err = pkg.ParseID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err = h.usuarioRepository.BuscarPorID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = h.usuarioRepository.Atualizar(&usuario)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *UsuarioHandler) Excluir(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err := h.usuarioRepository.BuscarPorID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = h.usuarioRepository.Excluir(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
