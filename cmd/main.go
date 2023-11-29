package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/fillipehmeireles/user-service/adapters/pg"
	"github.com/fillipehmeireles/user-service/adapters/repositories"
	"github.com/fillipehmeireles/user-service/core/usecases"
	"github.com/fillipehmeireles/user-service/internal/config"
	"github.com/fillipehmeireles/user-service/pkg/handlers/user/api"
)

const ConfigPath = "./"

var (
	binding string
)

func init() {
	flag.StringVar(&binding, "httpbind", ":4000", "address/port to bind listen socket")
	flag.Parse()
}

func setupConfig() (*config.Config, error) {
	return config.NewConfig(ConfigPath)
}
func main() {
	config, err := setupConfig()
	if err != nil {
		log.Fatal(err)
	}

	pgInstance, err := pg.NewPGInstance(config.PgConfig.Dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer pgInstance.Close()
	pgInstance.Migrate()
	userRepo := repositories.NewUserRepository(pgInstance.DB)
	userUseCase := usecases.NewUserUseCase(userRepo)

	ws := new(restful.WebService)
	ws = ws.Path("/api")

	api.NewUserHandler(userUseCase, ws)
	restful.Add(ws)

	log.Println(binding)

	if err := http.ListenAndServe(binding, nil); err != nil {
		log.Fatal(err)
	}
}
