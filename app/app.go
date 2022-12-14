package app

import (
	"github.com/gorilla/mux"
	"github.com/rajatxs/go-hyper/config"
	"github.com/rajatxs/go-hyper/db"
	"github.com/rajatxs/go-hyper/log"
	"github.com/rajatxs/go-hyper/rpc"
	"github.com/rajatxs/go-hyper/server"
	"github.com/rajatxs/go-hyper/util"
)

type App struct {
	server *server.Server
	db     *db.Database
}

// Returns new instance of App
func New() *App {
	return new(App)
}

// Registers RPC endpoint and starts webserver
func (a *App) startServer() (err error) {
	router := mux.NewRouter()
	rpch := rpc.New(a.db)

	if err = rpch.RegisterMethods(); err != nil {
		log.Fatalf("couldn't register rpc methods %s\n", err.Error())
	}

	router.Handle("/rpc", rpch.Handler())

	a.server = server.New(
		config.Hostname(),
		config.Port(),
		router)

	return a.server.Start()
}

// Opens global database
func (a *App) openDatabase() error {
	a.db = db.New(config.DbRoot())
	log.Infof("using database %s\n", a.db.FilePath())
	return a.db.Open()
}

// Starts app instance
func (a *App) Run() {
	util.Attempt(a.openDatabase())
	util.Attempt(a.startServer())
}

// Handles SIGINT and SIGTERM signal
func (a *App) Terminate() {
	util.Attempt(a.db.Close())
	util.Attempt(a.server.Stop())
}
