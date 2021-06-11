package router

import "net/http"

type Router interface {
	GET(uri string, f func(w http.ResponseWriter, r *http.Request))
	Handler() http.Handler
}

type defaultRouter struct {
	router *http.ServeMux
}

func NewDefaultRouter() Router {
	return &defaultRouter{router:http.NewServeMux()}
}

func (r *defaultRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	r.router.HandleFunc(uri, f)
}
func (r *defaultRouter) Handler() http.Handler {
	return r.router
}