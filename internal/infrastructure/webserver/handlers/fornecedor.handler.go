package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/eltonCasacio/controle-estoque/internal/domain/fornecedor/entity"
	value_object "github.com/eltonCasacio/controle-estoque/internal/domain/fornecedor/value-object"
	database "github.com/eltonCasacio/controle-estoque/internal/infrastructure/database/fornecedor"
	"github.com/eltonCasacio/controle-estoque/internal/infrastructure/dto"
)

type FornecedorHandler struct {
	FornecedorRepository database.FornecedorRepositoryInterface
}

func NovoFornecedorHandler(repo database.FornecedorRepositoryInterface) *FornecedorHandler {
	return &FornecedorHandler{FornecedorRepository: repo}
}

func (h *FornecedorHandler) CriarFornecedor(w http.ResponseWriter, r *http.Request) {
	fornecedor := dto.CriarFornecedorInput{}
	err := json.NewDecoder(r.Body).Decode(&fornecedor)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	endereco, _ := value_object.NovoEndereco("Cidade", "uf", "rua", "complemento", "bairro", "123", 12345678)
	contato, _ := value_object.NovoContato("telefone", "email", "celular", "elton")
	contatos := []value_object.Contato{*contato}

	f, err := entity.NovoFornecedor(fornecedor.NomeFantasia, *endereco, contatos, []string{"1"})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.FornecedorRepository.Criar(f)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
