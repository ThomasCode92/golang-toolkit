package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestConfig_AddDefaultData(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	ctx := getCtx(req)
	req = req.WithContext(ctx)

	testApp.Session.Put(ctx, "flash", "flash message")
	testApp.Session.Put(ctx, "warning", "warning message")
	testApp.Session.Put(ctx, "error", "error message")

	td := testApp.AddDefaultData(&TemplateData{}, req)

	if td.Flash != "flash message" {
		t.Errorf("failed to get flash message from session")
	}

	if td.Warning != "warning message" {
		t.Errorf("failed to get warning message from session")
	}

	if td.Error != "error message" {
		t.Errorf("failed to get error message from session")
	}
}

func TestConfig_IsAuthenticated(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	ctx := getCtx(req)
	req = req.WithContext(ctx)

	isAuth := testApp.IsAuthenticated(req)

	if isAuth {
		t.Error("returns true for unauthenticated user, when it should be false")
	}

	testApp.Session.Put(ctx, "userID", 1)
	isAuth = testApp.IsAuthenticated(req)

	if !isAuth {
		t.Error("returns false for authenticated user, when it should be true")
	}
}

func TestConfig_Render(t *testing.T) {
	pathToTemplates = "./templates"

	rr := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/", nil)
	ctx := getCtx(req)
	req = req.WithContext(ctx)

	testApp.render(rr, req, "home.page.gohtml", &TemplateData{})

	if rr.Code != 200 {
		t.Error("failed to render page")
	}
}
