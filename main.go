package main

import (
	"encoding/json"
	mux2 "github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
	"log"
	"net/http"
	"sort"
	"strconv"
)

var rd *render.Render

type Success struct {
	Success bool `json:"success"`
}

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
	var newTodo Todo
	err := json.NewDecoder(request.Body).Decode(&newTodo)
	if err != nil {
		log.Fatal(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	vars := mux2.Vars(request)
	id, _ := strconv.Atoi(vars["id"])
	if todo, ok := todoMap[id]; ok {
		todo.Name = newTodo.Name
		todo.Completed = newTodo.Completed
		rd.JSON(writer, http.StatusOK, Success{true})
		return
	}
	rd.JSON(writer, http.StatusNotFound, Success{false})
}

func DeleteTodoListHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux2.Vars(request)
	id, _ := strconv.Atoi(vars["id"])
	if _, ok := todoMap[id]; ok {
		delete(todoMap, id)
		rd.JSON(writer, http.StatusOK, Success{true})
		return
	}
	rd.JSON(writer, http.StatusNotFound, Success{false})
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
	rd = render.New()
	m := MakeWebHandler()
	n := negroni.Classic()
	n.UseHandler(m)

	log.Println("started app")
	err := http.ListenAndServe(":3000", n)
	if err != nil {
		panic(err)
	}
}
