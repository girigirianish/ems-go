package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/girigirianish/ems-go/config"
	_userHttpDelivery "github.com/girigirianish/ems-go/internal/user/delivery/http"
	_userRepo "github.com/girigirianish/ems-go/internal/user/repositories/mysql"
	_userUsecase "github.com/girigirianish/ems-go/internal/user/usecases"

	datastore "github.com/girigirianish/ems-go/pkg/database"
)

var (
	configPath = kingpin.Flag("config", "Location of config.yml").Default("./config.yml").String()
)

func main() {
	// Parse the CLI flags and load the config
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	// Load the config
	config.ReadConfig(*configPath)
	db, err := datastore.NewDB()
	if err != nil {
		log.Error(err)
	}
	defer func() {
		log.Info("Closing database connection")
		if err := db.Close(); err != nil {
			log.Error(err)
		}
	}()
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	userRepo := _userRepo.Init(db)
	userUsecase := _userUsecase.NewUserUseCaseImpl(
		userRepo,
		viper.GetString("auth.hash_salt"),
		[]byte(viper.GetString("auth.signing_key")),
		viper.GetDuration("auth.token_ttl"))
	_userHttpDelivery.RegisterHTTPEndpoints(e, userUsecase)

	go func() {
		if err := e.Start(viper.GetString("server.address")); err != nil {
			log.Info("Shutting down the server")
		}
	}()

	//Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
