package main

import(
	"regexp"
	"net/http"
	"log"
)
type Route struct {
	pattern *regexp.Regexp
	handler http.Handler
}

type RegexpRouter struct {
	routes []*Route
}

func MakeRouter() *RegexpRouter {
	return &RegexpRouter{make([]*Route, 10)}
}

func (r RegexpRouter) Handler(matchRegexp *regexp.Regexp, handler http.Handler) {
	r.routes = append(r.routes, &Route{matchRegexp, handler})
}

func (r RegexpRouter) HandleFunc(matchRegexp *regexp.Regexp, handlerFunc func(http.ResponseWriter, *http.Request)) {
	handler := http.HandlerFunc(handlerFunc)

	r.routes = append(r.routes, &Route{matchRegexp, handler})
	log.Printf("Added %s (n=%d)\n", matchRegexp.String(), len(r.routes))
}

func (r RegexpRouter) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	log.Println(req.URL.Path)
	log.Println(len(r.routes))

	for _, route := range r.routes {
		log.Printf("Serving request for %s", route.pattern.String())
		if route.pattern.MatchString(req.URL.Path) {
			route.handler.ServeHTTP(resp, req)
		}
	}

	http.NotFound(resp, req)
}
