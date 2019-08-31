package application

import (
	"reflect"
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

func TestBuildParam(t *testing.T) {
	tests := []struct {
		name string
		ap   app
		want []string
	}{
		{
			name: "all Param",
			ap: app{
				true,
				true,
				true,
				"path",
				"binPath",
			},
			want: []string{"-s", "-l", "-d", "path"},
		},
		{
			name: "Path only",
			ap: app{
				false,
				false,
				false,
				"",
				"",
			},
			want: []string{""},
		},
		{
			name: "list true",
			ap: app{
				true,
				false,
				false,
				"",
				"",
			},
			want: []string{"-l", ""},
		},
		{
			name: "simple true",
			ap: app{
				false,
				true,
				false,
				"",
				"",
			},
			want: []string{"-s", ""},
		},
		{
			name: "display true",
			ap: app{
				false,
				false,
				true,
				"",
				"",
			},
			want: []string{"-d", ""},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.ap.buildParam()
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("want: %v, got: %v", tt.want, got)
			}
		})
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
