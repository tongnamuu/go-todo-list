package main

import (
	"encoding/json"
	mux2 "github.com/gorilla/mux"
	"github.com/unrolled/render"
	"log"
	"net/http"
	"sort"
)

var rd *render.Render

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
	var todo Todo
	err := json.NewDecoder(request.Body).Decode(&todo)
	if err != nil {
		log.Fatal(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	lastId++
	todo.ID = lastId
	todoMap[lastId] = todo
	rd.JSON(writer, http.StatusCreated, todo)
}

func GetTodoListHandler(writer http.ResponseWriter, request *http.Request) {
	list := make(Todos, 0)
	for _, todo := range todoMap {
		list = append(list, todo)
	}
	sort.Sort(list)
	rd.JSON(writer, http.StatusOK, list)
}

func main() {

}
