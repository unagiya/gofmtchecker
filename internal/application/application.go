package application

type app struct {
	flags string
}

func NewApplication(flag string) *app {
	return &app{ flag }
}

func (a *app) Run() (exitCode int) {
	exitCode = 0

	return
}