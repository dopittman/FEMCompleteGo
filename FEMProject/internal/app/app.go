package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dopittman/femProject/internal/api"
)

type Application struct {
	Logger         *log.Logger
	WorkoutHandler *api.WorkoutHandler
}

func NewApplication() (*Application, error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// our stores will go here

	// our handlers will go here
	workoutHandler := api.NewWorkoutHandler()

	app := &Application{
		Logger:         logger,
		WorkoutHandler: workoutHandler,
	}

	return app, nil
}

// pass http.Request as a pointer to the function
// add HealthCheck to Application struct by adding (a *Application)
func (a *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	// fprint is specifically used to send data back data to the http client
	fmt.Fprint(w, "Status is available")
}
