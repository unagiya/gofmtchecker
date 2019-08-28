package main

import (
	"github.com/bikun-bikun/gofmtchecker/internal/application"
	"os"
)

func main() {
	app := application.NewApplication(" ")
	code := app.Run()
	os.Exit(code)
}