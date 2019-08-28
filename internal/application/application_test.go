package application

import "testing"

func TestRun(t *testing.T) {
	if 0 != Run() {
		t.Error("exsec RunFunc is Faild")
	}
}