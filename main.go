package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/vsouza/go-gin-boilerplate/config"
	"github.com/vsouza/go-gin-boilerplate/db"
	"github.com/vsouza/go-gin-boilerplate/server"
	"github.com/vsouza/go-gin-boilerplate/app_environments"
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
