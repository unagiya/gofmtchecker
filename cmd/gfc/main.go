package main

import (
	"os"

	"github.com/bikun-bikun/gofmtchecker/internal/application"
)

func main() {
	app := application.NewApplication()
	code := app.Run()
	os.Exit(code)
}
