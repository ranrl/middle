package app

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type middleware func(httprouter.Handle) httprouter.Handle
type Router struct {
	router          *httprouter.Router
	middlewareChain []middleware
}

func NewRouter() *Router {
	return &Router{
		router: httprouter.New(),
	}
}
func (r *Router) Use(midd middleware) {
	r.middlewareChain = append(r.middlewareChain, midd)
}

func (r *Router) Add(method, path string, handle httprouter.Handle) {
	for i := len(r.middlewareChain) - 1; i >= 0; i-- {
		handle = r.middlewareChain[i](handle)
	}
	r.router.Handle(method, path, handle)
}

func (r *Router) Run(address string) error {
	return http.ListenAndServe(address, r.router)
}
