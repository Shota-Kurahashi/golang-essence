package main

import (
	"fmt"
	"net/http"
	"regexp"
)

type routerParam map[string]string

type routerFunc func(routerParam, http.ResponseWriter, *http.Request)

type routerItem struct {
	method  string
	matcher *regexp.Regexp
	fnc     routerFunc
}

type router struct {
	items []routerItem
}

func (rt *router) GET(prefix string, fnc routerFunc) {
	rt.items = append(rt.items, routerItem{method: http.MethodGet, matcher: regexp.MustCompile(prefix), fnc: fnc})
}

func (rt *router) POST(prefix string, fnc routerFunc) {
	rt.items = append(rt.items, routerItem{method: http.MethodPost, matcher: regexp.MustCompile(prefix), fnc: fnc})
}

// ServerHTTP は、http.Handler インターフェースを実装する。
func (rt *router) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	for _, v := range rt.items {
		if v.method == r.Method && v.matcher.MatchString(r.RequestURI) {
			match := v.matcher.FindStringSubmatch(r.RequestURI)
			param := make(routerParam)

			for i, name := range v.matcher.SubexpNames() {
				param[name] = match[i]
			}

			v.fnc(param, w, r)
			return
		}
	}

	http.NotFound(w, r)
}

func main() {
	rt := router{}

	rt.GET(`^/$`, func(param routerParam, w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	rt.POST(`^/hello/(?P<name>.+)$`, func(param routerParam, w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s!", param["name"])
	})

	http.ListenAndServe(":8080", &rt)

}
