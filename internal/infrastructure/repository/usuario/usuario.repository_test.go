package usuario_repository

import (
	"testing"

	usuario_entity "github.com/eltonCasacio/controle-estoque/internal/domain/usuario/entity"
	"github.com/eltonCasacio/controle-estoque/pkg/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreate(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&usuario_entity.Usuario{})
	usuario, _ := usuario_entity.NovoUsuario("elton", "123")
	usuarioRepository := NovoUsuarioRpository(db)

	err = usuarioRepository.Create(usuario)
	assert.Nil(t, err)

	var usuarioEncontrado usuario_entity.Usuario
	err = db.First(&usuarioEncontrado, "nome = ?", usuario.Nome).Error
	assert.Nil(t, err)
	assert.Equal(t, usuarioEncontrado.Nome, usuario.Nome)
	assert.Equal(t, usuarioEncontrado.Ativo, true)
	assert.Equal(t, usuarioEncontrado.Id, usuario.Id)
	assert.NotEmpty(t, usuarioEncontrado.Senha)
}

func TestFindUsuario(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&usuario_entity.Usuario{})
	usuario, _ := usuario_entity.NovoUsuario("elton", "123")
	usuarioRepository := NovoUsuarioRpository(db)

	err = usuarioRepository.Create(usuario)
	assert.Nil(t, err)

	usuarioEncontrado, err := usuarioRepository.Find(usuario.Id)
	assert.Nil(t, err)
	assert.NotNil(t, usuarioEncontrado)
}

func TestFindUsuarioQueNaoExiste(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&usuario_entity.Usuario{})
	usuario, _ := usuario_entity.NovoUsuario("elton", "123")
	usuarioRepository := NovoUsuarioRpository(db)

	err = usuarioRepository.Create(usuario)
	assert.Nil(t, err)

	usuarioEncontrado, err := usuarioRepository.Find(entity.NewID())
	assert.Nil(t, usuarioEncontrado)
	assert.NotNil(t, err)
}
