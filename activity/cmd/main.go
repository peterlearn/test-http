package main

import (
	"flag"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"os"
	"x-server/activity/internal/config"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

func newApp(hs *http.Server, gs *grpc.Server,

// r registry.Registrar
) (*kratos.App, func()) {
	return kratos.New(
			kratos.ID(id),
			kratos.Name(Name),
			kratos.Version(Version),
			kratos.Metadata(map[string]string{}),
			kratos.Server(
				hs,
				gs,
			),
			//kratos.Registrar(r),
		), func() {
		}
}

func main() {
	flag.Parse()
	if err := config.Init(); err != nil {
		panic(err)
	}

	app, cleanup, err := initApp()
	if err != nil {
		log.Error("initApp err:%v", err)
		return
	}
	defer cleanup()
	// start and wait for stop signal
	if err := app.Run(); err != nil {
		log.Error("app Run err:%v", err)
		return
	}
}
