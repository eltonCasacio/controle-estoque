package fornecedor

import (
	"testing"

	e "github.com/eltonCasacio/controle-estoque/src/domain/fornecedor/value-object"
	"github.com/stretchr/testify/assert"
)

func MakeFornecedor() (*fornecedor, e.EnderecoInterface, []e.ContatoInterface, error) {
	endereco, _ := e.NovoEndereco("Cidade", "uf", "rua", "complemento", "bairro", "123", 12345678)
	contato, _ := e.NovoContato("telefone", "email", "celular", "elton")
	contatos := []e.ContatoInterface{contato}
	f, err := NovoFornecedor("nome fantasia", endereco, contatos, []string{"1"})
	return f, endereco, contatos, err
}

func TestNovoFornecedor(t *testing.T) {
	f, e, c, err := MakeFornecedor()
	assert.Nil(t, err)
	assert.NotNil(t, f)
	assert.NotNil(t, f.GetID())
	assert.Equal(t, f.NomeFantasia, "nome fantasia")
	assert.Equal(t, f.Endereco, e)
	assert.Equal(t, f.Contatos, c)
	assert.NotNil(t, f.NomeFantasia)
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
	f, err := NovoFornecedor("any_fantasy_name", endereco, []e.ContatoInterface{}, []string{})
	assert.Nil(t, f)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), CONTATO_OBRIGATORIO)
}

func TestAdicionarContato(t *testing.T) {
	f, _, c, err := MakeFornecedor()
	assert.Nil(t, err)
	assert.NotNil(t, f)

	contato, _ := e.NovoContato("1345234543", "email", "celular", "robert")
	c = append(c, contato)

	err = f.AdicionarContato(contato)
	assert.Nil(t, err)
	assert.Equal(t, f.Contatos, c)
	assert.Equal(t, len(f.Contatos), 2)
}

func TestAdicionarContatoWhenItsEmpty(t *testing.T) {
	f, _, _, err := MakeFornecedor()
	assert.Nil(t, err)

	err = f.AdicionarContato(nil)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), CONTATO_OBRIGATORIO)
}

func TestRemoverContatoWhitNameEmpty(t *testing.T) {
	f, _, _, err := MakeFornecedor()
	assert.Nil(t, err)

	contato, _ := e.NovoContato("1345234543", "email", "celular", "robert")
	_ = f.AdicionarContato(contato)

	err = f.RemoverContato("")
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), NOME_CONTATO_OBRIGATORIO)
}

func TestRemoveContato(t *testing.T) {
	f, _, _, err := MakeFornecedor()
	assert.Nil(t, err)

	contato, _ := e.NovoContato("1345234543", "email", "celular", "robert")
	_ = f.AdicionarContato(contato)
	assert.Equal(t, len(f.Contatos), 2)

	err = f.RemoverContato("robert")
	assert.Nil(t, err)
	assert.Equal(t, len(f.Contatos), 1)
}

func TestRemoveContatoWhenContactLengthIsOne(t *testing.T) {
	f, _, _, err := MakeFornecedor()
	assert.Nil(t, err)
	err = f.RemoverContato("elton")
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), CONTATO_NAO_PODE_SER_REMOVIDO)
}
