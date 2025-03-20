package main

import (
	"my-App/controller"
	"net/http"
)

func main() {
	mux := controller.Register()
	http.ListenAndServe("localhost:3000", mux)
}
