package common

import (
	"github.com/didikprabowo/rajaongkir-go/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Method  string
	Handler http.HandlerFunc
	Path    string
}

func AppRegister() *mux.Router {

	r := mux.NewRouter().StrictSlash(true)
	routes := DataRoutes()
	for _, v := range routes {
		r.Path(v.Path).HandlerFunc(v.Handler).Methods(v.Method)
	}

	return r
}

func JSONFormated(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		h.ServeHTTP(w, r)
	}
}

func DataRoutes() []Route {

	routes := []Route{
		Route{
			Method:  "GET",
			Handler: Welcome,
			Path:    "/welcome",
		},
		Route{
			Method:  "GET",
			Handler: JSONFormated(handlers.GetProvince),
			Path:    "/province",
		},
		Route{
			Method:  "GET",
			Handler: JSONFormated(handlers.GetCity),
			Path:    "/city",
		},
	}
	return routes
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome My Blog .."))
}
