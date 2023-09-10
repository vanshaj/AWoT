package main

import (
	"os"

	"github.com/vanshaj/awot/app"
	"github.com/vanshaj/awot/internal"
)

func main() {
	f, err := os.Create("debug.log")
	if err != nil {
		os.Exit(1)
	}
	internal.NewLogger("debug", f)
	if err := internal.NewAwsConfig(); err != nil {
		os.Exit(1)
	}
	if err := app.Run(); err != nil {
		os.Exit(1)
	}
}
