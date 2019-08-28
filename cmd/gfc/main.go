package main

import (
	"github.com/bikun-bikun/gofmtchecker/internal/exec"
	"os"
)

func main() {
	code := exec.Run()
	os.Exit(code)
}