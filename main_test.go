package main

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestWordcount(t *testing.T) {
	tests := []struct {
		name, in, want string
	}{
		{"empty", "", "0"},
		{"single", "ok", "1"},
		{"several", "go is awesome", "3"},
		{"even more", "php - not so much", "5"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			old, reader, writer := captureStdout()
			os.Args = []string{"main.go", test.in}
			main()
			got := restoreStdout(old, reader, writer)
			if got != test.want {
				t.Errorf("got '%s', want '%s'", got, test.want)
			}
		})
	}
}

// captureStdout redirects os.Stdout into the custom writer
func captureStdout() (old *os.File, reader *os.File, writer *os.File) {
	old = os.Stdout
	reader, writer, _ = os.Pipe()
	os.Stdout = writer
	return old, reader, writer
}

// restoreStdout returns everything printed since the last captureStdout()
// and restores the old os.Stdout
func restoreStdout(old *os.File, reader *os.File, writer *os.File) string {
	writer.Close()
	out, _ := ioutil.ReadAll(reader)
	os.Stdout = old
	return strings.TrimSuffix(string(out), "\n")
}
