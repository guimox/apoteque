package api

import (
	"apoteque/store"
	"net/http"
)

type application struct {
	serverCfg ServerCfg
	store     store.Storage
}

type dbConfig struct {
	addr string
}

type ServerCfg struct {
	Addr string
	Env  string
	Db   dbConfig
}

func (app *application) run(mux http.Handler) error {
	server := http.Server{
		Addr:    app.serverCfg.Addr,
		Handler: app.routes(),
	}

	return server.ListenAndServe()
}
