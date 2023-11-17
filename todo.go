package main

type Todo struct {
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name"`
	Completed bool   `json:"completed,omitempty"`
}

var todoMap map[int]Todo
var lastId int = 0

type Todos []Todo

func (t Todos) Len() int {
	return len(t)
}

func (t Todos) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t Todos) Less(i, j int) bool {
	return t[i].ID > t[j].ID
}
