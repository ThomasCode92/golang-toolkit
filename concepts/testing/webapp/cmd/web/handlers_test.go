package main

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_application_Handlers(t *testing.T) {
	theTests := []struct {
		name               string
		url                string
		expectedStatusCode int
	}{
		{"home", "/", http.StatusOK},
		{"not found", "/something", http.StatusNotFound},
	}

	routes := app.routes()

	// create a test server
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	// range through test data
	for _, e := range theTests {
		resp, err := ts.Client().Get(ts.URL + e.url)
		if err != nil {
			t.Log(err)
			t.Fatal(err)
		}

		if resp.StatusCode != e.expectedStatusCode {
			t.Errorf("for %s: expected status %d, but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
		}
	}
}

func Test_application_Home(t *testing.T) {
	// create a request
	req, _ := http.NewRequest("GET", "/", nil)
	req = addContextAndSession(req, app)

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(app.Home)
	handler.ServeHTTP(rr, req)

	// check the status code
	if rr.Code != http.StatusOK {
		t.Errorf("returned wrong status code; expected %d, but got %d", http.StatusOK, rr.Code)
	}

	body, _ := io.ReadAll(rr.Body)
	if !strings.Contains(string(body), "<small>From Session:") {
		t.Errorf("did not find correct text in html")
	}
}

func getCtx(req *http.Request) context.Context {
	ctx := context.WithValue(req.Context(), contextUserKey, "unknown")
	return ctx
}

func addContextAndSession(req *http.Request, app application) *http.Request {
	req = req.WithContext(getCtx(req))

	ctx, _ := app.Session.Load(req.Context(), req.Header.Get("X-Session"))

	return req.WithContext(ctx)
}
