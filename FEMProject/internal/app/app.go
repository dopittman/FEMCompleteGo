package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dopittman/femProject/internal/api"
	"github.com/dopittman/femProject/internal/store"
	"github.com/dopittman/femProject/migrations"
)

type Application struct {
	Logger         *log.Logger
	WorkoutHandler *api.WorkoutHandler
	DB             *sql.DB
}

func NewApplication() (*Application, error) {
	pgDB, err := store.Open()
	if err != nil {
		return nil, err
	}

	err = store.MigrateFS(pgDB, migrations.FS, ".")
	if err != nil {
		panic(err)
	}

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// our stores will go here
	workoutStore := store.NewPostgresWorkoutStore(pgDB)

	// our handlers will go here
	workoutHandler := api.NewWorkoutHandler(workoutStore)

	app := &Application{
		Logger:         logger,
		WorkoutHandler: workoutHandler,
		DB:             pgDB,
	}

	return app, nil
}

// pass http.Request as a pointer to the function
// add HealthCheck to Application struct by adding (a *Application)
func (a *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	// fprint is specifically used to send data back data to the http client
	fmt.Fprint(w, "Status is available")
}
