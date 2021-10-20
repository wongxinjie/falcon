package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"falcon/config"
	"falcon/infra"
	"falcon/instance/loginst"
	"falcon/web/errorenum"
	"falcon/web/routers"
)

var router *gin.Engine

func init() {
	router = gin.New()
	router.HandleMethodNotAllowed = true

	router.NoRoute(routers.NoRouteHandler())
}

func setupInfra() error {
	err := config.LoadConfig()
	if err != nil {
		return err
	}

	err = infra.SetUp().
		WithDB().
		WithRedis().
		Build()
	if err != nil {
		return err
	}

	return err
}

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)

	if err := setupInfra(); err != nil {
		log.Fatal(err)
	}

	routers.SetRoutes(router)
	router.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorenum.UnknownError)
	}))

	s := &http.Server{
		Addr:         ":18001",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	loginst.Inst().Info("OK")

	go func() {
		if err := s.ListenAndServe(); err != nil {
			fmt.Println(err)
		}
	}()

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-ch

	if err := s.Shutdown(context.Background()); err != nil {
		fmt.Println(err)
	}
}
