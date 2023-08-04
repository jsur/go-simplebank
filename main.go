package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/techschool/simplebank/api"
	db "github.com/techschool/simplebank/db/sqlc"
	"github.com/techschool/simplebank/util"
)

func main() {
	config, configErr := util.LoadConfig(".")
	if configErr != nil {
		log.Fatal("cannot load config", configErr)
	}

	conn, dbOpenErr := sql.Open(config.DBDriver, config.DBSource)
	if dbOpenErr != nil {
		log.Fatal("cannot connect to db:", dbOpenErr)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server", err)
	}

	startErr := server.Start(config.ServerAddress)
	if startErr != nil {
		log.Fatal("cannot start server:", startErr)
	}
}
