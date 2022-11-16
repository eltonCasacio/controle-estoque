package entity

import (
	"testing"

	e "github.com/eltonCasacio/controle-estoque/internal/domain/fornecedor/value-object"
	"github.com/stretchr/testify/assert"
)

func MakeFornecedor() (*Fornecedor, e.Endereco, []e.Contato, error) {
	endereco, _ := e.NovoEndereco("Cidade", "uf", "rua", "complemento", "bairro", "123", 12345678)
	contato, _ := e.NovoContato("telefone", "email", "celular", "elton")
	contatos := []e.Contato{}
	contatos = append(contatos, *contato)
	f, err := NovoFornecedor("nome fantasia", *endereco, contatos, []string{"1"})
	return f, *endereco, contatos, err
}

func TestNovoFornecedor(t *testing.T) {
	f, e, c, err := MakeFornecedor()
	assert.Nil(t, err)
	assert.NotNil(t, f)
	assert.NotNil(t, f.id)
	assert.Equal(t, f.nomeFantasia, "nome fantasia")
	assert.Equal(t, f.endereco, e)
	assert.Equal(t, f.contatos, c)
	assert.NotNil(t, f.nomeFantasia)
}

func TestNovoFornecedorWhenNomeFatasiaIsEmpty(t *testing.T) {
	_, e, c, _ := MakeFornecedor()
	f, err := NovoFornecedor("", e, c, []string{"1"})
	assert.NotNil(t, err)
	assert.Nil(t, f)
	assert.Equal(t, err.Error(), NOME_FANTASIA_OBRIGATORIO)
}

func TestNovoFornecedorWhenIdsIsEmpty(t *testing.T) {
	_, e, c, _ := MakeFornecedor()
	f, err := NovoFornecedor("any_fantasy_name", e, c, []string{})
	assert.Nil(t, f)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), PECA_OBRIGATORIO)
}

func TestNovoFornecedorWhenContatoIsEmpty(t *testing.T) {
	_, endereco, _, _ := MakeFornecedor()
	f, err := NovoFornecedor("any_fantasy_name", endereco, []e.Contato{}, []string{})
	assert.Nil(t, f)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), CONTATO_OBRIGATORIO)
}

func TestAdicionarContato(t *testing.T) {
	f, _, c, err := MakeFornecedor()
	assert.Nil(t, err)
	assert.NotNil(t, f)

	contato, _ := e.NovoContato("1345234543", "email", "celular", "robert")
	c = append(c, *contato)

	err = f.AdicionarContato(*contato)
	assert.Nil(t, err)
	assert.Equal(t, f.contatos, c)
	assert.Equal(t, len(f.contatos), 2)
}

func TestAdicionarContatoWhenItsEmpty(t *testing.T) {
	f, _, _, err := MakeFornecedor()
	assert.Nil(t, err)

	err = f.AdicionarContato(e.Contato{})
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), CONTATO_OBRIGATORIO)
}

func TestRemoverContatoWhitNameEmpty(t *testing.T) {
	f, _, _, err := MakeFornecedor()
	assert.Nil(t, err)

	contato, _ := e.NovoContato("1345234543", "email", "celular", "robert")
	_ = f.AdicionarContato(*contato)

	err = f.RemoverContato("")
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), NOME_CONTATO_OBRIGATORIO)
}

func TestRemoveContato(t *testing.T) {
	f, _, _, err := MakeFornecedor()
	assert.Nil(t, err)

	contato, _ := e.NovoContato("1345234543", "email", "celular", "robert")
	_ = f.AdicionarContato(*contato)
	assert.Equal(t, len(f.contatos), 2)

	err = f.RemoverContato("robert")
	assert.Nil(t, err)
	assert.Equal(t, len(f.contatos), 1)
}

func TestRemoveContatoWhenContactLengthIsOne(t *testing.T) {
	f, _, _, err := MakeFornecedor()
	assert.Nil(t, err)
	err = f.RemoverContato("elton")
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), CONTATO_NAO_PODE_SER_REMOVIDO)
}
