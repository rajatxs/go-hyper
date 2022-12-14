package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/rajatxs/go-hyper/app"
	"github.com/rajatxs/go-hyper/config"
)

func init() {
	config.Parse()
}

func main() {
	app := app.New()
	sigch := make(chan os.Signal, 1)

	go app.Run()
	signal.Notify(sigch, syscall.SIGINT, syscall.SIGTERM)

	<-sigch
	app.Terminate()
}
