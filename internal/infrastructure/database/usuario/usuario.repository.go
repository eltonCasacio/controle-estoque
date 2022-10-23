package database

import (
	e "github.com/eltonCasacio/controle-estoque/internal/domain/usuario/entity"
	"gorm.io/gorm"
)

type UsuarioRepository struct {
	DB *gorm.DB
}

func NovoUsuarioRpository(db *gorm.DB) *UsuarioRepository {
	return &UsuarioRepository{DB: db}
}

func (u *UsuarioRepository) Criar(usuario *e.Usuario) error {
	return u.DB.Create(usuario).Error
}

func (u *UsuarioRepository) BuscarPorID(id string) (*e.Usuario, error) {
	var usuario e.Usuario
	if err := u.DB.Where("id = ?", id).First(&usuario).Error; err != nil {
		return nil, err
	}
	return &usuario, nil
}

func (u *UsuarioRepository) BuscarTodos() ([]e.Usuario, error) {
	var usuarios []e.Usuario
	if err := u.DB.Find(&usuarios).Error; err != nil {
		return nil, err
	}
	return usuarios, nil
}

func (u *UsuarioRepository) Atualizar(usuario *e.Usuario) error {
	_, err := u.BuscarPorID(usuario.Id.String())
	if err != nil {
		return err
	}
	return u.DB.Save(usuario).Error
}

func (u *UsuarioRepository) Excluir(id string) error {
	usuario, err := u.BuscarPorID(id)
	if err != nil {
		return err
	}
	usuario.Ativo = false
	return u.DB.Save(usuario).Error
}

func (u *UsuarioRepository) BuscarPorNome(nome string) (*e.Usuario, error) {
	return &e.Usuario{}, nil
}

func (u *UsuarioRepository) BuscarPaginado(page, limit int, sort string) ([]e.Usuario, error) {
	var usuarios []e.Usuario
	var err error
	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}
	if page != 0 && limit != 0 {
		err = u.DB.Limit(limit).Offset((page - 1) * limit).Order("created_at " + sort).Find(&usuarios).Error
	} else {
		err = u.DB.Order("created_at " + sort).Find(&usuarios).Error
	}
	return usuarios, err
}
