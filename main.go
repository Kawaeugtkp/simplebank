package main

import (
	"database/sql"
	"log"

	"github.com/Kawaeugtkp/simplebank/api"
	db "github.com/Kawaeugtkp/simplebank/db/sqlc"
	"github.com/Kawaeugtkp/simplebank/util"
	_ "github.com/lib/pq"
)

func main() {
	// main.goとapp.envが同じ階層にあるのでpath指定は"."でいい
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}
	// まずDBと接続をする
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	// 次にstoreとserverを作成する
	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start sever:", err)
	}
}