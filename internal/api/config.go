package api

import (
	"apoteque/config"
	"apoteque/store"
	"net/http"
)

type application struct {
	server server
	store  store.Storage
}

type server struct {
	addr string
	db   dbConfig
	env  string
}

type dbConfig struct {
	addr         string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

func NewApplication(server *server, storage store.Storage) *application {
	return &application{
		server: *server,
		store:  storage,
	}
}

func NewServer(serverCfg config.ServerCfg) *server {
	return &server{
		addr: serverCfg.Addr,
		db: dbConfig{
			addr:         serverCfg.Db.Addr,
			maxOpenConns: serverCfg.Db.MaxOpenConns,
			maxIdleConns: serverCfg.Db.MaxIdleConns,
			maxIdleTime:  serverCfg.Db.MaxIdleTime,
		},
		env: serverCfg.Env,
	}
}

func (app *application) Run() error {
	server := http.Server{
		Addr:    app.server.addr,
		Handler: app.routes(),
	}

	return server.ListenAndServe()
}
