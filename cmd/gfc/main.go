package main

import (
	"flag"
	"os"

	"github.com/bikun-bikun/gofmtchecker/internal/application"
)

func main() {
	var (
		list    bool
		simple  bool
		display bool
		path    string
		binPath string
	)
	flag.Parse()
	flag.BoolVar(&list, "l", false, "gofmt option -l")
	flag.BoolVar(&simple, "s", false, "gofmt option -s")
	flag.BoolVar(&display, "d", false, "gofmt option -d")
	flag.StringVar(&path, "p", ".", "gofmt path")
	flag.StringVar(&binPath, "b", "", "gofmt binPath")

	app := application.NewApplication(false)
	code := app.Run()
	os.Exit(code)
}
