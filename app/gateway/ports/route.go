package ports

import "github.com/gorilla/mux"

type Route interface {
	Init(router *mux.Router)
}
