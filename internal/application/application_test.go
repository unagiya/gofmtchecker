package application

import "testing"

func TestRun(t *testing.T) {
	app := NewApplication("")
	if 0 != app.Run() {
		t.Error("exsec RunFunc is Faild")
	}
}