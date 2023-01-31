package main

import (
	"example-go/app"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joinpickup/middleware-go/database"
	"github.com/joinpickup/middleware-go/logging"
	"github.com/joinpickup/middleware-go/support"
)

func main() {
	// setup inializations
	// envFile := ".env"

	// setup middleware
	// support.SetupEnv(true, &envFile)
	support.SetupEnv(false, nil)
	database.Init()
	logging.SetupLogging()

	r := app.NewRouter()
	port := os.Getenv("PORT")
	if port == "" {
		port = "6000"
	}

	// setup application
	// listen on port
	logging.Logger.Debug().Msg(fmt.Sprintf("Server running on port: %s\n", port))
	if support.Env != "prod" {
		logging.Logger.Debug().Msg("Development Server")
		log.Panic(http.ListenAndServeTLS(fmt.Sprintf(":%s", port), "./ssl/joinpickup-dev.cer.pem", "./ssl/joinpickup-dev.key.pem", r))
	} else {
		logging.Logger.Debug().Msg("Production Server")
		log.Panic(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
	}

	defer database.Close()
}
