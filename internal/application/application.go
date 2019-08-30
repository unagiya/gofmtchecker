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
	return &app{flag, flag, flag, "", ""}
}

func (a *app) Run() (exitCode int) {
	exitCode = 0
	cmd := exec.Command(getGofmtPath(a.binPath), "-s", "-l", ".")
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

func getGofmtPath(binPath string) string {
	if len(binPath) > 0 {
		binPath = binPath + "/"
	}
	return fmt.Sprintf("%sgofmt", binPath)

}
