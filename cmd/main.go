package main

import (
	"apoteque/internal/api"
	"apoteque/store"
	"log"
)

func main() {
	cfg := api.ServerCfg{
		addr: ":8080",
		env:  "DEV",
		db: {
			addr: "localhost:5432",
		},
	}

	storage := store.NewStorage() // Your storage implementation
	app := internal.NewApplication(cfg, storage)
	err := app.run()
	if err != nil {
		log.Fatal(err)
	}

}
