package main

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"
	"webapp/pkg/data"

	"github.com/go-chi/chi/v5"
)

func Test_app_authenticate(t *testing.T) {
	var theTests = []struct {
		name               string
		requestBody        string
		expectedStatusCode int
	}{
		{"valid user", `{"email": "admin@example.com", "password": "secret"}`, http.StatusOK},
		{"not json", "I am not json", http.StatusUnauthorized},
		{"empty body", "{}", http.StatusUnauthorized},
		{"non existing user", `{"email": "user@example.com", "password": "secret"}`, http.StatusUnauthorized},
		{"wrong password", `{"email": "admin@example.com", "password": "password"}`, http.StatusUnauthorized},
	}

	for _, e := range theTests {
		var reader io.Reader
		reader = strings.NewReader(e.requestBody)

		req, _ := http.NewRequest("POST", "/auth", reader)
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(app.authenticate)
		handler.ServeHTTP(rr, req)

		if e.expectedStatusCode != rr.Code {
			t.Errorf("%s: returned wrong status code; expected %d but got %d", e.name, e.expectedStatusCode, rr.Code)
		}
	}
}

func Test_app_refresh(t *testing.T) {
	var theTests = []struct {
		name                    string
		token                   string
		expectedStatusCode      int
		resetRefreshTokenExpiry bool
	}{
		{"valid token", "", http.StatusOK, true},
		{"expired token", expiredToken, http.StatusBadRequest, false},
		{"valid, but not yet ready to expire", "", http.StatusTooEarly, false},
	}

	testUser := data.User{
		ID:       1,
		Username: "Admin",
		Email:    "admin@example.com",
	}

	oldRefreshTokenExpiry := refreshTokenExpiry

	for _, e := range theTests {
		var refreshToken string
		if e.token == "" {
			if e.resetRefreshTokenExpiry {
				refreshTokenExpiry = time.Second
			}
			tokenPair, _ := app.generateTokenPair(&testUser)
			refreshToken = tokenPair.RefreshToken
		} else {
			refreshToken = e.token
		}

		postedData := url.Values{"refresh_token": {refreshToken}}

		req, _ := http.NewRequest("POST", "/refresh-token", strings.NewReader(postedData.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(app.refresh)
		handler.ServeHTTP(rr, req)

		if e.expectedStatusCode != rr.Code {
			t.Errorf("%s: returned wrong status code; expected %d but got %d", e.name, e.expectedStatusCode, rr.Code)
		}

		refreshTokenExpiry = oldRefreshTokenExpiry
	}
}

func Test_app_userHandlers(t *testing.T) {
	var theTests = []struct {
		name           string
		method         string
		json           string
		paramID        string
		handler        http.HandlerFunc
		expectedStatus int
	}{
		{"allUsers", "GET", "", "", app.allUsers, http.StatusOK},
		{"deleteUser", "DELETE", "", "1", app.deleteUser, http.StatusNoContent},
		{"deleteUser - invalid param", "DELETE", "", "Y", app.deleteUser, http.StatusBadRequest},
		{"getUser - valid id", "GET", "", "1", app.getUser, http.StatusOK},
		{"getUser - invalid id param", "GET", "", "Y", app.getUser, http.StatusBadRequest},
		{"getUser - invalid id", "GET", "", "100", app.getUser, http.StatusNotFound},
		{
			"updateUser - valid id", "PATCH",
			`{"id": 1, "username": "Admin", "email": "admin@example.com"}`,
			"", app.updateUser, http.StatusNoContent,
		},
		{
			"updateUser - invalid id", "PATCH",
			`{"id": 100, "username": "Admin", "email": "admin@example.com"}`,
			"", app.updateUser, http.StatusInternalServerError,
		},
		{
			"updateUser - invalid json", "PATCH",
			`{"id": 1, username: "Admin", "email": "admin@example.com"}`,
			"", app.updateUser, http.StatusBadRequest,
		},
		{
			"insertUser - valid", "PUT",
			`{"username": "Admin", "email": "admin@example.com"}`,
			"", app.insertUser, http.StatusCreated,
		},
		{
			"insertUser - invalid body", "PUT",
			`{"username": "Admin", "email": "admin@example.com", "foo": "bar"}`,
			"", app.insertUser, http.StatusBadRequest,
		},
		{
			"insertUser - invalid json", "PUT",
			`{username: "Admin", "email": "admin@example.com" }`,
			"", app.insertUser, http.StatusBadRequest,
		},
	}

	for _, e := range theTests {
		var req *http.Request
		if e.json == "" {
			req, _ = http.NewRequest(e.method, "/", nil)
		} else {
			req, _ = http.NewRequest(e.method, "/", strings.NewReader(e.json))
		}

		if e.paramID != "" {
			chiCtx := chi.NewRouteContext()
			chiCtx.URLParams.Add("userID", e.paramID)
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
		}

		rr := httptest.NewRecorder()
		e.handler.ServeHTTP(rr, req)

		if rr.Code != e.expectedStatus {
			t.Errorf("%s: wrong status returned; expected %d but got %d", e.name, e.expectedStatus, rr.Code)
		}
	}
}
