package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/robfig/cron"
)

func main() {
	log.Println("scheduler starting")

	c := cron.New()

	c.AddFunc("*/1 * * * * *", func() {
		log.Println("current time", time.Now())
	})

	c.Start()
	log.Println("scheduler started")

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-ch

	c.Stop()
	log.Println("scheduler stop")
}
