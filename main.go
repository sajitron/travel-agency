package main

import (
	"log"

	"github.com/sajitron/travel-agency/api"
	"github.com/sajitron/travel-agency/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config")
	}

	runGinServer(config)
}

func runGinServer(config util.Config) {
	server, err := api.NewServer(config)
	if err != nil {
		log.Fatal("cannot create server")
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot start server")
	}
}
