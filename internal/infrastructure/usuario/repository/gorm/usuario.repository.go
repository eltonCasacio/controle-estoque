package gorm_repository

import (
	"github.com/eltonCasacio/controle-estoque/internal/domain/usuario/entity"
	"gorm.io/gorm"
)

type UsuarioRepository struct {
	DB *gorm.DB
}

func NovoUsuarioRpository(db *gorm.DB) *UsuarioRepository {
	return &UsuarioRepository{DB: db}
}

func (u *UsuarioRepository) Criar(usuario *entity.Usuario) error {
	usuarioModel := ConvertUsuarioDomainToModel(usuario)
	return u.DB.Create(usuarioModel).Error
}

func (u *UsuarioRepository) BuscarPorID(id string) (*entity.Usuario, error) {
	var usuario Usuario
	if err := u.DB.Where("id = ?", id).First(&usuario).Error; err != nil {
		return nil, err
	}
	usuarioConvertido := ConvertUsuarioModelToDomain(&usuario)
	return usuarioConvertido, nil
}

func (u *UsuarioRepository) BuscarTodos() ([]entity.Usuario, error) {
	var usuarios []Usuario
	if err := u.DB.Find(&usuarios).Error; err != nil {
		return nil, err
	}
	var usuariosDomain []entity.Usuario
	for _, usuario := range usuarios {
		u_convertido := ConvertUsuarioModelToDomain(&usuario)
		usuariosDomain = append(usuariosDomain, *u_convertido)
	}
	return usuariosDomain, nil
}

func (u *UsuarioRepository) Atualizar(usuario *entity.Usuario) error {
	usuarioModel := ConvertUsuarioDomainToModel(usuario)
	_, err := u.BuscarPorID(usuario.Id.String())
	if err != nil {
		return err
	}
	return u.DB.Save(usuarioModel).Error
}

func (u *UsuarioRepository) Excluir(id string) error {
	usuario, err := u.BuscarPorID(id)
	if err != nil {
		return err
	}
	usuario.Ativo = false

	usuarioModel := ConvertUsuarioDomainToModel(usuario)
	return u.DB.Save(usuarioModel).Error
}

func (u *UsuarioRepository) BuscarPorNome(nome string) (*entity.Usuario, error) {
	var usuario Usuario
	if err := u.DB.Where("nome = ?", nome).First(&usuario).Error; err != nil {
		return nil, err
	}

	usuarioConvertido := ConvertUsuarioModelToDomain(&usuario)
	return usuarioConvertido, nil
}

func (u *UsuarioRepository) BuscarPaginado(page, limit int, sort string) ([]entity.Usuario, error) {
	var usuarios []Usuario
	var err error
	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}
	if page != 0 && limit != 0 {
		err = u.DB.Limit(limit).Offset((page - 1) * limit).Order("created_at " + sort).Where("deleted_at = null").Find(&usuarios).Error
	} else {
		err = u.DB.Order("created_at " + sort).Find(&usuarios).Error
	}

	var usuariosDomain []entity.Usuario
	for _, usuario := range usuarios {
		u_convertido := ConvertUsuarioModelToDomain(&usuario)
		usuariosDomain = append(usuariosDomain, *u_convertido)
	}
	return usuariosDomain, err
}
