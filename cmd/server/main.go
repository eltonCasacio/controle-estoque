package main

import (
	"net/http"

	"github.com/eltonCasacio/controle-estoque/configs"
	_ "github.com/eltonCasacio/controle-estoque/docs"
	usuario_repository "github.com/eltonCasacio/controle-estoque/internal/infrastructure/usuario/repository/gorm"
	handlers "github.com/eltonCasacio/controle-estoque/internal/infrastructure/webserver/handlers/usuario-handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// @title           API Controle de estoque
// @version         1.0
// @description     API para controle de estoque de peças
// @termsOfService  http://swagger.io/terms/
// @contact.name   	Elton Casacio & Wevyrton Antero
// @contact.url    	https://www.instagram.com/elton_casacio/
// @contact.email  	eltoncasacio@hotmail.com.br
// @license.name   	C3R Innovation
// @license.url    	https://c3rinnovation.com
// @host      		localhost:8000
// @BasePath  /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	configs, err := configs.LoadConfig(".env")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("./safisa.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&usuario_repository.Usuario{})

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.WithValue("jwt", configs.TokenAuth))
	r.Use(middleware.WithValue("jwtExperiesIn", configs.JwtExperesIn))

	usuarioDB := usuario_repository.NovoUsuarioRpository(db)
	usuarioHandler := handlers.NovoUsuarioHandler(usuarioDB)

	r.Route("/usuario", func(r chi.Router) {
		r.Post("/", usuarioHandler.CriarUsuario)
		r.Get("/", usuarioHandler.BuscarTodos)
		r.Get("/{id}", usuarioHandler.BuscarUsuarioPorID)
		r.Put("/{id}", usuarioHandler.Atualizar)
		r.Delete("/{id}", usuarioHandler.Excluir)
	})

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))
	http.ListenAndServe(":8000", r)
}
