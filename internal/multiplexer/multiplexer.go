package multiplexer

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

const (
	PARAMS = "params"
)

// Route contains data about route
type Route struct {
	Pattern string
	Method  string
	Handler http.HandlerFunc
}

// resolveURL resolve pattern
func (r *Route) resolveURL(urlPath string) (map[string]*string, error) {
	params := make(map[string]*string)

	reg := regexp.MustCompile(`{([^{}]*)}`)
	matches := reg.FindAllStringSubmatch(r.Pattern, -1)
	patternParts := strings.Split(r.Pattern, "/")
	urlParts := strings.Split(urlPath, "/")

	for _, match := range matches {
		params[match[1]] = nil
	}

	for i := range urlParts {
		findString := reg.FindStringSubmatch(patternParts[i])
		if len(findString) > 1 {
			if _, ok := params[findString[1]]; ok {
				params[findString[1]] = &urlParts[i]
				continue
			}
		}

		if urlParts[i] != patternParts[i] {
			return nil, errors.New("url does not match pattern")
		}

	}

	return params, nil
}

// Server serves as a route multiplexer
type Server struct {
	routes       []Route
	defaultRoute http.HandlerFunc
}

// NewServer construct new multiplexer object and set default values
func NewServer() *Server {
	srv := &Server{
		routes: nil,
		defaultRoute: func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, err := w.Write([]byte(fmt.Sprintf(`{"message": "%s method requested"}`, r.Method)))
			if err != nil {
				w.WriteHeader(http.StatusNoContent)
			}
		},
	}

	return srv
}

// ServeHTTP serve function
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Set the return Content-Type as JSON like before
	w.Header().Set("Content-Type", "application/json")

	for _, route := range s.routes {
		if len(route.Method) == 0 || route.Method == r.Method {
			params, err := route.resolveURL(r.URL.Path)
			if err == nil {

				// set params to the context
				ctx := context.WithValue(r.Context(), PARAMS, params)

				route.Handler(w, r.WithContext(ctx))
				return
			}
		}
	}

	s.defaultRoute(w, r)
}
