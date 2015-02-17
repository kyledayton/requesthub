package main

import(
	"regexp"
	"net/http"
)
type Route struct {
	pattern *regexp.Regexp
	handler http.Handler
}

type RegexpRouter struct {
	routes []*Route
}

func MakeRouter() *RegexpRouter {
	r := new(RegexpRouter)
	r.routes = make([]*Route, 0, 20)
	return r
}

func (r *RegexpRouter) Handler(matchRegexp string, handler http.Handler) {
	r.routes = append(r.routes, &Route{regexp.MustCompile(matchRegexp), handler})
}

func (r *RegexpRouter) HandleFunc(matchRegexp string, handlerFunc func(http.ResponseWriter, *http.Request)) {
	handler := http.HandlerFunc(handlerFunc)

	r.routes = append(r.routes, &Route{regexp.MustCompile(matchRegexp), handler})
}

func (r *RegexpRouter) ServeHTTP(resp http.ResponseWriter, req *http.Request) {

	for _, route := range r.routes {
		if route.pattern.MatchString(req.URL.Path) {
			route.handler.ServeHTTP(resp, req)
			return
		}
	}

	http.NotFound(resp, req)
}
