package multiplexer

import (
	"net/http"
)

func (s *Server) AddRoute(pattern string, method string, handler http.HandlerFunc) {
	s.routes = append(s.routes, Route{
		Pattern: pattern,
		Method:  method,
		Handler: handler,
	})
}
