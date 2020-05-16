package main

import (
	"00_web_framework/server"

	"00_web_framework/router.go"

	"go.uber.org/fx"
)

//https://godoc.org/go.uber.org/fx
func main() {
	app := fx.New(
		// put all ingestion libs here
		fx.Provide(
			router.NewRouter,
		),
		// call the function to spin up the server
		fx.Invoke(server.NewServer),
	)
	// run calls the invoke functions
	app.Run()
}
