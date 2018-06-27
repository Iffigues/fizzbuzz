package main

import(
	"net/http"
	"strings"
)

/*
Routes system of rest api server
*/


type route struct{
	nbr int
	handler http.Handler
}


type routes struct{
	route []*route
}

func (rr *routes)Handler(i int,handler http.Handler){
	rr.route = append(rr.route,&route{i,handler})
}

func (rr  *routes) HandleFunc(i int, handler func(http.ResponseWriter, *http.Request)) {
	rr.route = append(rr.route, &route{i, http.HandlerFunc(handler)})
}


func (rr *routes) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, router := range rr.route {
		var b string = r.URL.Path
		var i int = len(strings.Split(b,"/"))-1
		if router.nbr == i {
			router.handler.ServeHTTP(w, r)
			return
		}
	}
	http.NotFound(w, r)
}

func NewRoute()(r *routes){
	return  &routes{}
}
