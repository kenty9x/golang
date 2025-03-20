package controller

import (
	"fmt"
	"net/http"
)

func Register() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/getgoing/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("request received")
		fmt.Println(r.Method)
		w.Write([]byte("Hello world"))
	})
	mux.HandleFunc("/getgoing/go", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("request received")
		fmt.Println(r.Method)
		w.Write([]byte("Hello world go"))
	})
	mux.HandleFunc("/ping", ping())
	return mux
}
