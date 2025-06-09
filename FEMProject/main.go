package main

import (
	"github.com/dopittman/femProject/internal/app"
)

func main() {
	app, err := app.NewApplication()
	if err != nil {
		// panic is a "s*^t has hit the fan" error and should be used for when the app should crash when hitting this error basically
		panic(err)
	}

	app.Logger.Println("We are running our app.")
}
