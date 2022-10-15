package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"simple-crud-app/api"
	"simple-crud-app/internal/lib/config"
	"simple-crud-app/internal/lib/logger"
	"syscall"
)

var (
	configFile = "config.yml"
)

func init() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	configFile = fmt.Sprintf("%s\\%s", cwd, configFile)
}

func main() {
	// init logger
	l := logger.NewLogger().SetMethod("Server")

	// load service config
	cfg, err := config.NewConfig(configFile)
	if err != nil {
		l.Fatal(err)
	}
	l.Infof("Loaded config file: %s", configFile)

	// create a new http server
	server := api.NewHttpServer(cfg)

	// initiate errors description
	if err := server.InitErrors(cfg.ErrorsFile); err != nil {
		l.Fatal(err)
	}

	// connect to database (connection parameters in the config)
	if err := server.ConnectToDatabase(); err != nil {
		l.Fatal(err)
	}
	l.Infof("Connected to database: %s:%s", cfg.Sql.Host, cfg.Sql.Port)

	// run the http server
	l.Infof("Server listening: %s:%s", cfg.Api.Host, cfg.Api.Port)
	server.Run()

	// almost graceful shutdown
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-exit
	l.Infof("[Control-C] Get signal: shutdown server ...")
	l.Infof("Server shutting down")
	if err := server.Shutdown(context.TODO()); err != nil {
		l.Fatal(err)
	}
}
