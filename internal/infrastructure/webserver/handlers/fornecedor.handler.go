package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/eltonCasacio/controle-estoque/internal/domain/fornecedor/entity"
	database "github.com/eltonCasacio/controle-estoque/internal/infrastructure/database/fornecedor"
	"github.com/eltonCasacio/controle-estoque/internal/infrastructure/dto"
)

type FornecedorHandler struct {
	FornecedorRepository database.FornecedorRepositoryInterface
}

func NovoFornecedorHandler(repo database.FornecedorRepositoryInterface) *FornecedorHandler {
	return &FornecedorHandler{FornecedorRepository: repo}
}

// Criar fornecedor godoc
// @Summary      Criar fornecedor
// @Description  Criar fornecedor
// @Tags         fornecedores
// @Accept       json
// @Produce      json
// @Param        request     body      dto.CriarFornecedorInput  true  "fornecedor request"
// @Success      201
// @Failure      500         {object}  Error
// @Router       /fornecedor [post]
func (h *FornecedorHandler) CriarFornecedor(w http.ResponseWriter, r *http.Request) {
	fornecedor := dto.CriarFornecedorInput{}
	err := json.NewDecoder(r.Body).Decode(&fornecedor)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	f, err := entity.NovoFornecedor(
		fornecedor.NomeFantasia,
		fornecedor.Endereco,
		fornecedor.Contatos,
		fornecedor.IdPecas,
	)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(f)
	err = h.FornecedorRepository.Criar(f)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *FornecedorHandler) BuscarTodos(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
func (h *FornecedorHandler) BuscarUsuarioPorID(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
func (h *FornecedorHandler) Atualizar(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
func (h *FornecedorHandler) Excluir(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
