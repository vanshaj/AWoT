package main

import (
	"os"

	"github.com/vanshaj/awot/app"
	"github.com/vanshaj/awot/internal"
)

func main() {
	internal.NewLogger("debug", os.Stdout)
	if err := app.Run(); err != nil {
		os.Exit(1)
	}
}
