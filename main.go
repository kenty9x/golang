package main

import (
	"my-App/controller"
	"my-App/model"
	"net/http"
)

func main() {
	mux := controller.Register()
	db := model.Connect()
	defer db.Close()
	http.ListenAndServe("localhost:3000", mux)
}
