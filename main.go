package main

import (
	mux2 "github.com/gorilla/mux"
	"net/http"
)

func MakeWebHandler() http.Handler {
	mux := mux2.NewRouter()
	mux.Handle("/", http.FileServer(http.Dir("public")))
	mux.HandleFunc("/todos", GetTodoListHandler).Methods("GET")
	mux.HandleFunc("/todos", PostTodoListHandler).Methods("POST")
	mux.HandleFunc("/todos/{id:[0-9]+}", DeleteTodoListHandler).Methods("DELETE")
	mux.HandleFunc("/todos/{id:[0-9]+}", UpdateTodoListHandler).Methods("PUT")
	return mux
}

func UpdateTodoListHandler(writer http.ResponseWriter, request *http.Request) {

}

func DeleteTodoListHandler(writer http.ResponseWriter, request *http.Request) {

}

func PostTodoListHandler(writer http.ResponseWriter, request *http.Request) {

}

func GetTodoListHandler(writer http.ResponseWriter, request *http.Request) {

}

func main() {

}
