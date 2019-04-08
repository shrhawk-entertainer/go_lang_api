package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/shrhawk-entertainer/go_lang_api/app_environments"
	"github.com/shrhawk-entertainer/go_lang_api/config"
	"github.com/shrhawk-entertainer/go_lang_api/db"
	"github.com/shrhawk-entertainer/go_lang_api/server"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	dbConnection := db.Init()
	app_environments.MigrateDb(dbConnection)
	defer dbConnection.Close()
	server.Init()
}
