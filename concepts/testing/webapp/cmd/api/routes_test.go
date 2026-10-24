package main

import (
	"net/http"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
)

func Test_app_Routes(t *testing.T) {
	registered := []struct {
		route  string
		method string
	}{
		{"/auth", "POST"},
		{"/refresh-token", "POST"},
		{"/users", "GET"},
		{"/users/{userID}", "GET"},
		{"/users/{userID}", "DELETE"},
		{"/users", "PUT"},
		{"/users/{userID}", "PATCH"},
	}

	mux := app.routes()
	chiRouter := mux.(chi.Routes) // cast to chi.Routes to access the routes

	for _, route := range registered {
		// check to see if the route exists
		if !routeExists(route.route, route.method, chiRouter) {
			t.Errorf("route %s is not registered", route.route)
		}
	}
}

func routeExists(testRoute, testMethod string, chiRoutes chi.Routes) bool {
	found := false

	_ = chi.Walk(chiRoutes, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		if strings.EqualFold(method, testMethod) && strings.EqualFold(strings.TrimSuffix(route, "/"), strings.TrimSuffix(testRoute, "/")) {
			found = true
		}
		return nil
	})

	return found
}
