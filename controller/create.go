package controller

import (
	"encoding/json"
	"fmt"
	"my-App/model"
	"my-App/views"
	"net/http"
)

func create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			data := views.PostRequest{}
			json.NewDecoder(r.Body).Decode(&data)
			fmt.Println(data)
			if err := model.CreateTodo(data.Name, data.Todo); err != nil {
				w.Write([]byte("some error"))
				return
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(data)
		}
	}
}
