package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/tmortx7/go-simplebank/api"
	"github.com/tmortx7/go-simplebank/config"
	db "github.com/tmortx7/go-simplebank/db/sqlc"
)

func main() {
	config.ReadConfig(config.ReadConfigOption{})
	
	conn, err := sql.Open(config.C.Database.DBDriver, config.C.Database.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.C.Server.Address)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
