package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"
	"webapp/pkg/data"
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
