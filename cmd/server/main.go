package main

import (
	"net/http"

	"github.com/eltonCasacio/controle-estoque/configs"
	_ "github.com/eltonCasacio/controle-estoque/docs"
	"github.com/eltonCasacio/controle-estoque/internal/domain/usuario/entity"
	dbFornecedor "github.com/eltonCasacio/controle-estoque/internal/infrastructure/database/fornecedor"
	dbUser "github.com/eltonCasacio/controle-estoque/internal/infrastructure/database/usuario"
	"github.com/eltonCasacio/controle-estoque/internal/infrastructure/webserver/handlers"
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
	configs, err := configs.LoadConfig("./cmd/.env")
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
	r.Use(middleware.WithValue("jwt", configs.TokenAuth))
	r.Use(middleware.WithValue("jwtExperiesIn", configs.JwtExperesIn))

	usuarioDB := dbUser.NovoUsuarioRpository(db)
	usuarioHandler := handlers.NovoUsuarioHandler(usuarioDB)

	r.Route("/usuario", func(r chi.Router) {
		r.Post("/", usuarioHandler.CriarUsuario)
		r.Get("/", usuarioHandler.BuscarTodos)
		r.Get("/{id}", usuarioHandler.BuscarUsuarioPorID)
		r.Put("/{id}", usuarioHandler.Atualizar)
		r.Delete("/{id}", usuarioHandler.Excluir)
	})

	fornecedorDB := dbFornecedor.NovoFornecedorRepository(db)
	fornecedorHandler := handlers.NovoFornecedorHandler(fornecedorDB)

	r.Route("/fornecedor", func(r chi.Router) {
		r.Post("/", fornecedorHandler.CriarFornecedor)
	})

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))
	http.ListenAndServe(":8000", r)
}
