package main

import (
	"chapter1/basicRouter"
	"chapter1/staticDynamic"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	basicRouter.BasicRouter(router)
	staticDynamic.StaticDynamic()
}
