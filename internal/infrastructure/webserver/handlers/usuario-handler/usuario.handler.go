package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/eltonCasacio/controle-estoque/internal/domain/usuario/entity"
	repository_interface "github.com/eltonCasacio/controle-estoque/internal/domain/usuario/repository-interface"
	"github.com/eltonCasacio/controle-estoque/internal/infrastructure/webserver/handlers"
	"github.com/eltonCasacio/controle-estoque/internal/infrastructure/webserver/handlers/usuario-handler/dto"
	pkg "github.com/eltonCasacio/controle-estoque/pkg/entity"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
)

type UsuarioHandler struct {
	usuarioRepository repository_interface.UsuarioRepositoryInterface
}

func NovoUsuarioHandler(repo repository_interface.UsuarioRepositoryInterface) *UsuarioHandler {
	return &UsuarioHandler{usuarioRepository: repo}
}

// GetJWT godoc
// @Summary      Gerar Token JWT
// @Description  Gerar Token JWT
// @Tags         usuarios
// @Accept       json
// @Produce      json
// @Param        request   			body     dto.GetJWTInput  true  "usuario credentials"
// @Success      200  {object}  	dto.GetJWTOutput
// @Failure      404  {object}  	Error
// @Failure      500  {object}  	Error
// @Router       /usuario/generate_token [post]
func (uh *UsuarioHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	jwtExperiesIn := r.Context().Value("jwtExperiesIn").(int)

	var usuario dto.GetJWTInput
	err := json.NewDecoder(r.Body).Decode(&usuario)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := handlers.Error{ErrorMessage: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	u, err := uh.usuarioRepository.BuscarPorID(usuario.Nome)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		error := handlers.Error{ErrorMessage: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	if !u.ValidarSenha(usuario.Senha) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	_, tokenString, _ := jwt.Encode(map[string]interface{}{
		"sub": u.Id.String(),
		"exp": time.Now().Add(time.Second * time.Duration(jwtExperiesIn)).Unix(),
	})
	accessToken := dto.GetJWTOutput{AccessToken: tokenString}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}

// Criar usuario godoc
// @Summary      Criar usuário
// @Description  Criar usuário
// @Tags         usuarios
// @Accept       json
// @Produce      json
// @Param        request     body      dto.CriarUsuarioInput  true  "usuario request"
// @Success      201
// @Failure      500         {object}  Error
// @Router       /usuario [post]
// @Security ApiKeyAuth
func (h *UsuarioHandler) CriarUsuario(w http.ResponseWriter, r *http.Request) {
	var usuario dto.CriarUsuarioInput
	err := json.NewDecoder(r.Body).Decode(&usuario)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	newUsuario, err := entity.NovoUsuario(usuario.Nome, usuario.Senha)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := handlers.Error{ErrorMessage: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	h.usuarioRepository.Criar(newUsuario)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := handlers.Error{ErrorMessage: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// BuscarUsuario godoc
// @Summary      Buscar Usuário
// @Description  Buscar Usuário por ID
// @Tags         usuarios
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "usuario ID" Format(uuid)
// @Success      200  {object}  entity.Usuario
// @Failure      404
// @Failure      500  {object}  Error
// @Router       /usuario/{id} [get]
// @Security ApiKeyAuth
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

// BuscarUsuario godoc
// @Summary      Buscar Usuários
// @Description  Buscar Todos Usuários
// @Tags         usuarios
// @Accept       json
// @Produce      json
// @Success      200  {array}  entity.Usuario
// @Failure      404
// @Failure      500  {object}  Error
// @Router       /usuario [get]
// @Security ApiKeyAuth
func (h *UsuarioHandler) BuscarTodos(w http.ResponseWriter, r *http.Request) {
	usuarios, err := h.usuarioRepository.BuscarTodos()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(usuarios)
}

// AtualizarUsuario godoc
// @Summary      Atualizar dados do usuário
// @Description  Atualizar dados do usuário
// @Tags         usuarios
// @Accept       json
// @Produce      json
// @Param        id        	path      string                  true  "usuario Id" Format(uuid)
// @Param        request    body      dto.CriarUsuarioInput     true  "usuario request"
// @Success      200
// @Failure      404
// @Failure      500       {object}  Error
// @Router       /usuario/{id} [put]
// @Security ApiKeyAuth
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

// ExcluirUsuario 	godoc
// @Summary      	Excluir Usuário
// @Description  	Excluir Usuário
// @Tags         	usuarios
// @Accept       	json
// @Produce      	json
// @Param        	id        		path      string   true  "usuario Id" Format(uuid)
// @Success      	200
// @Failure      	404
// @Failure      	500       		{object}  Error
// @Router       	/usuario/{id} 	[delete]
// @Security 		ApiKeyAuth
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
