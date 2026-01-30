package main

import (
	"os"
	"testing"
)

var app application

func TestMain(m *testing.M) {
	pathToTemplates = "../../templates/"

	app.Session = getSessions()
	os.Exit(m.Run())
}
