package main

import (
	"os"
	"os/signal"
	"simple-crud-app/api"
	"simple-crud-app/internal/lib/config"
	"simple-crud-app/internal/lib/logger"
)

const (
	configFile = "config.yml"
)

func main() {
	// init logger
	l := logger.NewLogger().SetMethod("Init Server")
	// load service config
	cfg, err := config.NewConfig(configFile)
	if err != nil {
		l.Fatal(err)
	}
	l.Infof("Loaded config file: %s", configFile)
	// create a new http server
	server, err := api.NewHttpServer(cfg)
	if err != nil {
		l.Fatal(err)
	}
	// connect to database (connection parameters in the config)
	/*if err := server.ConnectToDatabase(); err != nil {
		l.Fatal(err)
	}
	l.Infof("Connected to database: %s:%s", cfg.Sql.Host, cfg.Sql.Port)*/
	// run the http server
	l.Infof("Server listening: %s:%s", cfg.Api.Host, cfg.Api.Port)
	if err := server.Run(); err != nil {
		l.Fatal(err)
	}
	// almost graceful shutdown
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt)
	<-exit
	l.Infof("[Control-C] Get signal: shutdown server ...")
	l.Infof("Server shutting down")
}
