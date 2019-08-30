package application

import (
	"testing"
)

func TestRun(t *testing.T) {
	ta := &app{
		true,
		true,
		true,
		"../../.",
		"",
	}
	cd := ta.Run()
	if cd != 0 {
		t.Errorf("exsec RunFunc is Faild: exitCode = %v", cd)
	}
}

func TestGetGofmtPath(t *testing.T) {
	tests := []struct {
		name    string
		binPath string
		want    string
	}{
		{
			name:    "Blanc binPath",
			binPath: "",
			want:    "gofmt",
		},
		{
			name:    "set path/to",
			binPath: "path/to",
			want:    "path/to/gofmt",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getGofmtPath(tt.binPath)
			if tt.want != got {
				t.Errorf("want: %v, got: %v", tt.want, got)
			}
		})
	}
}
