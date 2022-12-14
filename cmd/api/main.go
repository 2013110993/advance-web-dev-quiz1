//Filename: cmd/api/main.go
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

//The application version number
const version = "1.0.0"

//The conifiguration settings
type config struct {
	port int
	env  string //development,   staging, production, etc.
}

//Dependency Injection
type application struct {
	config config
	logger *log.Logger
}

func main() {
	// Declare an instance of the config struct.
	var cfg config
	//read in the flags tahat are need to populateouuuuur config
	flag.IntVar(&cfg.port, "port", 3000, "API server pport")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development | staging | production)")
	flag.Parse()
	//create a logger
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	//create an instance of our application struct
	app := &application{
		config: cfg,
		logger: logger,
	}

	//create our HTTP server
	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", cfg.port),
		//Custom routes
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	//Start our server
	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)

}
