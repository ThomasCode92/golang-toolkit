package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"subscription-service/data"
)

var pageTests = []struct {
	name               string
	url                string
	method             string
	expectedStatusCode int
	handler            http.HandlerFunc
	sessionData        map[string]any
	expectedHTML       string
}{
	{
		name:               "home",
		url:                "/",
		method:             "GET",
		expectedStatusCode: http.StatusOK,
		handler:            testApp.HomePage,
	},
	{
		name:               "login",
		url:                "/login",
		method:             "GET",
		expectedStatusCode: http.StatusOK,
		handler:            testApp.LoginPage,
		expectedHTML:       `<h1 class="mt-5">Login</h1>`,
	},
	{
		name:               "logout",
		url:                "/logout",
		method:             "GET",
		expectedStatusCode: http.StatusSeeOther,
		handler:            testApp.Logout,
		sessionData: map[string]any{
			"userID": 1,
			"user":   data.User{},
		},
	},
}

func TestConfig(t *testing.T) {
	pathToTemplates = "./templates/"

	for _, e := range pageTests {
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(e.method, e.url, nil)

		ctx := getCtx(req)
		req = req.WithContext(ctx)

		if len(e.sessionData) > 0 {
			for key, value := range e.sessionData {
				testApp.Session.Put(ctx, key, value)
			}
		}

		e.handler.ServeHTTP(rr, req)

		if rr.Code != e.expectedStatusCode {
			t.Errorf("%s failed: expected %d, got %d", e.name, e.expectedStatusCode, rr.Code)
		}

		if len(e.expectedHTML) > 0 {
			html := rr.Body.String()
			if !strings.Contains(html, e.expectedHTML) {
				t.Errorf("%s failed: expected HTML to contain %s", e.name, e.expectedHTML)
			}
		}
	}
}

func TestConfig_PostLoginPage(t *testing.T) {
	pathToTemplates = "./templates/"

	postedData := url.Values{
		"email":    {"admin@example.com"},
		"password": {"password123"},
	}

	rr := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/login", strings.NewReader(postedData.Encode()))
	ctx := getCtx(req)
	req = req.WithContext(ctx)

	handler := http.HandlerFunc(testApp.PostLoginPage)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusSeeOther {
		t.Error("wrong status code")
	}

	if !testApp.Session.Exists(ctx, "userID") {
		t.Error("did not found userID in session")
	}
}
