package main

import (
	"net/http"

	"github.com/eltonCasacio/controle-estoque/configs"
	"github.com/eltonCasacio/controle-estoque/internal/domain/usuario/entity"
	database "github.com/eltonCasacio/controle-estoque/internal/infrastructure/database/usuario"
	"github.com/eltonCasacio/controle-estoque/internal/infrastructure/webserver/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	_, err := configs.LoadConfig("./cmd/.env")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("./cmd/safisa.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Usuario{})
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	usuarioDB := database.NovoUsuarioRpository(db)
	usuarioHandler := handlers.NovoUsuarioHandler(usuarioDB)

	r.Route("/usuario", func(r chi.Router) {
		r.Post("/", usuarioHandler.CriarUsuario)
		r.Get("/", usuarioHandler.BuscarTodos)
		r.Get("/{id}", usuarioHandler.BuscarUsuarioPorID)
		r.Put("/{id}", usuarioHandler.Atualizar)
		r.Delete("/{id}", usuarioHandler.Excluir)
	})

	http.ListenAndServe(":8000", r)
}
