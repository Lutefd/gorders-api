package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/Lutefd/gorders-api/application"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	app := application.NewApp(application.LoadConfig())
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	err := app.Start(ctx)
	if err != nil {
		panic(err)
	}
}
