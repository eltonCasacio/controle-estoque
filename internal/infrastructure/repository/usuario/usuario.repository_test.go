package usuario_repository

import (
	"testing"

	usuario_entity "github.com/eltonCasacio/controle-estoque/internal/domain/usuario/entity"
	"github.com/eltonCasacio/controle-estoque/pkg/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetRepository() (*gorm.DB, *UsuarioRepository, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	db.AutoMigrate(&usuario_entity.Usuario{})
	repo := NovoUsuarioRpository(db)
	return db, repo, err
}

func TestCreate(t *testing.T) {
	db, repository, err := GetRepository()
	assert.Nil(t, err)
	usuario, _ := usuario_entity.NovoUsuario("roberto", "123")
	err = repository.Create(usuario)
	assert.Nil(t, err)

	var usuarioEncontrado usuario_entity.Usuario
	err = db.First(&usuarioEncontrado, "nome = ?", usuario.Nome).Error
	assert.Nil(t, err)
	assert.Equal(t, usuarioEncontrado.Nome, usuario.Nome)
	assert.Equal(t, usuarioEncontrado.Ativo, true)
	assert.Equal(t, usuarioEncontrado.Id, usuario.Id)
	assert.NotEmpty(t, usuarioEncontrado.Senha)
}

func TestFind(t *testing.T) {
	_, repository, err := GetRepository()
	assert.Nil(t, err)
	usuario, _ := usuario_entity.NovoUsuario("roberto", "123")
	_ = repository.Create(usuario)
	usuarioEncontrado, err := repository.Find(usuario.Id)
	assert.Nil(t, err)
	assert.NotNil(t, usuarioEncontrado)
	assert.Equal(t, usuarioEncontrado.Nome, usuario.Nome)
	assert.Equal(t, usuarioEncontrado.Id, usuario.Id)
}

func TestFindQuandoNaoExiste(t *testing.T) {
	_, repository, err := GetRepository()
	assert.Nil(t, err)
	usuarioEncontrado, err := repository.Find(entity.NewID())
	assert.Nil(t, usuarioEncontrado)
	assert.NotNil(t, err)
}

func TestFindAll(t *testing.T) {
	_, repository, err := GetRepository()
	assert.Nil(t, err)
	usuario1, _ := usuario_entity.NovoUsuario("elton", "123")
	usuario2, _ := usuario_entity.NovoUsuario("roberto", "123")
	usuario3, _ := usuario_entity.NovoUsuario("daniel", "123")

	_ = repository.Create(usuario1)
	_ = repository.Create(usuario2)
	_ = repository.Create(usuario3)

	usuarios, err := repository.FindAll()
	assert.Nil(t, err)
	assert.NotNil(t, usuarios)
	assert.Equal(t, len(*usuarios), 3)
}

func TestFindAll_NotFound(t *testing.T) {
	_, repository, err := GetRepository()
	assert.Nil(t, err)
	usuarios, err := repository.FindAll()
	assert.Nil(t, err)
	assert.NotNil(t, usuarios)
	assert.Equal(t, len(*usuarios), 0)
}

func TestUpdateUser(t *testing.T) {
	_, repository, err := GetRepository()
	assert.Nil(t, err)

	usuario, _ := usuario_entity.NovoUsuario("Elton", "123")
	_ = repository.Create(usuario)

	uEncontrado, _ := repository.Find(usuario.Id)
	uEncontrado.Nome = "Casacio"

	repository.Update(uEncontrado)
	usuarioAtualizado, err := repository.Find(usuario.Id)
	assert.Nil(t, err)
	assert.NotNil(t, usuarioAtualizado)
	assert.Equal(t, usuarioAtualizado.Nome, "Casacio")
}

func TestUpdateUserInvalid(t *testing.T) {
	_, repository, err := GetRepository()
	assert.Nil(t, err)
	err = repository.Update(&usuario_entity.Usuario{})
	assert.NotNil(t, err)
}

func TestDelete(t *testing.T) {
	_, repository, err := GetRepository()
	assert.Nil(t, err)

	usuario, _ := usuario_entity.NovoUsuario("Elton", "123")
	_ = repository.Create(usuario)

	repository.Delete(usuario.Id)
	assert.Nil(t, err)

	uEncontrado, err := repository.Find(usuario.Id)
	assert.Nil(t, err)
	assert.NotNil(t, uEncontrado)
	assert.Equal(t, uEncontrado.Nome, usuario.Nome)
	assert.Equal(t, uEncontrado.Ativo, false)
}

func TestDeleteWithIdInvalid(t *testing.T) {
	_, repository, err := GetRepository()
	assert.Nil(t, err)
	err = repository.Delete(entity.NewID())
	assert.NotNil(t, err)
}
