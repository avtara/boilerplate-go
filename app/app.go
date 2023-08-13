package app

import (
	"log"
	"os"
)

type App struct {
	Hostname string
}

func New() App {
	var cfg App

	cfg.Hostname, _ = os.Hostname()
	most(cfg.InitViper())
	most(cfg.InitLogrus())

	return cfg
}

func (cfg *App) Start() (err error) {
	ch := make(chan bool)
	go func() {

		ch <- false
	}()

	<-ch
	return
}

func most(err error) {
	if err != nil {
		log.Panic(err)
	}
}
