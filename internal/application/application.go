package application

import (
	"fmt"
	"os/exec"
)

type app struct {
	list    bool
	simple  bool
	display bool
	path    string
	binPath string
}

func NewApplication(flag bool) *app {
	return &app{flag}
}

func (a *app) Run() (exitCode int) {
	exitCode = 0
	cmd := exec.Command("godfmt", "s", "l", ".")
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("%v", err)
		exitCode = 2
		return
	}
	if len(out) != 0 {
		fmt.Print("Please take  product ")
		exitCode = 1
	}
	return
}
