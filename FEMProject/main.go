package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/dopittman/femProject/internal/app"
	"github.com/dopittman/femProject/internal/routes"
)

func main() {
	// adds flag that will use 8080 as default unless port is passed in
	var port int
	flag.IntVar(&port, "port", 8080, "go backend server port")
	flag.Parse()

	app, err := app.NewApplication()
	if err != nil {
		// panic is a "s*^t has hit the fan" error and should be used for when the app should crash when hitting this error basically
		panic(err)
	}

	// Pass HealthCheck as a first class citizen and setup route
	r := routes.SetupRoutes(app)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.Logger.Printf("We are running our app on port %d.\n", port)

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}
}
