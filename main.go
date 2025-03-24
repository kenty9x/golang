package main

import (
	"fmt"
	"log"
	"my-App/controller"
	"my-App/model"
	"net/http"
)

func main() {
	mux := controller.Register()
	db := model.Connect()
	defer db.Close()
	fmt.Println("Serving...")
	log.Fatal(http.ListenAndServe("localhost:3000", mux))
}
