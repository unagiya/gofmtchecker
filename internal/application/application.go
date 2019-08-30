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

//NewApplication ...
func NewApplication(list, simple, display bool, path, binPath string) *app {
	return &app{
list,
simple,
display,
		path,
		binPath,
	}
}

func (a *app) Run() (exitCode int) {
	exitCode = 0
	cmd := exec.Command(getGofmtPath(a.binPath), a.buildParam()...)
	out, err := cmd.Output()
	fmt.Println(string(out))
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

func (a *app) buildParam() (param []string) {
	if a.simple {
		param = append(param, "-s")
	}
	if a.list {
		param = append(param, "-l")
	}
	if a.display {
		param = append(param, "-d")
	}
	param = append(param, a.path)

	return
}

func getGofmtPath(binPath string) string {
	if len(binPath) > 0 {
		binPath = binPath + "/"
	}
	return fmt.Sprintf("%sgofmt", binPath)

}
