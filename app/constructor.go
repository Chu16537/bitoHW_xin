package app

import (
	"bitohw_xin/app/config"
	"bitohw_xin/app/module/database"
	"bitohw_xin/app/server"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func Run() {

	c, err := config.Get("./env.yaml")

	if err != nil {
		fmt.Println(err)
		return
	}

	db := database.New()

	s := server.New(c.Server, db)

	s.Run()

	doneChan := make(chan os.Signal, 1)
	signal.Notify(doneChan, os.Interrupt, syscall.SIGTERM)
	<-doneChan

	s.Shutdown()
}
