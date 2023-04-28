package main

import (
	"github.com/t67y110v/driver/internal/driver/config"
	"github.com/t67y110v/driver/internal/driver/logging"
	"github.com/t67y110v/driver/internal/driver/server"
)

func main() {
	logging.Init()
	l := logging.GetLogger()
	l.Infoln("Config initialization")
	config, err := config.LoadConfig()
	if err != nil {

		l.Fatal(err)
	}

	l.Infoln("Starting grpc server on :%s", config.ServerPort)
	if err := server.Start(&config); err != nil {
		l.Fatal(err)
	}
}
