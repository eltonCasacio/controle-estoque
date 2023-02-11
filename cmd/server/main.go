package main

import (
	"database/sql"
	"net/http"

	"github.com/eltonCasacio/controle-estoque/configs"
	_ "github.com/eltonCasacio/controle-estoque/docs"
	usuario_repository "github.com/eltonCasacio/controle-estoque/internal/infrastructure/usuario/mysql"
	handlers "github.com/eltonCasacio/controle-estoque/internal/infrastructure/webserver/handlers/usuario-handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/go-sql-driver/mysql"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	configs, err := configs.LoadConfig(".env")
	if err != nil {
		panic(err)
	}
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.WithValue("jwt", configs.TokenAuth))
	r.Use(middleware.WithValue("jwtExperiesIn", configs.JwtExperesIn))

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/safisa")
	if err != nil {
		panic(err)
	}
	usuarioDB := usuario_repository.NovoUsuarioRpository(db)
	usuarioHandler := handlers.NovoUsuarioHandler(usuarioDB)

	r.Route("/usuario", func(r chi.Router) {
		r.Post("/", usuarioHandler.CriarUsuario)
		r.Get("/", usuarioHandler.BuscarTodos)
		r.Get("/{id}", usuarioHandler.BuscarUsuarioPorID)
		r.Put("/{id}", usuarioHandler.Atualizar)
		r.Delete("/{id}", usuarioHandler.Excluir)
		r.Get("/paginado", usuarioHandler.BuscarPaginado)
	})

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))
	http.ListenAndServe(":8000", r)
}
