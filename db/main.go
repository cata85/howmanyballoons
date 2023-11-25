package db

import (
	"log"
	"os"

	helper "github.com/cata85/balloons/helpers"
	"github.com/go-pg/pg"
)

var opts *(pg.Options)

func Connect() *pg.DB {
	if opts == nil {
		config := helper.Config()
		databaseConfig := config["database"]
		password := helper.GetPassword(
			os.Getenv(helper.String(helper.Get(databaseConfig, "key"))),
			os.Getenv(helper.String(helper.Get(databaseConfig, "iv"))),
			helper.String(helper.Get(databaseConfig, "password")))
		opts = new(pg.Options)
		opts.User = helper.String(helper.Get(databaseConfig, "username"))
		opts.Password = password
		opts.Addr = helper.String(helper.Get(databaseConfig, "host"))
	}
	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		log.Printf("Failed to connect to database.\n")
	}
	return db
}

func InitializeTables() {
	var db *pg.DB = Connect()
	defer db.Close()
	CreateBalloonObjectTable(db)
	return
}
