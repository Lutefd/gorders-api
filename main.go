package main

import (
	"context"

	"github.com/Lutefd/gorders-api/application"
)

func main() {
	app := application.NewApp()
	err := app.Start(context.TODO())
	if err != nil {
		panic(err)
	}
}
