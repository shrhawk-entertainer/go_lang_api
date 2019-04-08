package server

import "github.com/shrhawk-entertainer/go_lang_api/config"

func Init() {
	config := config.GetConfig()
	r := NewRouter()
	r.Run(config.GetString("server.port"))
}
