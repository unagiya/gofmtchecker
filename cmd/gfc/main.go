package main

import (
	"github.com/bikun-bikun/gofmtchecker/internal/application"
	"os"
)

func main() {
	code := application.Run()
	os.Exit(code)
}