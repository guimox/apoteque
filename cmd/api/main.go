package main

import (
	"apoteque/config"
	"apoteque/env"
	"apoteque/internal/api"
	"apoteque/internal/db"
	"apoteque/store"
	"log"
)

func main() {
	cfg := config.ServerCfg{
		Addr: ":8080",
		Env:  "DEV",
		Db: config.DbCfg{
			Addr:         env.GetString("DB_ADDR", "postgres://user:adminpassword@localhost/social?sslmode=disable"),
			MaxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			MaxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			MaxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}

	db, err := db.NewConnection(cfg.Db.Addr, cfg.Db.MaxOpenConns, cfg.Db.MaxIdleConns, cfg.Db.MaxIdleTime)
	if err != nil {
		log.Panic(err)
	}

	defer db.Close()
	log.Println("database connection pool established")

	storage := store.NewStorage(db)
	server := api.NewServer(cfg)

	app := api.NewApplication(server, storage)

	log.Fatal(app.Run())
}
